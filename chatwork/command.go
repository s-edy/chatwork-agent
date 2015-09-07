package chatwork

import (
	"../_vendor/src/github.com/BurntSushi/toml"
	"../_vendor/src/github.com/codegangsta/cli"
	"log"
	"os"
)

type Command struct {
	app *cli.App
}

func NewCommand() *Command {
	app := cli.NewApp()
	app.Name = "chatwork-agent"
	app.Usage = "Chatwork proxy agent"
	app.Version = VERSION
	app.Author = "Shinichiro Yuki"
	app.Email = "sinycourage@gmail.com"

	command := &Command{
		app: app,
	}

	command.app.Commands = []cli.Command{
		cli.Command{
			Name: "start",
			Usage: "Start agent",
			Description: "Listening on HTTP",
			Action: command.Start,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "config, c",
					Usage: "Specify TOML file",
				},
			},
		},
	}

	return command
}

func (this *Command) Run() {
	this.app.Run(os.Args)
}

func (this *Command) Start(cli *cli.Context) {
	Configuration.Port = DEFAULT_LISTEN_PORT
	_, err := toml.DecodeFile(cli.String("config"), &Configuration)
	if err != nil {
		log.Fatalln(err)
	}

	server := NewServer(Configuration)
	err = server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
