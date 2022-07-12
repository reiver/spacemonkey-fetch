package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var target string
	{
		if len(os.Args) < 2 {
			fmt.Fprintln(os.Stderr, "ERROR: bad request: missing target")
			os.Exit(1)
			return
		}

		target = os.Args[1]
	}

	var kind string
	{
		const colon string = ":"

		index := strings.Index(target, colon)

		if index < 0 {
			fmt.Fprintf(os.Stderr, "ERROR: bad request: bad target (%q)\n", target)
			os.Exit(1)
			return
		}

		kind = target[:index]

		if " " == kind {
			fmt.Fprintf(os.Stderr, "ERROR: bad request: bad target (%q)\n", target)
			os.Exit(1)
			return
		}
	}

	var filename string
	{
		const format string = "spacemonkey-fetch-%s"

		filename = fmt.Sprintf(format, kind)

	}

	var path string
	{
		var err error

		path, err = exec.LookPath(filename)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: not found: command %q not found in path\n", filename)
			os.Exit(1)
			return
		}
	}

	var workingdir string
	{
		var err error

		workingdir, err = os.Getwd()
		if nil != err {
			fmt.Fprintln(os.Stderr, "ERROR: internal error")
			os.Exit(1)
			return
		}
	}

	var cmd exec.Cmd
	{

		cmd.Path = path

		if nil == cmd.Args {
			cmd.Args = make([]string, 2)
		}
		cmd.Args[0] = path
		cmd.Args[1] = target

		cmd.Dir = workingdir

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	{

		err := cmd.Run()
		if nil != err {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
			return
		}
	}
}
