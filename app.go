package core

import (
	"github.com/lightningsdk/core/model"
	"github.com/lightningsdk/core/service"
)

// NewApp creates and returns an instance of App
func Run(cfg string, plugins map[string]model.Plugin) {
	app := &service.App{
		Plugins: map[string]model.Plugin{
			"core": &core{},
		},
	}
	// this loads any included plugins
	app.RegisterPlugins(plugins)
	app.Bootstrap(cfg)
	app.Run()
}
