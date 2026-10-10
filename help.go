package commandline

import (
	"fmt"
)

// HELP COMMAND START

type HelpCommand struct {
}

func (helpCommand HelpCommand) Run(cli *CommandLine) TerminateOnCompletion {
	categories := make(map[string][]*CommandWrapper)

	if cli.Graphics.Banner != "" {
		fmt.Println(cli.Graphics.Banner)
	}

	for _, commandWrapper := range cli.commands {
		category := commandWrapper.Category()
		if array, found := categories[category]; found {
			categories[category] = append(array, commandWrapper)
		} else {
			categories[category] = make([]*CommandWrapper, 0)
			categories[category] = append(categories[category], commandWrapper)
		}
	}

	for category, commandWrapperArray := range categories {
		// do not print a category if it is an empty string
		if len(category) != 0 {
			fmt.Printf("[%s]\n", category)
		}
		// group commands together by category
		for _, commandWrapper := range commandWrapperArray {
			if category == "flags" {
				fmt.Printf("\t%s", commandWrapper.identifier)
			} else {
				fmt.Printf("\t%s\t%s\n", commandWrapper.identifier, commandWrapper.description)
				for flag, description := range commandWrapper.variants {
					fmt.Printf("\t\t%s => %s\n", flag, description)
				}
			}
		}
		fmt.Println()
	}

	return Terminate
}

// HELP SCOPE COMMAND

type HelpScopeCommand struct {
}

func (helpScopeCommand HelpScopeCommand) Run(cli *CommandLine) TerminateOnCompletion {
	if len(cli.args) == 1 {
		fmt.Println("you cannot use '?' without specifying a command")
	}

	if (len(cli.args) == 2) && (cli.args[1] == "?") {
		if commandWrapper, found := cli.commands[cli.args[0]]; found {
			fmt.Println(commandWrapper.description)
		} else {
			fmt.Println(cli.args[0] + " is not a command")
		}
	} else if (len(cli.args) == 3) && (cli.args[2] == "?") {
		if commandWrapper, found := cli.commands[cli.args[0]]; found {
			if variant, found := commandWrapper.variants[cli.args[1]]; found {
				fmt.Println(variant)
			} else {
				fmt.Printf("%s [%s] either does not exist or does not have a description")
			}
		}
	} else {
		fmt.Println("invalid format, you must use '[command] ?' or '[command] [flag] ?'")
	}

	return Terminate
}
