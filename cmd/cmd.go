package cmd

import (
	"fmt"
	"github.com/lightningsdk/core/model"
	"os"
	"strings"
)

type commandInput struct {
	Commands []string
	Flags    map[string][]string
}

// the project should wrap this function to build the cli tools, passing in the configuration
func RunCommand(a model.App) error {
	cmds := a.GetCommands()
	ci := parseInput()
	if len(ci.Commands) == 0 {
		fmt.Println("Missing command")
		printHelp(cmds)
	} else {
		var exec func(a model.App) error
		for _, c := range ci.Commands {
			if fc, ok := cmds[c]; ok {
				exec = fc.Function
				if fc.SubCommands != nil {
					cmds = fc.SubCommands
				}
			}
		}
		if exec != nil {
			return exec(a)
		}
		// TODO: a partial match should print requirements or subcommands
		fmt.Println(fmt.Sprintf("Command not found: %s", strings.Join(ci.Commands, " ")))
	}

	return nil
}

func printHelp(cs map[string]model.Command) {
	fmt.Println("Available commands:")
	for k, v := range cs {
		fmt.Println(fmt.Sprintf("    %s: %s", k, v.Help))
	}
}

func parseInput() commandInput {
	ci := commandInput{
		Commands: []string{},
		Flags:    map[string][]string{},
	}

	for i := 1; i < len(os.Args); i++ {

		if len(os.Args[i]) > 1 && os.Args[i][:2] == "--" {
			// Handle arguments starting with '--'
			ci.Flags[os.Args[i]] = []string{os.Args[i+1]}
			i++
		} else if os.Args[i][0] == '-' {
			// Handle arguments starting with a single '-'
			k, v := splitCommand(os.Args[i])
			if v == "" {
				ci.Flags[k] = []string{os.Args[i+1]}
				i++
			} else {
				ci.Flags[k] = []string{v}
			}
		} else {
			ci.Commands = append(ci.Commands, os.Args[i])
		}
	}
	return ci
}

func splitCommand(f string) (string, string) {
	// Remove leading "-" or "--" from the key
	key := f
	if len(f) > 1 && f[:2] == "--" {
		key = f[2:]
	} else if len(f) > 0 && f[0] == '-' {
		key = f[1:]
	}

	// Split by the first occurrence of "="
	equalIndex := -1
	inQuotes := false

	for i, r := range key {
		switch r {
		case '=':
			if !inQuotes {
				equalIndex = i
				break
			}
		case '"', '\'':
			inQuotes = !inQuotes
		}
	}

	if equalIndex == -1 {
		// No "=" outside of quotes, return key and empty value
		return key, ""
	}

	// Separate key and value
	k := key[:equalIndex]
	v := key[equalIndex+1:]

	// Remove surrounding quotes from value
	if len(v) > 1 && ((v[0] == '"' && v[len(v)-1] == '"') || (v[0] == '\'' && v[len(v)-1] == '\'')) {
		v = v[1 : len(v)-1]
	}

	return k, v
}
