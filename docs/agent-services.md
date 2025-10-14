# 🤖 Agent Service

The Lybic Go SDK provides a powerful client for interacting with the Lybic Agent Cluster, enabling you to programmatically control and automate your GUI-based AI agents. This document provides a detailed guide on how to use the Agent services, with code examples for each feature.

## 🚀 Getting Started

### 1. Installation

First, ensure you have the Lybic Go SDK installed:

```bash
go get github.com/lybic/lybic-sdk-go
```

### 2. Initialization & Configuration

To use the Agent services, you need to initialize an `AgentClient`. The client can be configured with the address of the Agent Cluster.

#### Basic Initialization

This example initializes the client with a specific address.

```go
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/lybic/lybic-sdk-go/pkg/agent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Initialize the Agent client
	var creds credentials.TransportCredentials
	//  if you use TLS, use the following line
	{
		certPool, err := x509.SystemCertPool()
		if err != nil {
			panic(err)
		}

		tlsConfig := &tls.Config{
			RootCAs: certPool,
		}
		creds = credentials.NewTLS(tlsConfig)
	} 
	// if you don't use TLS, use the following line instead
	{
		creds = insecure.NewCredentials()
	}

	conn, err := grpc.NewClient("your-agent-cluster-address:port",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := agent.NewAgentClient(conn)
	ctx := context.Background()

	// Get Agent Info
	agentInfo, err := cli.GetAgentInfo(ctx, &agent.GetAgentInfoRequest{})
	if err != nil {
		fmt.Printf("Error getting agent info: %v", err)
		return
	}
	fmt.Printf("%s\n", agentInfo)
}
```

## ✨ Agent Service Features

The Agent service provides a range of features for managing and interacting with your AI agents.

### Agent Management
- `GetAgentInfo(ctx, req)`: Retrieve information about the agent.
- `GetGlobalCommonConfig(ctx, req)`: Get the global common configuration for the agent.
- `GetCommonConfig(ctx, req)`: Get the common configuration for a specific agent.
- `SetGlobalCommonConfig(ctx, req)`: Set the global common configuration for the agent.
- `SetGlobalCommonLLMConfig(ctx, req)`: Set the global common LLM configuration for the agent.
- `SetGlobalGroundingLLMConfig(ctx, req)`: Set the global grounding LLM configuration for the agent.
- `SetGlobalEmbeddingLLMConfig(ctx, req)`: Set the global embedding LLM configuration for the agent.

### Agent Interaction
- `RunAgentInstruction(ctx, req)`: Run an agent instruction and stream the task results.
- `RunAgentInstructionAsync(ctx, req)`: Run an agent instruction asynchronously.
- `GetAgentTaskStream(ctx, req)`: Get the task stream for a specific task.
- `QueryTaskStatus(ctx, req)`: Query the status of a task.

## 📝 Code Examples

### Get Agent Info

```go
agentInfo, err := cli.GetAgentInfo(context.Background(), &agent.GetAgentInfoRequest{})
if err != nil {
    fmt.Printf("Error getting agent info: %v", err)
    return
}
fmt.Printf("Agent Info: %+v", agentInfo)
```

### Run Agent Instruction (Sync)

```go
stream, err := cli.RunAgentInstruction(context.Background(), &agent.RunAgentInstructionRequest{
    Instruction: "your-instruction",
    // Optionally specify a sandbox and running config
})
if err != nil {
    fmt.Printf("Error running agent instruction: %v", err)
    return
}

for {
    taskStream, err := stream.Recv()
    if err == io.EOF {
        break // due to stream close: task completed or failed
    }
    if err != nil {
        fmt.Printf("Error receiving task stream: %v", err)
        break
    }
    fmt.Printf("Task Stream: %+v", taskStream)
}
```

### Run Agent Instruction (Async)

```go
resp, err := cli.RunAgentInstructionAsync(context.Background(), &agent.RunAgentInstructionRequest{
    Instruction: "your-instruction",
})
if err != nil {
    fmt.Printf("Error running agent instruction async: %v", err)
    return
}
fmt.Printf("Task ID: %s", resp.TaskId)
```

### Query Task Status

```go
status, err := cli.QueryTaskStatus(context.Background(), &agent.QueryTaskStatusRequest{
    TaskId: "your-task-id",
})
if err != nil {
    fmt.Printf("Error querying task status: %v", err)
    return
}
fmt.Printf("Task Status: %+v", status)
```


## 📚 Additional Resources

Deploy Agent Cluster: [Link to Deployment Guide](https://hub.docker.com/r/agenticlybic/guiagent/)