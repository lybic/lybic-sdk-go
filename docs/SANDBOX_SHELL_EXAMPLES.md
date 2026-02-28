# Sandbox Shell Command Examples

This document provides examples of how to use the new Sandbox Shell Command APIs.

## Basic Shell Command Execution

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/lybic/lybic-sdk-go"
)

func main() {
    client, err := lybic.NewClient(nil)
    if err != nil {
        log.Fatal(err)
    }
    
    ctx := context.Background()
    sandboxId := "your-sandbox-id"
    
    // Create a shell session
    createReq := lybic.SandboxShellCommandCreateRequestDto{
        Command: "ls -la",
        UseTty: false,
    }
    
    session, err := client.CreateSandboxShellCommand(ctx, sandboxId, createReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Shell session created: %s\n", session.SessionId)
    
    // Read the output
    output, err := client.ReadSandboxShellCommand(ctx, sandboxId, session.SessionId)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Is running: %v\n", output.IsRunning)
    for _, out := range output.Output {
        switch out.OneofKind {
        case "stdout":
            if out.Stdout != nil {
                fmt.Printf("STDOUT: %s\n", *out.Stdout)
            }
        case "stderr":
            if out.Stderr != nil {
                fmt.Printf("STDERR: %s\n", *out.Stderr)
            }
        case "waiting":
            fmt.Println("Command completed, waiting to exit")
        }
    }
    
    // Terminate the session
    err = client.TerminateSandboxShellCommand(ctx, sandboxId, session.SessionId)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Interactive Shell with TTY

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/lybic/lybic-sdk-go"
)

func main() {
    client, err := lybic.NewClient(nil)
    if err != nil {
        log.Fatal(err)
    }
    
    ctx := context.Background()
    sandboxId := "your-sandbox-id"
    
    timeout := int32(3600)
    rows := int32(24)
    cols := int32(80)
    
    // Create an interactive shell session with TTY
    createReq := lybic.SandboxShellCommandCreateRequestDto{
        Command: "/bin/bash",
        UseTty: true,
        TimeoutSeconds: &timeout,
        TtyRows: &rows,
        TtyCols: &cols,
    }
    
    session, err := client.CreateSandboxShellCommand(ctx, sandboxId, createReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Interactive shell session created: %s\n", session.SessionId)
    
    // Write a command to the shell
    writeReq := lybic.SandboxShellCommandWriteRequestDto{
        Data: "echo 'Hello from shell'\n",
    }
    
    err = client.WriteSandboxShellCommand(ctx, sandboxId, session.SessionId, writeReq)
    if err != nil {
        log.Fatal(err)
    }
    
    // Read the output
    output, err := client.ReadSandboxShellCommand(ctx, sandboxId, session.SessionId)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, out := range output.Output {
        if out.OneofKind == "stdout" && out.Stdout != nil {
            fmt.Printf("Output: %s\n", *out.Stdout)
        }
    }
    
    // Finish writing (send EOF)
    err = client.FinishSandboxShellCommand(ctx, sandboxId, session.SessionId)
    if err != nil {
        log.Fatal(err)
    }
    
    // Terminate the session
    err = client.TerminateSandboxShellCommand(ctx, sandboxId, session.SessionId)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Streaming Shell Command (SSE)

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/lybic/lybic-sdk-go"
)

func main() {
    client, err := lybic.NewClient(nil)
    if err != nil {
        log.Fatal(err)
    }
    
    ctx := context.Background()
    sandboxId := "your-sandbox-id"
    
    timeout := int32(300)
    workDir := "/home/user"
    
    // Create a streaming shell session
    createReq := lybic.SandboxShellCommandStreamCreateRequestDto{
        Command: "for i in {1..5}; do echo \"Line $i\"; sleep 1; done",
        UseTty: false,
        TimeoutSeconds: &timeout,
        WorkingDirectory: &workDir,
    }
    
    eventChan, err := client.CreateSandboxShellCommandStream(ctx, sandboxId, createReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Streaming shell output:")
    
    // Process events from the stream
    for event := range eventChan {
        switch event.Type {
        case "stdout":
            fmt.Printf("STDOUT: %s", event.Data)
        case "stderr":
            fmt.Printf("STDERR: %s", event.Data)
        case "waiting":
            fmt.Println("Command completed")
        case "timeout":
            fmt.Printf("Command timed out: %s\n", event.Data)
        case "end":
            fmt.Println("Stream ended")
        }
    }
}
```

## API Overview

### Non-Streaming APIs

1. **CreateSandboxShellCommand**: Creates a shell session and returns a session ID
2. **WriteSandboxShellCommand**: Writes input data to a shell session
3. **FinishSandboxShellCommand**: Signals that no more input will be sent (sends EOF)
4. **ReadSandboxShellCommand**: Reads accumulated output from the shell session
5. **TerminateSandboxShellCommand**: Terminates a shell session

### Streaming API (SSE)

- **CreateSandboxShellCommandStream**: Creates a shell session with real-time streaming output via Server-Sent Events (SSE). Returns a channel that emits events as they occur.

## Event Types in SSE Stream

- `stdout`: Standard output from the command
- `stderr`: Standard error output from the command
- `waiting`: Command completed and is waiting to exit
- `timeout`: Command execution timed out
- `end`: Stream has ended

## Notes

- All data in SSE events is base64-encoded on the server side and automatically decoded by the SDK
- The streaming API is useful for long-running commands where you want real-time feedback
- The non-streaming APIs are better for short commands or when you need more control over the interaction
- Use TTY mode for interactive shells that require terminal features
- Always call `TerminateSandboxShellCommand` to clean up shell sessions (non-streaming only)
