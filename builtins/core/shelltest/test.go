package shelltest

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	proc.GoFunctions["test"] = cmdTest

	defaults.AppendProfile(`
		autocomplete set test { [{
			"Flags": [
				"on",
				"off",
				"auto-report",
				"!auto-report"
			]
		}] }
    `)
}

type testArgs struct {
	OutBlock  string
	OutRegexp string
	ErrBlock  string
	ErrRegexp string
	ExitNum   int
}

func cmdTest(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters.")
	}

	if p.Parameters.Len() == 1 {
		return testConfig(p)
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "define":
		return testDefine(p)

	case "run":
		return testRun(p)

	default:
		return errors.New("Invalid paramter: " + option)
	}
}

func testConfig(p *proc.Process) (err error) {
	option, _ := p.Parameters.String(0)

	switch option {
	case "enable":
		err = p.Config.Set("test", "enabled", true)

	case "!enable", "disable":
		err = p.Config.Set("test", "enabled", false)

	case "auto-report":
		err = p.Config.Set("test", "auto-report", true)

	case "!auto-report":
		err = p.Config.Set("test", "auto-report", false)

	default:
		err = p.Config.Set("test", "enabled", types.IsTrue([]byte(option), 0))
	}

	return
}

func testDefine(p *proc.Process) error {
	enabled, err := p.Config.Get("test", "enabled", types.Boolean)
	if err != nil || !enabled.(bool) {
		return err
	}

	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	b, err := p.Parameters.Byte(2)
	if err != nil {
		return err
	}

	var args testArgs
	err = json.UnmarshalMurex(b, &args)
	if err != nil {
		return err
	}

	// stdout
	rx, err := regexp.Compile(args.OutRegexp)
	if err != nil {
		return err
	}
	stdout := &proc.TestChecks{
		Regexp:   rx,
		Block:    []rune(args.OutBlock),
		RunBlock: runBlock,
	}

	// stderr
	rx, err = regexp.Compile(args.ErrRegexp)
	if err != nil {
		return err
	}
	stderr := &proc.TestChecks{
		Regexp:   rx,
		Block:    []rune(args.ErrBlock),
		RunBlock: runBlock,
	}

	err = p.Tests.Define(name, stdout, stderr, args.ExitNum)
	return err
}

func runBlock(p *proc.Process, block []rune) ([]byte, error) {
	stdout := streams.NewStdin()
	_, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, proc.ShellProcess.Stderr, p)
	if err != nil {
		return nil, err
	}

	b, err := stdout.ReadAll()
	if err != nil {
		return nil, err
	}
	return utils.CrLfTrim(b), nil
}

func testRun(p *proc.Process) error {
	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	branch := p.BranchFID()
	defer branch.Close()

	err = branch.Process.Config.Set("test", "enabled", true)
	if err != nil {
		return err
	}

	err = branch.Process.Config.Set("test", "auto-report", true)
	if err != nil {
		return err
	}

	h := md5.New()
	_, err = h.Write([]byte(time.Now().String() + ":" + strconv.Itoa(p.Id)))
	if err != nil {
		return err
	}

	pipeName := "system_test_" + hex.EncodeToString(h.Sum(nil))

	err = proc.GlobalPipes.CreatePipe(pipeName)
	if err != nil {
		return err
	}

	pipe, err := proc.GlobalPipes.Get(pipeName)
	if err != nil {
		return err
	}

	err = branch.Process.Config.Set("test", "report-pipe", pipeName)
	if err != nil {
		return err
	}

	_, err = lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stderr, branch.Process)
	if err != nil {
		return err
	}

	err = proc.GlobalPipes.Close(pipeName)
	if err != nil {
		return err
	}

	reportType, err := p.Config.Get("test", "report-format", types.String)
	if err != nil {
		return err
	}
	if reportType.(string) == "table" {
		p.Stderr.Writeln([]byte(consts.TestTableHeadings))
	}

	_, err = io.Copy(p.Stderr, pipe)
	return err
}
