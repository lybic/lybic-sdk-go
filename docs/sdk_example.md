# Lybic Go SDK: API Usage Guide

This document provides a comprehensive guide with examples for using the Lybic Go SDK to interact with the Lybic platform.

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
	"fmt"
	"github.com/lybic/lybic-sdk-go"
)

func main() {
	// Passing nil initializes a client with environment variables
	// Make sure LYBIC_ORG_ID and LYBIC_API_KEY are set.
	client, err := lybic.NewClient(nil)
	if err != nil {
		panic(err)
	}

	// Use the client...
	fmt.Println("Lybic client initialized successfully!")
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

---

## ✨ Platform API Examples

Here are examples of how to use the SDK to interact with various Lybic platform features.

### Sandbox Management

Sandboxes are on-demand, isolated environments for running your GUI-based agents.

#### List Sandboxes

Retrieve a list of all your available sandboxes.

```go
ctx := context.Background()
sandboxes, err := client.ListSandboxes(ctx)
if err != nil {
    fmt.Printf("Error listing sandboxes: %v", err)
    return
}
fmt.Printf("Found %d sandboxes.\n", len(sandboxes))
for _, sandbox := range sandboxes {
    fmt.Printf("- Sandbox ID: %s, Name: %s\n", sandbox.Id, sandbox.Name)
}
```

#### Create a Sandbox

Create a new sandbox. Not all parameters need to be filled in. Parameters marked with * are optional.

```go
_sandboxName := "Test Sandbox"
createDto := lybic.CreateSandboxDto{
    Name: &_sandboxName,
}

sandbox, err := client.CreateSandbox(ctx, createDto)
if err != nil {
    fmt.Printf("Error creating sandbox: %v", err)
    return
}
fmt.Printf("Created sandbox: %+v\n", sandbox)
```

#### Get Sandbox Details

Retrieve detailed information for a specific sandbox.

```go
sandbox, err := client.GetSandbox(ctx, "sandbox-id")
if err != nil {
    fmt.Printf("Error getting sandbox: %v\n", err)
    return
}
fmt.Printf("Sandbox info: %+v\n", sandbox)
```

#### Interact with a Sandbox

You can perform actions like mouse clicks or keyboard inputs, and take screenshots (previews).

<!-- todo: The serializable structure is not ready yet and needs to wait for the DTO component to be updated. Please visit the official website to view the currently supported parameters. --->

**Execute an Action:**
```go
actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: lybic.ComputerUseActionDtoActionOneOf{
        ComputerUseActionDtoActionOneOfInterface: map[string]any{
            "type": "mouse:click",
            "x": map[string]any{
                "type":  "px",
                "value": 10,
            },
            "y": map[string]any{
                "type":  "px",
                "value": 10,
            },
            "button": 1,
        },
    },
})
if err != nil {
    fmt.Println("Error executing computer use action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Get a Preview (Screenshot):**
```go
previewSandbox, err := client.PreviewSandbox(ctx, "sandbox-Id")
if err != nil {
    fmt.Printf("Error previewing sandbox: %v\n", err)
    return
}
fmt.Printf("Previewed sandbox: %+v\n", previewSandbox)
```

#### Extend a Sandbox

Extend the lifetime of a running sandbox.

```go
var _maxLife float32 = 86400 // 24 hours in seconds
err = client.ExtendSandbox(ctx, "sandbox_id", lybic.ExtendSandboxDto{
    MaxLifeSeconds: &_maxLife,
})
if err != nil {
    fmt.Printf("Error extending sandbox: %v\n", err)
    return
}
fmt.Println("Sandbox extended successfully.")
```

#### Delete a Sandbox

Delete a sandbox when you are finished with it.

```go
sandboxId := "your-sandbox-id"
err := client.DeleteSandbox(context.Background(), sandboxId)
if err != nil {
    fmt.Printf("Error deleting sandbox: %v", err)
    return
}
fmt.Println("Sandbox deleted successfully.")
```

### Project Management

Organize your work into projects.

#### List Projects

```go
projects, err := client.ListProjects(ctx)
if err != nil {
    fmt.Println("Error listing projects:", err)
    return
}
for _, project := range projects {
    fmt.Printf("Project ID: %s, Name: %s\n", project.Id, project.Name)
}
```

#### Create a Project

```go
projectDto := lybic.CreateProjectDto{
    Name: "My New AI Agent",
}
project, err := client.CreateProject(ctx, projectDto)
if err != nil {
    fmt.Println("Error creating project:", err)
	return
}
fmt.Println("Project created successfully:", project.Name)
```

#### Delete a Project

```go
err = client.DeleteProject(ctx, "project-id")
if err != nil {
    fmt.Println("Failed to delete project:", err)
    return
}
fmt.Println("Project deleted successfully")
```

### Organization Stats

Retrieve current usage statistics for your organization.

```go
stats, err := client.GetStats(ctx)
if err != nil {
    fmt.Println("Error getting stats:", err)
    return
}
fmt.Printf("Platform Statistics: %+v\n", stats)
```

---

## 🤖 Using the MCP Client

For advanced agent development involving tool calling, you need to use the Model Context Protocol (MCP) client.

### MCP Client Initialization

You can initialize the `McpClient` from an existing `lybic.Client` or a `lybic.Config` object.

```go
// Assuming 'client' is your initialized lybic.Client
mcpClient, err := lybic.NewMcpClient(ctx, lybic.McpOption{
    UsingClient: client,
})
if err != nil {
    panic(err)
}
defer mcpClient.Close()

// Use the mcpClient...
fmt.Println("MCP client initialized successfully!")
```

### Calling Tools

The primary function of the MCP client is to call tools, such as the `computer-use` service.

```go
// This example assumes you have an MCP server associated with a sandbox.
args := map[string]any{
    "action": "doubleClick",
    "x":      120,
    "y":      240,
}
service := "computer-use"

result, err := mcpClient.CallTools(context.Background(), args, &service)
if err != nil {
    fmt.Printf("Error calling tool: %v", err)
    return
}
fmt.Printf("Tool call result: %+v\n", result)
```

### MCP Server Management

You can also manage the MCP servers themselves.

#### List MCP Servers
```go
servers, err := mcpClient.ListMcpServers(ctx)
if err != nil {
    fmt.Println("Error listing MCP servers:", err)
    return
}
fmt.Println("MCP Servers:")
for _, server := range servers {
    fmt.Printf("ID: %s, Name: %s\n", server.Id, server.Name)
}
```

#### Create an MCP Server
```go
m, err := mcpClient.CreateMcpServer(ctx, lybic.CreateMcpServerDto{
    Name: "MCP-server-01",
})
if err != nil {
    fmt.Println("Error creating MCP server:", err)
    return
}
fmt.Println("Created MCP Server:")
fmt.Printf("ID: %s, Name: %s\n", m.Id, m.Name)
```
