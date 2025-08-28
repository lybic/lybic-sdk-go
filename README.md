<p align="center">
  <a href="https://lybic.ai/">
    <img src="https://avatars.githubusercontent.com/lybic" alt="Lybic Logo" width="120" height="120">
  </a>
</p>

<h1 align="center">Lybic SDK for Golang</h1>

<p align="center">
  <a href="https://github.com/lybic/lybic-sdk-go/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/pypi/l/lybic"></a>
  <a href="https://lybic.ai/docs"><img alt="Documentation" src="https://img.shields.io/badge/documentation-Lybic-orange"></a>
  <a href="https://github.com/lybic/lybic-sdk-go/actions/workflows/quality.yml"><img alt="Golanglint" src="https://github.com/lybic/lybic-sdk-go/actions/workflows/quality.yml/badge.svg"></a>
  <a href="https://github.com/lybic/lybic-sdk-go/releases"><img alt="GitHub Release" src="https://img.shields.io/github/v/release/lybic/lybic-sdk-go?display_name=tag&sort=semver"></a>
</p>

Developing, testing, and deploying GUI-based AI agents is complex. Developers waste precious time wrestling with cloud instances, VNC servers, and environment configurations instead of focusing on what matters: building intelligent agents.

**Lybic is the infrastructure layer for your GUI agents.**

**Lybic** (/ˈlaɪbɪk/) provides a robust, on-demand infrastructure platform designed specifically for the AI agent development lifecycle. This SDK for Go is your command center for programmatically controlling the entire Lybic ecosystem, empowering you to build, test, and scale your agents with unprecedented speed and simplicity.

## 🚀 Getting Started

### 1. Installation

First, install the package from the Go package repository:

```bash
go get github.com/lybic/lybic-sdk-go
```

### 2. Initialization & Configuration

Initialize the client in your Go application. You can pass `nil` to the `NewClient` function to configure the client from environment variables, or provide a `lybic.Config` struct for programmatic configuration.

#### Basic Initialization

This example initializes the client using environment variables.

```go
package main

import (
	"context"
	"fmt"

	"github.com/lybic/lybic-sdk-go"
)

func main() {
	// Passing nil initializes a client with environment variables
	client, err := lybic.NewClient(nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	sandboxes, err := client.ListSandboxes(ctx)
	if err != nil {
		fmt.Printf("Error listing sandboxes: %v", err)
		return
	}
	fmt.Printf("sandboxes: %+v", sandboxes)
}
```

#### Programmatic Configuration

You can also configure the client programmatically by creating a `Config` object.

```go
config := lybic.NewConfig() // Initializes with defaults and env variables
config.OrgId = "your-org-id"
config.ApiKey = "your-api-key"
config.Timeout = 20 // seconds

client, err := lybic.NewClient(config)
if err != nil {
    panic(err)
}
```

#### Configuration Options

The client can be configured with the following options, either through the `Config` struct or environment variables:

| Struct Field     | Environment Variable   | Description                                               | Default Value        |
|------------------|------------------------|-----------------------------------------------------------|----------------------|
| `OrgId`          | `LYBIC_ORG_ID`         | **Required**. Your organization ID.                       | `""`                 |
| `ApiKey`         | `LYBIC_API_KEY`        | Your API key for authentication.                          | `""`                 |
| `Endpoint`       | `LYBIC_API_ENDPOINT`   | The API endpoint URL.                                     | `https://api.lybic.cn` |
| `Timeout`        | -                      | HTTP request timeout in seconds.                          | `10`                 |
| `ExtraHeaders`   | -                      | A map of extra HTTP headers to send with each request.    | `nil`                |
| `Logger`         | -                      | A custom logger instance. See the [Logging](#-logging) section. | `nil` (disabled)     |

## ✨ Platform API Features

The Lybic SDK provides a comprehensive client for interacting with all major platform features.

All examples are available in the [examples.md](docs/sdk_example.md)

### Sandbox Management
- `ListSandboxes(ctx)`: Retrieve a list of all available sandboxes.
- `CreateSandbox(ctx, dto)`: Create a new sandbox.
- `GetSandbox(ctx, sandboxId)`: Retrieve details for a specific sandbox.
- `DeleteSandbox(ctx, sandboxId)`: Delete a sandbox.
- `ExtendSandbox(ctx, sandboxId, dto)`: Extend or modify a sandbox.

### Sandbox Interaction
- `ExecuteComputerUseAction(ctx, sandboxId, dto)`: Perform an action (e.g., mouse click, keyboard input) on a sandbox.
- `PreviewSandbox(ctx, sandboxId)`: Generate a preview (screenshot) of the sandbox state.

### Project Management
- `ListProjects(ctx)`: Get a list of all projects.
- `CreateProject(ctx, dto)`: Create a new project.
- `DeleteProject(ctx, projectId)`: Delete a project.

### Other Utilities
- `GetStats(ctx)`: Retrieve current platform statistics.
- `ParseComputerUse(ctx, dto)`: Parse and validate computer use actions.

## 🤖 Using the MCP Client

For interacting with the Model Context Protocol (MCP), which enables tool calling, you need to initialize a separate `McpClient`.

### MCP Client Initialization

The `NewMcpClient` function creates a client for MCP services. It requires an `McpOption` struct, which can be configured in several ways.

#### Example 1: Initialize with a Lybic Config

You can initialize the `McpClient` directly from a `lybic.Config` object. The client will be created internally.

```go
import (
    "context"

    "github.com/lybic/lybic-sdk-go"
)

func main() {
    // Config can be loaded from environment variables
    lybicConfig := lybic.NewConfig()

    mcpOpt := lybic.McpOption{
        UsingClientConfig: lybicConfig,
    }
    mcpClient, err := lybic.NewMcpClient(context.Background(), mcpOpt)
    if err != nil {
        panic(err)
    }
    defer mcpClient.Close()
    
    // ... use the mcpClient
}
```

#### Example 2: Initialize with an Existing Lybic Client

If you already have an instance of the main `lybic.Client`, you can reuse it.

```go
client, err := lybic.NewClient(nil) // Main client
if err != nil {
    panic(err)
}

mcpOpt := lybic.McpOption{
    UsingClient: client,
}
mcpClient, err := lybic.NewMcpClient(context.Background(), mcpOpt)
if err != nil {
    panic(err)
}
defer mcpClient.Close()
```

### MCP Client Features

#### Tool Calling
- `CallTools(ctx, args, service)`: Call a tool service like `computer-use` with the specified arguments.

#### MCP Server Management
- `ListMcpServers(ctx)`: Get a list of all MCP servers.
- `CreateMcpServer(ctx, dto)`: Create a new MCP server.
- `GetDefaultMcpServer(ctx)`: Retrieve the default MCP server.
- `DeleteMcpServer(ctx, mcpServerId)`: Delete an MCP server.
- `SetMcpServerToSandbox(ctx, mcpServerId, dto)`: Associate an MCP server with a sandbox.

## 📝 Logging

The SDK uses a flexible logging interface that allows you to integrate your own preferred logging library. Any logger that implements the `lybic.Logger` interface is supported. The interface is compatible with popular libraries like `zap (SugaredLogger)`, `logrus`, `zerolog`, and `slog`.

The `lybic.Logger` interface is defined as:
```go
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
```

### Example: Using Logrus

```go
import (
    "github.com/lybic/lybic-sdk-go"
    "github.com/sirupsen/logrus"
)

func main() {
    // Create a new logrus logger
    customLogger := logrus.New()
    customLogger.SetLevel(logrus.DebugLevel)

    // Create a new Lybic client configuration
    config := lybic.NewConfig()
    config.OrgId = "your-org-id"
    config.ApiKey = "your-api-key"
    config.Logger = customLogger // Set the custom logger

    // Initialize the client
    client, err := lybic.NewClient(config)
    if err != nil {
        customLogger.Fatalf("Failed to create Lybic client: %v", err)
    }
    
    // ... use the client
}
```

If no logger is provided, logging will be disabled.

## 📚 Full Documentation & API Reference

This README provides a high-level overview. For detailed, up-to-date code examples, tutorials, and a complete API reference, please visit our **[official documentation site](https://lybic.ai/docs)**.

## 🤝 Contributing

We welcome contributions from the community! Please see our [Contributing Guidelines](https://github.com/lybic/lybic-sdk-go/blob/main/CONTRIBUTING.md) for more details on how to get involved.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/lybic/lybic-sdk-go/blob/main/LICENSE) file for details.
