package consts

import (
	"io/ioutil"
	"os"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang/tty"
)

// TempDir is the location of temp directory
var TempDir string

func init() {
	var err error

	TempDir, err = ioutil.TempDir("", app.Name)
	if err != nil || TempDir == "" {
		TempDir = tempDir
	}

	if TempDir[len(TempDir)-1:] != PathSlash {
		TempDir += PathSlash
	}

	createDirIfNotExist(TempDir)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			_, err = tty.Stderr.WriteString("WARNING: temp directory doesn't exist and unable to create it. This might cause problems.\nTemp directory: " + dir)

			if err != nil {
				panic("Unable to create tmp directories, unable to write to STDERR. Something is amiss")
			}
		}
	}
}
