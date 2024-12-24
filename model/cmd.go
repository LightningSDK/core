package model

type Command struct {
	Function    func(a App) error
	Help        string
	SubCommands map[string]Command
}
