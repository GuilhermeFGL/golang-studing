package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// GenerateCommandLineApplication Command line application
func GenerateCommandLineApplication() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "To run command line commands"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search IPs in internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}
