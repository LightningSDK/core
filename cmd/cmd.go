package cmd

import (
	"fmt"
	"github.com/lightningsdk/core/model"
	"os"
)

// the project should wrap this function to build the cli tools, passing in the configuration
func RunCommand(a model.App) error {
	fc := os.Args
	cmds := a.GetCommands()
	if len(fc) == 1 {
		fmt.Println("Missing command")
		printHelp(cmds)
	} else if cmd, ok := cmds[fc[1]]; ok {
		return cmd.Function(a)
	} else {
		fmt.Println(fmt.Sprintf("Command not found: %s", fc[0]))
	}

	return nil
}

func printHelp(cs map[string]model.Command) {
	fmt.Println("Available commands:")
	for k, v := range cs {
		fmt.Println(fmt.Sprintf("    %s: %s", k, v.Help))
	}
}
