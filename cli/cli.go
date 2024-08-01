package cli

import cli "directory/cli/server"

func Run() {
	command := &cli.Command{}

	err := command.Run()
	if err != nil {
		return
	}
}
