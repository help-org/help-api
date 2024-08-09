package cli

import cli "directory/cli/server"

func Run() {
	command := &cli.Command{}

	err := command.Serve()
	if err != nil {
		return
	}
}
