package main

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// NSLookup looks up the host
func NSLookup(c *cli.Context) error {
	ns, err := net.LookupNS(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, n := range ns {
		fmt.Println(n.Host)
	}
	return nil
}

// IPAddress looks up IP address
func IPAddress(c *cli.Context) error {
	ip, err := net.LookupIP(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, a := range ip {
		fmt.Println(a)
	}
	return nil
}

// Cname looks up cname
func Cname(c *cli.Context) error {
	cname, err := net.LookupCNAME(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cname)
	return nil
}

// MXLookup looks up MX records
func MXLookup(c *cli.Context) error {
	mx, err := net.LookupMX(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, m := range mx {
		fmt.Println(m.Host, m.Pref)
	}
	return nil
}

// PortLookup looks up a port of a service
func PortLookup(c *cli.Context) error {
	port, err := net.LookupPort(c.String("network"), c.String("service"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(port)
	return nil
}
