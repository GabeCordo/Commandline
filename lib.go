package commandline

import (
	"github.com/GabeCordo/Commandline/paths"
)

// DEBUG COMMAND START

type DebugCommand struct {
}

func (dc DebugCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Debug] = true

	return Continue // do not terminate
}

// CREATE COMMAND START

type CreateCommand struct {
}

func (cc CreateCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Create] = true

	return Continue // do not terminate
}

// DELETE COMMAND START

type DeleteCommand struct {
}

func (dc DeleteCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Delete] = true

	return Continue // do not terminate
}

// SHOW COMMAND START

type ShowCommand struct {
}

func (sc ShowCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Show] = true

	return Continue
}

// SWITCH COMMAND START

type SwitchCommand struct {
}

func (sc SwitchCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Switch] = true

	return Continue
}

// INSTALL COMMAND START

type InstallCommand struct {
}

func (ic InstallCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Install] = true

	return Continue
}

// ADD COMMAND

type AddCommand struct {
}

func (addCommand AddCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Add] = true

	return Continue
}

// UPDATE COMMAND

type UpdateCommand struct {
}

func (uc UpdateCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Update] = true

	return Continue
}

// REVOKE COMMAND

type RevokeCommand struct {
}

func (rc RevokeCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Revoke] = true

	return Continue
}

// STOP COMMAND

type StopCommand struct {
}

func (command StopCommand) Run(cli *CommandLine) TerminateOnCompletion {
	cli.flags[Stop] = true

	return Continue
}

// DEFAULT COMMAND LINE

func NewCommandLine(path ...paths.Path) *CommandLine {

	cli := new(CommandLine)

	// there may be a persistent JSON file to load Config data from
	if len(path) > 0 {
		config := new(Config)
		config.FromJson(path[0])

		cli.Config = config
	}

	// make Command map
	cli.commands = make(map[string]*CommandWrapper)

	// help commands
	cli.AddCommand("help", HelpCommand{}).SetCategory("core").SetDescription("displays all commands that can be invoked")
	cli.AddCommand("?", HelpScopeCommand{}).SetCategory("core").SetDescription("displays help information for a command or command flag variant")
	// cli flag sub-commands
	cli.AddCommand("create", CreateCommand{}).SetCategory("flags")
	cli.AddCommand("delete", DeleteCommand{}).SetCategory("flags")
	cli.AddCommand("show", ShowCommand{}).SetCategory("flags")
	cli.AddCommand("switch", SwitchCommand{}).SetCategory("flags")
	cli.AddCommand("debug", DebugCommand{}).SetCategory("flags")
	cli.AddCommand("install", InstallCommand{}).SetCategory("flags")
	cli.AddCommand("add", AddCommand{}).SetCategory("flags")
	cli.AddCommand("update", UpdateCommand{}).SetCategory("flags")
	cli.AddCommand("revoke", RevokeCommand{}).SetCategory("flags")
	cli.AddCommand("stop", StopCommand{}).SetCategory("flags")

	return cli
}
