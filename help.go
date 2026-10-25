package commandline

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

	// the help command shall sort the categories by alphabetical order.
	// note: iterating through a map would result in non-deterministic behaviour.
	categoryNames := make([]string, 0, len(categories))
	for category := range categories {
		categoryNames = append(categoryNames, category)
	}
	sort.Strings(categoryNames)

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	defer writer.Flush()

	for _, category := range categoryNames {
		commandWrapperArray := categories[category]

		// do not print the core or flags category
		if (category == "flags") || (category == "core") {
			continue
		}

		// do not print a category if it is an empty string
		if len(category) != 0 {
			fmt.Fprintf(writer, "[%s]\n", category)
		}

		// the help command shall sort the commands by alphabetical order.
		// note: iterating through a map would result in non-deterministic behaviour.
		sort.Slice(commandWrapperArray, func(i, j int) bool {
			return commandWrapperArray[i].identifier < commandWrapperArray[j].identifier
		})

		for _, commandWrapper := range commandWrapperArray {
			fmt.Fprintf(writer, "  %s\t%s\n", commandWrapper.identifier, commandWrapper.description)

			for flag, description := range commandWrapper.variants {
				fmt.Fprintf(writer, "    %s\t%s\n", flag, description)
			}
		}

		fmt.Fprintln(writer)
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
				fmt.Printf("%s [%s] either does not exist or does not have a description\n", cli.args[0], cli.args[1])
			}
		}
	} else {
		fmt.Println("invalid format, you must use '[command] ?' or '[command] [flag] ?'")
	}

	return Terminate
}
