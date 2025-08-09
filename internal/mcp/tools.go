package mcp

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/theankitbhardwaj/openrgb-mcp-server/internal/app"
	"github.com/theankitbhardwaj/openrgb-mcp-server/internal/openrgb"
)

func RegisterTools(server *mcp.Server, svc *app.Service) {
	listDevices := &mcp.Tool{
		Name:        "list_devices",
		Description: "List all devices connected to OpenRGB",
	}
	setDeviceColor := &mcp.Tool{
		Name:        "set_device_color",
		Description: "Set the color of a specific device",
	}
	mcp.AddTool(server, listDevices, handleListDevices(svc))
	mcp.AddTool(server, setDeviceColor, handleSetDeviceColor(svc))
}

func handleListDevices(svc *app.Service) mcp.ToolHandlerFor[ListDevicesParams, []openrgb.DeviceInfo] {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[ListDevicesParams]) (*mcp.CallToolResultFor[[]openrgb.DeviceInfo], error) {
		// Call the service to list devices
		devices, err := svc.ListDevices(ctx)
		if err != nil {
			return nil, err
		}

		return &mcp.CallToolResultFor[[]openrgb.DeviceInfo]{
			Content:           []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("List of devices: %v", devices)}},
			StructuredContent: devices,
		}, nil
	}
}

func handleSetDeviceColor(svc *app.Service) mcp.ToolHandlerFor[SetDeviceColorParams, string] {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[SetDeviceColorParams]) (*mcp.CallToolResultFor[string], error) {
		err := svc.SetDeviceColor(ctx, params.Arguments.DeviceID, params.Arguments.R, params.Arguments.G, params.Arguments.B)
		if err != nil {
			return nil, fmt.Errorf("failed to set device color: %w", err)
		}

		return &mcp.CallToolResultFor[string]{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Successfully set color for device: %v", params.Arguments.DeviceID)}},
		}, nil
	}
}

type ListDevicesParams struct {
	// No parameters needed for listing devices
}

type SetDeviceColorParams struct {
	DeviceID int `json:"device_id"`
	R        int `json:"r"`
	G        int `json:"g"`
	B        int `json:"b"`
}
