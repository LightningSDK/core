package core

import (
	"github.com/lightningsdk/core/cmd"
	"github.com/lightningsdk/core/http"
	"github.com/lightningsdk/core/model"
	"github.com/lightningsdk/core/service"
)

type core struct {
	model.DefaultPlugin
}

func NewPlugin() model.Plugin {
	return &core{}
}

func (c *core) GetRoutes() []model.Route {
	return nil
}

func (c *core) GetCommands() map[string]model.Command {
	return map[string]model.Command{
		"build": cmd.GetAutogenerateCommand(),
		"http":  http.GetCommand(),
	}
}

// NewApp creates and returns an instance of App
func NewApp() *service.App {
	return &service.App{
		Plugins: map[string]model.Plugin{
			"core": &core{},
		},
	}
}
