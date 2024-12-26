package service

import (
	"fmt"
	"github.com/lightningsdk/core/cmd"
	"github.com/lightningsdk/core/model"
	"golang.org/x/exp/maps"
)

type App struct {
	Include  []string
	Plugins  map[string]model.Plugin
	Routes   []model.Route
	Commands map[string]model.Command
	configs  *Config
}

func (a *App) Bootstrap(configPath string) {
	// load the configurations into raw json, ready for each plugin to parse itself
	a.configs = loadConfigurations(configPath)

	for p, plugin := range a.Plugins {
		perm := createPermissionToken(p)
		plugin.SetPermissionToken(perm)
		cfg := a.configs.getConfig(p)
		plugin.Configure(cfg)
	}

	// load the commands and routes
	if a.Commands == nil {
		a.Commands = map[string]model.Command{}
	}
	if a.Routes == nil {
		a.Routes = []model.Route{}
	}
	for _, plugin := range a.Plugins {
		a.Routes = append(a.Routes, plugin.GetRoutes()...)
		maps.Copy(a.Commands, plugin.GetCommands())
	}
}

func (a *App) Run() {
	err := cmd.RunCommand(a)
	fmt.Println(err)
}

// TOOD: this should be protected
func (a *App) GetCommands() map[string]model.Command {
	return a.Commands
}

func (a *App) GetPlugins() map[string]model.Plugin {
	return a.Plugins
}

func (a *App) RegisterPlugins(p map[string]model.Plugin) {
	a.Plugins = p
}
