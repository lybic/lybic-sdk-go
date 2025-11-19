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

You can perform actions like mouse clicks or keyboard inputs, and take screenshots. The following examples demonstrate how to use various actions available in the SDK.

**Execute a Mouse Click Action:**
```go
// Create a mouse click action at position (100, 200) with the left mouse button.
action := lybic.NewMouseClickAction(
    lybic.NewPixelLength(100),
    lybic.NewPixelLength(200),
    1, // 1 for left button
)

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing mouse click action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Execute a Mouse Double-Click Action:**
```go
// Create a mouse double-click action at the center of the screen.
action := lybic.NewMouseDoubleClickAction(
    lybic.NewFractionalLength(1, 2), // 1/2 of screen width
    lybic.NewFractionalLength(1, 2), // 1/2 of screen height
    1, // 1 for left button
)

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing mouse double-click action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Execute a Mouse Move Action:**
```go
// Move the mouse to position (300, 400).
action := lybic.NewMouseMoveAction(
    lybic.NewPixelLength(300),
    lybic.NewPixelLength(400),
)

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing mouse move action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Execute a Mouse Scroll Action:**
```go
// Scroll vertically by 10 steps at position (300, 400).
action := lybic.NewMouseScrollAction(
    lybic.NewPixelLength(300), // x position
    lybic.NewPixelLength(400), // y position
    10,  // vertical scroll steps
    0,   // horizontal scroll steps
)

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing mouse scroll action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Execute a Mouse Drag Action:**
```go
// Drag the mouse from (100, 100) to (500, 500).
action := lybic.NewMouseDragAction(
    lybic.NewPixelLength(100),
    lybic.NewPixelLength(100),
    lybic.NewPixelLength(500),
    lybic.NewPixelLength(500),
)

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing mouse drag action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Execute a Keyboard Type Action:**
```go
// Create a keyboard type action to type "Hello, Lybic!".
action := lybic.NewKeyboardTypeAction("Hello, Lybic!", false)

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing keyboard type action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Execute a Keyboard Hotkey Action:**
```go
// Press Ctrl+C.
action := lybic.NewKeyboardHotkeyAction("ctrl+c")

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing keyboard hotkey action:", err.Error())
    return
}
fmt.Println("Action executed successfully:", actionResult)
```

**Take a Screenshot:**
```go
// Take a screenshot of the sandbox.
action := lybic.NewScreenshotAction()

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error taking screenshot:", err.Error())
    return
}
// The screenshot data will be in actionResult.
fmt.Println("Screenshot taken successfully:", actionResult)
```

**Get a Preview (Convenience Method for Screenshot):**
```go
// This is a helper function that wraps the screenshot action.
previewSandbox, err := client.PreviewSandbox(ctx, "sandbox-Id")
if err != nil {
    fmt.Printf("Error previewing sandbox: %v\n", err)
    return
}
fmt.Printf("Previewed sandbox: %+v", previewSandbox)
```

**Execute a Wait Action:**
```go
// Wait for 5 seconds (5000 milliseconds).
action := lybic.NewWaitAction(5000)

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error executing wait action:", err.Error())
    return
}
fmt.Println("Wait action completed successfully:", actionResult)
```

**Signal Task Finished:**
```go
// Signal that the task is finished successfully.
action := lybic.NewFinishedAction()
// Optionally, you can add a message.
// action.Message = "Task completed with flying colors!"

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error sending finished signal:", err.Error())
    return
}
fmt.Println("Finished signal sent successfully:", actionResult)
```

**Signal Task Failed:**
```go
// Signal that the task has failed.
action := lybic.NewFailedAction()
// Optionally, you can add a message.
// action.Message = "Something went wrong."

actionResult, err := client.ExecuteComputerUseAction(ctx, "sandbox-Id", lybic.ComputerUseActionDto{
    Action: action,
})
if err != nil {
    fmt.Println("Error sending failed signal:", err.Error())
    return
}
fmt.Println("Failed signal sent successfully:", actionResult)
```

#### Extend a Sandbox

Extend the lifetime of a running sandbox.

```go
var maxLife float32 = 86400 // 24 hours in seconds
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

#### Copy Files With Sandbox

Transfer files to/from the sandbox using multiple protocols (sandbox file paths, HTTP PUT/GET, HTTP POST forms).

**Example 1: Copy a file from sandbox to HTTP endpoint**

```go
import "encoding/base64"

// Define the source (sandbox file) and destination (HTTP PUT upload)
dto := lybic.SandboxFileCopyRequestDto{
    Files: []lybic.SandboxFileCopyRequestDtoFiles{
        {
            Id: "file-1", // Optional ID to track this operation in the response
            Src: map[string]interface{}{
                "type": "sandboxFileLocation",
                "path": "/home/user/output.txt",
            },
            Dest: map[string]interface{}{
                "type": "httpPutLocation",
                "url":  "https://storage.example.com/upload",
                "headers": map[string]string{
                    "Authorization": "Bearer your-token",
                },
            },
        },
    },
}

resp, err := client.CopyFilesWithSandbox(ctx, "sandbox-id", dto)
if err != nil {
    fmt.Printf("Error copying files: %v\n", err)
    return
}

// Check results
for _, result := range resp.Results {
    if result.Success {
        fmt.Printf("File %s copied successfully\n", result.Id)
    } else {
        fmt.Printf("File %s failed: %s\n", result.Id, result.Error)
    }
}
```

**Example 2: Upload a file from HTTP to sandbox**

```go
dto := lybic.SandboxFileCopyRequestDto{
    Files: []lybic.SandboxFileCopyRequestDtoFiles{
        {
            Id: "download-1",
            Src: map[string]interface{}{
                "type": "httpGetLocation",
                "url":  "https://example.com/data.csv",
            },
            Dest: map[string]interface{}{
                "type": "sandboxFileLocation",
                "path": "/home/user/data.csv",
            },
        },
    },
}

resp, err := client.CopyFilesWithSandbox(ctx, "sandbox-id", dto)
if err != nil {
    fmt.Printf("Error downloading file to sandbox: %v\n", err)
    return
}
fmt.Println("File downloaded to sandbox successfully")
```

**Example 3: Copy multiple files in one request**

```go
dto := lybic.SandboxFileCopyRequestDto{
    Files: []lybic.SandboxFileCopyRequestDtoFiles{
        {
            Id: "file-1",
            Src: map[string]interface{}{
                "type": "sandboxFileLocation",
                "path": "/home/user/report.pdf",
            },
            Dest: map[string]interface{}{
                "type": "httpPutLocation",
                "url":  "https://storage.example.com/reports/report.pdf",
            },
        },
        {
            Id: "file-2",
            Src: map[string]interface{}{
                "type": "sandboxFileLocation",
                "path": "/home/user/data.json",
            },
            Dest: map[string]interface{}{
                "type": "httpPutLocation",
                "url":  "https://storage.example.com/data/data.json",
            },
        },
    },
}

resp, err := client.CopyFilesWithSandbox(ctx, "sandbox-id", dto)
if err != nil {
    fmt.Printf("Error copying files: %v\n", err)
    return
}

// Check which files succeeded
for _, result := range resp.Results {
    fmt.Printf("File %s: success=%v\n", result.Id, result.Success)
}
```

#### Execute Process in Sandbox

Run commands and scripts inside the sandbox, with support for arguments, working directory, stdin, and capturing stdout/stderr.

**Example 1: Execute a simple command**

```go
processDto := lybic.SandboxProcessRequestDto{
    Executable: "python3",
    Args:       []string{"-c", "print('Hello from sandbox!')"},
}

result, err := client.ExecSandboxProcess(ctx, "sandbox-id", processDto)
if err != nil {
    fmt.Printf("Error executing process: %v\n", err)
    return
}

// Decode and print stdout
stdoutBytes, _ := base64.StdEncoding.DecodeString(result.StdoutBase64)
fmt.Printf("Output: %s\n", string(stdoutBytes))
fmt.Printf("Exit code: %d\n", result.ExitCode)
```

**Example 2: Execute with working directory and arguments**

```go
processDto := lybic.SandboxProcessRequestDto{
    Executable:       "/usr/bin/ls",
    Args:             []string{"-la"},
    WorkingDirectory: "/home/user",
}

result, err := client.ExecSandboxProcess(ctx, "sandbox-id", processDto)
if err != nil {
    fmt.Printf("Error executing ls: %v\n", err)
    return
}

// Decode stdout to see the directory listing
stdoutBytes, _ := base64.StdEncoding.DecodeString(result.StdoutBase64)
fmt.Printf("Directory listing:\n%s\n", string(stdoutBytes))
```

**Example 3: Execute a script with stdin input**

```go
import "encoding/base64"

// Prepare stdin data
stdinData := "line1\nline2\nline3"
stdinBase64 := base64.StdEncoding.EncodeToString([]byte(stdinData))

processDto := lybic.SandboxProcessRequestDto{
    Executable:       "python3",
    Args:             []string{"-c", "import sys; print(f'Read {len(sys.stdin.readlines())} lines')"},
    WorkingDirectory: "/tmp",
    StdinBase64:      stdinBase64,
}

result, err := client.ExecSandboxProcess(ctx, "sandbox-id", processDto)
if err != nil {
    fmt.Printf("Error executing process: %v\n", err)
    return
}

stdoutBytes, _ := base64.StdEncoding.DecodeString(result.StdoutBase64)
fmt.Printf("Process output: %s\n", string(stdoutBytes))
fmt.Printf("Exit code: %d\n", result.ExitCode)
```

**Example 4: Handle errors from process execution**

```go
processDto := lybic.SandboxProcessRequestDto{
    Executable: "false", // Command that always fails with exit code 1
}

result, err := client.ExecSandboxProcess(ctx, "sandbox-id", processDto)
if err != nil {
    fmt.Printf("Error calling API: %v\n", err)
    return
}

// Check exit code to determine success/failure
if result.ExitCode != 0 {
    stderrBytes, _ := base64.StdEncoding.DecodeString(result.StderrBase64)
    fmt.Printf("Process failed with exit code %d\n", result.ExitCode)
    fmt.Printf("Error output: %s\n", string(stderrBytes))
} else {
    fmt.Println("Process executed successfully")
}
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
