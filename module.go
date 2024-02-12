package core

import (
	"errors"
	"gopkg.in/yaml.v3"
)

type Module interface {
	GetRenderers() map[string]Renderer
	GetHandlers() []Handler
	GetEmptyConfig() any
	SetConfig(any)
}

type Modules map[string]Module

func (ms *Modules) UnmarshalYAML(value *yaml.Node) error {
	for i := 0; i < len(value.Content); i += 2 {
		if value.Content[i].Tag != "!!str" {
			return errors.New("module config name is not a string")
		}
		n := value.Content[i].Value
		c := (*ms)[n].GetEmptyConfig()
		if c == nil {
			continue
		}
		if value.Content[i+1].Tag != "!!map" {
			return errors.New("module config is not a map")
		}
		err := value.Content[i+1].Decode(c)
		if err != nil {
			return err
		}
		(*ms)[n].SetConfig(c)
	}
	return nil
}

type DefaultModule struct{}

func (m *DefaultModule) GetRenderers() map[string]Renderer {
	return nil
}
func (m *DefaultModule) GetHandlers() []Handler {
	return nil
}
func (m *DefaultModule) GetEmptyConfig() any {
	return nil
}
func (m *DefaultModule) SetConfig(cfg any) {}

func SetConfig(cfg any) {}
