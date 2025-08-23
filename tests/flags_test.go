package tests

import (
	"github.com/FortifiedCode/commandline"
	"testing"
)

func TestFlagToString(t *testing.T) {
	if "add" != commandline.Add.ToString() {
		t.Error("flag.ToString() did not pass")
	}

	if "revoke" != commandline.Revoke.ToString() {
		t.Error("flag.ToString() did not pass")
	}
}

func TestStringToFlag(t *testing.T) {
	if commandline.NotAFlag != commandline.FlagFromString("fwefo") {
		t.Error("flag.FlagFromString() not rejecting invalid flags")
	}

	if commandline.Create != commandline.FlagFromString("create") {
		t.Error("flag.FlagFromString() not converting string to flags correctly")
	}
}

func TestCommandLineFlag(t *testing.T) {
	cli := commandline.NewCommandLine()

	if cli.Flag(commandline.Add) != false {
		t.Error("CommandLine flags are not being initialized to false")
	}
}
