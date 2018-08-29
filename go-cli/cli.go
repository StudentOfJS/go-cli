package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = " Lets you query IPs, CNAMEs, MX records and Name servers"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "the host to lookup, for example github.com",
		},
		cli.StringFlag{
			Name:  "service",
			Value: "the service to check, for example http",
		},
		cli.StringFlag{
			Name:  "network",
			Value: "the network to check, for exmaple TCP",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ns",
			Usage:  "looks up the name service for a given host",
			Flags:  myFlags,
			Action: NSLookup,
		},
		{
			Name:   "ip",
			Usage:  "looks up the ip address for a requested host",
			Flags:  myFlags,
			Action: IPAddress,
		},
		{
			Name:   "cname",
			Usage:  "Looks up CNAME for a particular host",
			Flags:  myFlags,
			Action: Cname,
		},
		{
			Name:   "mx",
			Usage:  "Looks up mx records for a provided host",
			Flags:  myFlags,
			Action: MXLookup,
		},
		{
			Name:   "port",
			Usage:  "Looks up the port of a service on a host",
			Flags:  myFlags,
			Action: PortLookup,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
