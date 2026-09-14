package tests

import (
	"github.com/GabeCordo/Commandline"
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
