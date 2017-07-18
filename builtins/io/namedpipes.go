package io

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	proc.GoFunctions["pipe"] = proc.GoFunction{Func: cmdPipe, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["<read-pipe>"] = proc.GoFunction{Func: cmdReadPipe, TypeIn: types.Null, TypeOut: types.Generic}
}

func cmdPipe(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	flag, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	args := p.Parameters.StringArray()
	if err != nil {
		return err
	}
	if len(args) < 1 {
		return errors.New("Not enough parameters!")
	}
	args = args[1:]

	switch flag {
	case "--create", "-c":
		for i := range args {
			err := proc.GlobalPipes.CreatePipe(args[i])
			if err != nil {
				return err
			}
		}

	case "--close", "-x":
		for i := range args {
			err := proc.GlobalPipes.Close(args[i])
			if err != nil {
				return err
			}
		}

	case "--file", "-f":
		err := proc.GlobalPipes.CreateFile(args[0], args[1])
		if err != nil {
			return err
		}

	default:
		return errors.New("Invalid parameters. Please include either `--create` or `--close`.")
	}

	return nil
}

func cmdReadPipe(p *proc.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	pipe, err := proc.GlobalPipes.Get(name)
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(pipe.GetDataType())
	_, err = io.Copy(p.Stdout, pipe)
	return err
}