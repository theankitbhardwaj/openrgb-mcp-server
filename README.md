# OpenRGB MCP Server

A server written in Go that exposes OpenRGB controls via the **Model Context Protocol (MCP)**. This allows AI agents and other tools that support MCP (like GitHub Copilot in VS Code) to interact with and control your local RGB lighting setup using natural language.

## About The Project

This project acts as a bridge between the abstract world of Large Language Models and your physical hardware. It uses the official `modelcontextprotocol/go-sdk` to create an MCP server and the `csutorasa/go-openrgb-sdk` to communicate with the OpenRGB application.

The result is a powerful tool that enables you to control your RGB devices directly from your code editor's chat interface.

### Implemented MCP Tools

* `list_devices`: Lists all connected RGB devices detected by OpenRGB.
* `set_device_color`: Sets a specific device to a solid RGB color.
* `set_all_color`: Sets all devices color to a solid RGB color.
* `list_profiles`: Lists all saved profiles on OpenRGB.
* `set_profile`: Sets a saved profile.

## Getting Started

Follow these steps to get the server up and running on your local machine.

### Prerequisites

* **Go**: Version 1.21 or later.
* **OpenRGB**: The OpenRGB application must be installed and running. (If running on another machine, host and port should be reachable by this server. Edit config/config.yaml to set server host and port).
* **Enable OpenRGB SDK Server**: In OpenRGB, go to the **SDK Server** tab and click **Start Server**. It must be running for this application to connect.


### Installation

1. **Clone the repository:**

```sh
git clone https://github.com/theankitbhardwaj/openrgb-mcp-server.git
cd openrgb-mcp-server
```

2. **Install Go dependencies:**

```sh
go mod tidy
```

3. **Build the application:**

```sh
go build -o ./bin/openrgb-mcp-server ./cmd/server/main.go
```


## Usage

### With Visual Studio Code

1. Create a `.vscode` folder in the root of your project if it doesn't exist.
2. Inside `.vscode`, create a file named `mcp.json` with the following content:

```json
{
    "servers": {
        "OpenRGB": {
            "type": "stdio",
            "command": "${workspaceFolder}/bin/openrgb-mcp-server",
            "args": []
        }
    }
}
```
3. Optionally update your config file under `config/config.yaml`.

4. Open the Chat view in VS Code (e.g., GitHub Copilot Chat in Agent Mode). You can now make requests to control your lights.

**Example Prompts:**
> `@workspace can you list my RGB devices for me?`

> `@workspace please set device 0 to a bright blue color.`

> `@workspace please set all devices to a dark red color.`

> `@workspace can you list my saved openrgb profiles?`

> `@workspace set GamingMode profile.`


**You can also follow your specific IDE or tool like Claude Code's instructions to setup a MCP Server.**

## Project Structure

The project follows the standard Go project layout for clarity and maintainability.

```
├── cmd/server/       # Main application entry point.
├── internal/
│   ├── app/          # Core application logic and service definitions.
│   ├── mcp/          # MCP server setup, tool definitions, and handlers.
│   └── openrgb/      # Wrapper around the OpenRGB SDK for stable interactions.
├── pkg/
│   └── util/         # Shared utility functions. 
└── config/           # Configuration files (e.g., config.yaml).
```


## Roadmap

* [ ] Add more tools (`set_device_mode`, `set_brightness`, `set_zone_color`).
* [ ] Support saving profiles.
* [ ] Implement an optional HTTP transport for the MCP server.


## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the GPL v3.0 License. See `LICENSE` for more information.

## Acknowledgments

* [modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk)
* [csutorasa/go-openrgb-sdk](https://github.com/csutorasa/go-openrgb-sdk)
* [OpenRGB](https://openrgb.org/)

