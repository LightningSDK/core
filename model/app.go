package model

type App interface {
	Run()
	Bootstrap(configPath string)
	GetCommands() map[string]Command
	GetPlugins() map[string]Plugin
}
