package tests

import (
	"github.com/FortifiedCode/commandline"
	"testing"
)

type TestCommand struct {
}

func (tc TestCommand) Run(cli *commandline.CommandLine) commandline.TerminateOnCompletion {
	return commandline.Terminate
}

func TestCommandLine(t *testing.T) {

	cli := commandline.NewCommandLine()
	cli.AddCommand("test", TestCommand{}).Description("test description")
}
