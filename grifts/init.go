package grifts

import (
	"github.com/avachen2005/buffalo_websocket_cmd/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
