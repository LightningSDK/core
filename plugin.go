package core

import (
	"encoding/json"
	"github.com/lightningsdk/core/cmd"
	"github.com/lightningsdk/core/http"
	"github.com/lightningsdk/core/model"
)

type Configuration struct {
	HTTP struct {
		Port int `json:"port"`
	}
}

type core struct {
	conf *Configuration
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

func (c *core) Configure(cfg []byte) {
	c.conf = &Configuration{}
	err := json.Unmarshal(cfg, c.conf)
	if err != nil {
		panic(err)
	}
}
