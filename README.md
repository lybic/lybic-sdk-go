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
</p>

Developing, testing, and deploying GUI-based AI agents is complex. Developers waste precious time wrestling with cloud instances, VNC servers, and environment configurations instead of focusing on what matters: building intelligent agents.

**Lybic is the infrastructure layer for your GUI agents.**

**Lybic** (/ˈlaɪbɪk/) provides a robust, on-demand infrastructure platform designed specifically for the AI agent development lifecycle. This SDK for Python is your command center for programmatically controlling the entire Lybic ecosystem, empowering you to build, test, and scale your agents with unprecedented speed and simplicity.

## 🚀 Getting Started

### 1. Installation & Setup

Getting started is simple. First, install the package from the Go package repository:

```bash
go get github.com/lybic/lybic-sdk-go
```

Then, initialize the client in your go application:

```go
package main

import (
	"context"
	"fmt"

	"github.com/lybic/lybic-sdk-go"
)

func main() {
	client, err := lybic.NewClient(nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	sandboxes, err := client.ListSandboxes(ctx)
	if err != nil {
		fmt.Printf("Error listing sandboxes: %v\n", err)
		return
	}
	fmt.Printf("sandboxes: %+v\n", sandboxes)
}
```

### 2. Core Workflow

With the client initialized, the typical workflow follows these logical steps:

1. **Register(Or be invited into) an Organization**: Lybic allows you to register a new organization to manage your projects and resources.

2. **Create a `Project`**: Projects are the primary way to organize your work. They act as containers for your sandboxes, team members, and secrets.

3. **Launch a `Sandbox`**: Within a project, you can launch a GUI sandbox. This is your agent's secure, cloud-based home.

4. **Automate and Interact**: Once the sandbox is running, your agent can begin its work. The SDK provides all the necessary tools to interact with the sandbox, from executing commands to capturing screenshots.

## 📚 Full Documentation & API Reference

This README provides a high-level overview of Lybic's capabilities. For detailed, up-to-date code examples, tutorials, and a complete API reference, please visit our **[official documentation site](https://lybic.ai/docs)**.

## 🤝 Contributing

We welcome contributions from the community! Please see our [Contributing Guidelines](https://github.com/lybic/lybic-sdk-go/blob/main/CONTRIBUTING.md) for more details on how to get involved.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/lybic/lybic-sdk-go/blob/main/LICENSE) file for details.
