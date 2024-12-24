package model

type Plugin interface {
	GetRoutes() []Route
	GetCommands() map[string]Command
	GetEmptyConfig() any
	SetConfig(any)
}

type DefaultPlugin struct{}

func (m *DefaultPlugin) GetCommands() map[string]Command {
	return nil
}
func (m *DefaultPlugin) GetRoutes() []Route {
	return nil
}
func (m *DefaultPlugin) GetEmptyConfig() any {
	return nil
}
func (m *DefaultPlugin) SetConfig(cfg any) {}

func SetConfig(cfg any) {}
