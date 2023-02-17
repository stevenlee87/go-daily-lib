package main

import "github.com/stevenlee87/go-daily-lib/cobra/cmd"

func main() {
	cmd.Execute()
}

/*
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cmd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version number of Hugo

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cmd
  -l, --license string   name of license for the project
      --viper            use Viper for configuration (default true)

Use "cmd [command] --help" for more information about a command.
*/
