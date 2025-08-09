package main

import (
	"fmt"

	"github.com/theankitbhardwaj/openrgb-mcp-server/internal/openrgb"
	"github.com/theankitbhardwaj/openrgb-mcp-server/pkg/util"
)

func main() {
	fmt.Println("OpenRGB MCP server")

	cfg, err := util.LoadConfig("config/config.yaml")

	if err != nil {
		fmt.Println(err)
	}

	_, err = openrgb.ConnectClient(cfg.OpenRGB.Host, cfg.OpenRGB.Port) // client, err

	if err != nil {
		fmt.Printf("Failed to connect to OpenRGB server: %v\n", err)
		return
	}
}
