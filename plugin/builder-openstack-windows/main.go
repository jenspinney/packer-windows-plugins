package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/packer-community/packer-windows-plugins/builder/openstack-windows"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(openstack.Builder))
	server.Serve()
}
