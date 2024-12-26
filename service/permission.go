package service

import "github.com/lightningsdk/core/model"

type certified struct{}

var c certified

func init() {
	c = certified{}
}

func createPermissionToken(plugin string) *model.PermissionToken {
	return model.CreatePermissionToken(plugin, c)
}

func ValidatePermission(plugin string, token *model.PermissionToken) bool {
	return token.GetPermission() == c && token.GetPlugin() == plugin
}
