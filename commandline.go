package commandline

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	FinalArg string = ""
)

type TerminateOnCompletion bool

const (
	Terminate TerminateOnCompletion = true
	Continue  TerminateOnCompletion = false
)

type CommandLine struct {
	Config *Config

	flags [numOfFlags]bool

	MetaData struct {
		WorkingDirectory string
		TimeCalled       time.Time
	}

	args        []string
	numOfArgs   int
	argsPointer int

	commands map[string]*CommandWrapper
}

func (cli *CommandLine) CollectMetaData() {
	cli.MetaData.TimeCalled = time.Now()
	cli.MetaData.WorkingDirectory, _ = os.Getwd()
}

func (cli *CommandLine) NextArg() string {
	cli.argsPointer++

	if cli.argsPointer < cli.numOfArgs {
		return cli.args[cli.argsPointer]
	} else {
		return FinalArg
	}
}

func (cli *CommandLine) AddCommand(identifier string, function Command) *CommandWrapper {
	if _, found := cli.commands[identifier]; found {
		panic("the command " + identifier + " already exists")
	}

	commandWrapper := NewCommandWrapper(identifier, function)
	cli.commands[identifier] = commandWrapper

	return commandWrapper
}

// UnknownCommand
// @runtime O(keywords * chars)
func (cli CommandLine) UnknownCommand(identifier string) TerminateOnCompletion {

	// populate a map with all keywords where the first char of the keyword matches the first char of the identifier
	// O(keywords)
	keywords := make(map[string]int)
	for keyword, _ := range cli.commands {
		if keyword[0] == identifier[0] {
			keywords[keyword] = 1
		}
	}

	// whatever word has the greatest number of chars common with the unknown identifier will be outputted
	// if there is more than one identifier that shared the number of common chars, output them all
	mostNumOfCharOccurrences := 1

	// start at the second index since we know the first one is already common
	for c := 1; c < len(identifier); c++ {
		char := string(identifier[c])
		for keyword, _ := range keywords {
			if strings.Contains(keyword, char) {
				keywords[keyword]++

				if keywords[keyword] > mostNumOfCharOccurrences {
					mostNumOfCharOccurrences = keywords[keyword]
				}
			}
		}
	}

	var output string
	if len(keywords) == 0 {
		output = "Unknown identifier \"%s\", there are no matching commands.\n"
	} else {
		output = "Unknown identifier \"%s\", did you mean?\n"
	}
	fmt.Printf(output, identifier)

	// if keywords exist in the map, for loop entered, otherwise it will skip (this is why
	// we do not need to worry about averageOfCharOccurrences being initialized if numOfKeywords is zero
	for keyword, count := range keywords {
		//if count >= averageOfCharOccurrences {
		if count == mostNumOfCharOccurrences {
			fmt.Println(keyword)
		}
	}

	return true // terminate since we don't know how to execute all commands
}

func (cli CommandLine) Flag(flag Flag) bool {
	if flag == NotAFlag {
		return false
	} else {
		return cli.flags[flag]
	}
}

func (cli *CommandLine) Run() {

	// this data can be useful for capturing the environment the exe was called
	cli.CollectMetaData()

	// push the command line arguments into struct
	cli.args = os.Args[1:] // strip out the file descriptor in position 0
	cli.numOfArgs = len(cli.args)
	cli.argsPointer = 0

	for arg := cli.args[0]; arg != FinalArg; arg = cli.NextArg() {
		var terminateFlag TerminateOnCompletion = false
		if commandWrapper, found := cli.commands[arg]; found {
			terminateFlag = commandWrapper.Command().Run(cli)
		} else {
			terminateFlag = cli.UnknownCommand(arg)
		}

		if terminateFlag {
			return
		}
	}

	// the commandline tool shall display the `help` callback when no arguments are provided
	if len(cli.args) == 0 {
		if commandWrapper, found := cli.commands["help"]; found {
			commandWrapper.Command().Run(cli)
		}
	}

	// stop reading cli arguments
}
