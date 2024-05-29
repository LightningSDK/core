package core

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"strings"
)

// the project should wrap this function to build the cli tools, passing in the configuration
func RunCommand(a *App) error {
	commands := getCommands(a)
	// load command from plugins here

	fc := os.Args
	if len(fc) == 1 {
		fmt.Println("Missing command")
		printHelp(commands)
	} else if cmd, ok := commands[fc[1]]; ok {
		return cmd.Function(a)
	} else {
		fmt.Println(fmt.Sprintf("Command not found: %s", fc[0]))
	}

	return nil
}

func printHelp(cs map[string]Command) {
	fmt.Println("Available commands:")
	for k, v := range cs {
		fmt.Println(fmt.Sprintf("    %s: %s", k, v.Help))
	}
}

type Command struct {
	Function func(a *App) error
	Help     string
}

func getCommands(a *App) map[string]Command {
	cmds := map[string]Command{
		"autogenerate": {
			Function: autognenerate,
			Help:     "Automatically generate required module imports",
		},
	}
	for _, m := range a.Modules {
		if add := m.GetCommands(); add != nil {
			maps.Copy(cmds, add)
		}
	}

	return cmds
}

// this builds the module registry file
func autognenerate(a *App) error {
	// create the contents
	fmt.Println("autogenerating ...")
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("package autogen\n\n")
	buf.WriteString("import (\n")
	// first add the coe
	buf.WriteString(fmt.Sprintf("\t\"%s\"\n", "github.com/lightningsdk/core"))
	if len(a.Modules) > 0 {
		for k, _ := range a.Modules {
			fmt.Println(fmt.Sprintf("\tadding module: %s", k))
			buf.WriteString(fmt.Sprintf("\t\"%s\"\n", k))
		}
	}
	buf.WriteString(")\n\n")

	buf.WriteString("func GetModules(app *core.App) map[string]core.Module {\n")
	buf.WriteString("\tmodules := map[string]core.Module{}\n")
	for k, _ := range a.Modules {
		s := strings.Split(k, "/")
		pkg := s[len(s)-1]
		buf.WriteString(fmt.Sprintf("\tmodules[\"%s\"] = %s.NewModule(app)\n", k, pkg))
	}
	buf.WriteString("\treturn modules\n")
	buf.WriteString("}\n")

	// write the file
	return os.WriteFile("./autogen/init_modules.go", buf.Bytes(), 0644)
}
