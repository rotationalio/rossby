package main

import (
	"os"

	"github.com/kansaslabs/rossby"
	"github.com/urfave/cli"
)

func main() {

	// Instantiate the CLI application
	app := cli.NewApp()
	app.Name = "rossby"
	app.Version = rossby.Version(false)
	app.Usage = "rossby is a distributed encrypted message service"

	// Define the commands available to the application
	app.Commands = []cli.Command{
		{
			Name:     "serve",
			Usage:    "run the rossby server",
			Action:   serve,
			Category: "server",
			Flags:    []cli.Flag{},
		},
	}

	// Run the CLI program
	app.Run(os.Args)
}

//===========================================================================
// Server Commands
//===========================================================================

func serve(c *cli.Context) (err error) {
	return nil
}
