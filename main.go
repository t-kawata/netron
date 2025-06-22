package main

//go:generate go run ./api/generate ./api/public/functions.tmpl ./api/public/index.tmpl ./api/public/index.html
import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/mudler/netron/cmd"
	internal "github.com/mudler/netron/internal"
)

func main() {

	app := &cli.App{
		Name:        "netron",
		Version:     internal.Version,
		Authors:     []*cli.Author{{Name: "ShyMe Inc. Toshimi Kawata"}},
		Usage:       "netron --config /etc/netron/config.yaml",
		Description: "netron builds an immutable trusted blockchain addressable p2p network",
		Copyright:   cmd.Copyright,
		Flags:       cmd.MainFlags(),
		Commands: []*cli.Command{
			cmd.Start(),
			cmd.API(),
			cmd.ServiceAdd(),
			cmd.ServiceConnect(),
			cmd.FileReceive(),
			cmd.Proxy(),
			cmd.FileSend(),
			cmd.DNS(),
			cmd.Peergate(),
		},

		Action: cmd.Main(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
