package model

type Plugin interface {
	GetRoutes() []Route
	GetCommands() map[string]Command
	Configure([]byte)
	SetPermissionToken(permissionToken *PermissionToken)
}

type DefaultPlugin struct {
	permissionToken *PermissionToken
}

func (m *DefaultPlugin) GetCommands() map[string]Command {
	return nil
}
func (m *DefaultPlugin) GetRoutes() []Route {
	return nil
}
func (m *DefaultPlugin) Configure(cfg []byte) {}
func (m *DefaultPlugin) SetPermissionToken(permissionToken *PermissionToken) {
	m.permissionToken = permissionToken
}
