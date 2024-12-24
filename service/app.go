package service

import (
	"fmt"
	"github.com/lightningsdk/core/cmd"
	"github.com/lightningsdk/core/model"
	"golang.org/x/exp/maps"
)

type App struct {
	Include  []string                 `yaml:"include"`
	Plugins  map[string]model.Plugin  `yaml:"plugins"`
	Routes   []model.Route            `yaml:"routes"`
	Commands map[string]model.Command `yaml:"commands"`
}

func (a *App) Bootstrap() {
	if a.Commands == nil {
		a.Commands = map[string]model.Command{}
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
