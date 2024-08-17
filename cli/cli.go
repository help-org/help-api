package cli

import (
	"directory/cli/server"
)

func Run() {
	err := server.Serve()
	if err != nil {
		return
	}
}
