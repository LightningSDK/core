package service

import (
	"github.com/lightningsdk/core/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Permission(t *testing.T) {
	token := createPermissionToken("plugin")
	assert.True(t, ValidatePermission("plugin", token))

	badToken := model.CreatePermissionToken("plugin", "test")
	assert.False(t, ValidatePermission("plugin", badToken))
}
