package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewServer(name, version string) *mcp.Server {
	return mcp.NewServer(&mcp.Implementation{
		Name:    name,
		Version: version,
	}, nil)
}

func RunStdio(ctx context.Context, server *mcp.Server) error {
	if err := server.Run(ctx, mcp.NewStdioTransport()); err != nil {
		return err
	}
	return nil
}
