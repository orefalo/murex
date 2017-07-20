// +build !windows

package man

import (
	"bufio"
	"compress/gzip"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/utils"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	paths        []string
	manId        []string       = []string{"man1", "man6", "man7", "man8"}
	rxMatchFlags *regexp.Regexp = regexp.MustCompile(`\\fB(\\-[a-zA-Z0-9]|\\-\\-[\\\-a-zA-Z0-9]+).*?\\f[RP]`)
)

/*
MANUAL SECTIONS
    The standard sections of the manual include:

    1      User Commands
    2      System Calls
    3      C Library Functions
    4      Devices and Special Files
    5      File Formats and Conventions
    6      Games et. al.
    7      Miscellanea
    8      System Administration tools and Daemons

    Distributions customize the manual section to their specifics,
    which often include additional sections.
*/

func Initialise() {
	debug.Log("Initialising man pages...")
	cmd := exec.Command("man", "-w")

	b, err := cmd.Output()
	if err != nil {
		os.Stderr.WriteString("Error initialising man pages: " + err.Error() + utils.NewLineString)
		return
	}

	s := strings.TrimSpace(string(b))
	paths = strings.Split(s, ":")
	debug.Log(paths)
}

func ScanManPages(exe string) (flags []string) {
	fMap := make(map[string]bool)
	for i := range paths {
		for j := range manId {
			files, err := filepath.Glob(paths[i] + "/" + manId[j] + "/" + exe + ".*.gz")
			if err != nil {
				continue
			}
			for k := range files {
				parseManPage(&fMap, files[k])
			}
		}
	}

	for f := range fMap {
		flags = append(flags, f)
	}
	sort.Strings(flags)
	return
}

func parseManPage(flags *map[string]bool, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		return
	}
	defer gz.Close()

	scanner := bufio.NewScanner(gz)

	for scanner.Scan() {
		match := rxMatchFlags.FindAllStringSubmatch(scanner.Text(), -1)
		for i := range match {
			if len(match[i]) > 0 {
				s := strings.Replace(match[i][1], `\`, "", -1)
				if strings.HasSuffix(s, "fR") || strings.HasSuffix(s, "fP") {
					s = s[:len(s)-2]
				}
				(*flags)[s] = true
			}
		}
	}

	return
}