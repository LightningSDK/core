package core

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func StartService(a *App) error {
	fmt.Println(a)

	r := gin.Default()
	configureGin(r, a)

	return r.Run("localhost:9000")
}

func configureGin(r *gin.Engine, app *App) {
	for _, m := range app.Modules {
		for _, h := range m.GetHandlers() {
			r.Handle(h.Method, h.Endpoint, func(c *gin.Context) {
				// this is a temporary wrapper and should be updated to a more elegant system
				ctx := context.WithValue(c.Request.Context(), "params", c.Params)
				resp, err := h.Handle(ctx, c.Request)
				if err != nil {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
				if resp.Body != nil {
					b, err := io.ReadAll(resp.Body)
					if err != nil {
						c.AbortWithStatus(http.StatusBadRequest)
						return
					}
					c.Data(200, resp.Header["Content-Type"][0], b)
				}
			})
		}
	}
}
