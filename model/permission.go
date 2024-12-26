package model

type PermissionToken struct {
	plugin    string
	certified any
}

func CreatePermissionToken(plugin string, certified any) *PermissionToken {
	return &PermissionToken{
		plugin:    plugin,
		certified: certified,
	}
}

func (p *PermissionToken) GetPlugin() any {
	return p.plugin
}

func (p *PermissionToken) GetPermission() any {
	return p.certified
}
