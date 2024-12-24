package model

type App interface {
	Run()
	Bootstrap()
	GetCommands() map[string]Command
	GetPlugins() map[string]Plugin
}
