package core

import "context"

type Renderer interface {
	RenderBlock(context.Context, map[string]string) (string, error)
}
