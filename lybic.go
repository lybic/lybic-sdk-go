// Copyright (c) 2019-2025   Beijing Tingyu Technology Co., Ltd.
// Copyright (c) 2025        Lybic Development Team <team@lybic.ai, lybic@tingyutech.com>
// Copyright (c) 2025        Lu Yicheng <luyicheng@tingyutech.com>
//
// These Terms of Service ("Terms") set forth the rules governing your access to and use of the website lybic.ai
// ("Website"), our web applications, and other services (collectively, the "Services") provided by Beijing Tingyu
// Technology Co., Ltd. ("Company," "we," "us," or "our"), a company registered in Haidian District, Beijing. Any
// breach of these Terms may result in the suspension or termination of your access to the Services.
// By accessing and using the Services and/or the Website, you represent that you are at least 18 years old,
// acknowledge that you have read and understood these Terms, and agree to be bound by them. By using or accessing
// the Services and/or the Website, you further represent and warrant that you have the legal capacity and authority
// to agree to these Terms, whether as an individual or on behalf of a company. If you do not agree to all of these
// Terms, do not access or use the Website or Services.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package lybic provides a set of utilities and functions for working with the Lybic API.
//
//	It offers comprehensive client functionality for managing sandboxes, projects, MCP servers,
//	and various other Lybic platform features.
package lybic

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Client defines the interface for interacting with the Lybic API.
// It provides methods for managing sandboxes, projects, MCP servers and other platform resources.
type Client interface {
	// GetConfig returns the current configuration of the client
	GetConfig() *Config

	// ListSandboxes retrieves a list of all available sandboxes
	ListSandboxes(ctx context.Context) ([]CreateSandboxResponseDto, error)

	// CreateSandbox creates a new sandbox with the specified configuration
	CreateSandbox(ctx context.Context, dto CreateSandboxDto) (*CreateSandboxResponseDto, error)

	// GetSandbox retrieves detailed information about a specific sandbox
	GetSandbox(ctx context.Context, sandboxId string) (*GetSandboxResponseDto, error)

	// DeleteSandbox removes a specific sandbox by its ID
	DeleteSandbox(ctx context.Context, sandboxId string) error

	// ExtendSandbox extends the duration or modifies settings of an existing sandbox
	ExtendSandbox(ctx context.Context, sandboxId string, dto ExtendSandboxDto) error

	// ExecuteComputerUseAction performs a specified action on a sandbox
	//  Deprecated: Use ExecuteSandboxAction instead.
	ExecuteComputerUseAction(ctx context.Context, sandboxId string, dto ComputerUseActionDto) (*SandboxActionResponseDto, error)

	// PreviewSandbox generates a preview of the sandbox state
	PreviewSandbox(ctx context.Context, sandboxId string) (*SandboxActionResponseDto, error)

	// ListProjects returns a list of all available projects
	ListProjects(ctx context.Context) ([]SingleProjectResponseDto, error)

	// CreateProject creates a new project with the specified configuration
	CreateProject(ctx context.Context, dto CreateProjectDto) (*SingleProjectResponseDto, error)

	// DeleteProject removes a specific project by its ID
	DeleteProject(ctx context.Context, projectId string) error

	// GetStats retrieves current platform statistics
	GetStats(ctx context.Context) (*StatsResponseDto, error)

	// ParseComputerUse parses and validates computer use actions
	ParseComputerUse(ctx context.Context, model string, dto ParseTextRequestDto) (*ComputerUseActionResponseDto, error)

	// ParseMobileUseModelTextOutput parses and validates mobile use actions from text input
	ParseMobileUseModelTextOutput(ctx context.Context, modelType string, dto ParseTextRequestDto) (*MobileUseActionResponseDto, error)

	// ExecuteSandboxAction performs a generic action on a sandbox
	ExecuteSandboxAction(ctx context.Context, sandboxId string, dto ExecuteSandboxActionDto) (*SandboxActionResponseDto, error)
}

// NewClient creates a new instance of the Lybic client with the provided configuration.
// It returns an error if the configuration is invalid or the client cannot be initialized.
//
//	if config is nil, it initializes a new Config with default values and environment variables.
func NewClient(config *Config) (Client, error) {
	return newClient(config)
}

// Config holds the configuration parameters for the Lybic client.
type Config struct {
	// OrgId is the organization ID required for API access
	OrgId string

	// ApiKey is the authentication key for API access (optional)
	ApiKey string

	// Endpoint is the API endpoint URL, defaults to "https://api.lybic.cn"
	Endpoint string

	// Timeout specifies the duration in seconds for HTTP requests, defaults to 10 seconds
	Timeout uint8

	// ExtraHeaders contains additional HTTP headers to be included in each request
	ExtraHeaders map[string]string

	// Logger provides an interface for logging operations, can be nil to disable logging
	Logger Logger

	// HttpTransport allows customization of the HTTP transport layer, can be nil to use the default transport
	HttpTransport http.RoundTripper
}

// NewConfig creates a new Config instance with default values and environment variables.
// It initializes the configuration with values from environment variables if available,
// otherwise uses default values.
func NewConfig() *Config {
	return &Config{
		OrgId:    getEnv(envOrgId, ""),
		ApiKey:   getEnv(envApiKey, ""),
		Endpoint: getEnv(envEndpoint, defaultEndpoint),
		Timeout:  defaultTimeout,
	}
}

// Mcp defines the interface for interacting with the lybic Model Context Protocol (MCP) services.
type Mcp interface {
	// ListMcpServers retrieves a list of all available MCP servers
	ListMcpServers(ctx context.Context) ([]McpServerResponseDto, error)

	// CreateMcpServer creates a new MCP server with the specified configuration
	CreateMcpServer(ctx context.Context, dto CreateMcpServerDto) (*McpServerResponseDto, error)

	// GetDefaultMcpServer retrieves the default MCP server configuration
	GetDefaultMcpServer(ctx context.Context) (*McpServerResponseDto, error)

	// DeleteMcpServer removes a specific MCP server by its ID
	DeleteMcpServer(ctx context.Context, mcpServerId string) error

	// SetMcpServerToSandbox associates an MCP server with a sandbox
	SetMcpServerToSandbox(ctx context.Context, mcpServerId string, dto SetMcpServerToSandboxResponseDto) error

	// CallTools calls the specified tool service with the given arguments.
	//
	//	args: MCP request content map[string]any
	//	service: "computer-use","mobile-use"
	//	If no service is specified, it defaults to "computer-use".
	CallTools(ctx context.Context, args map[string]any, service *string) (*mcp.CallToolResult, error)

	GetTools(ctx context.Context) ([]*mcp.Tool, error)
	// Close releases any resources held by the MCP client
	Close() error
}

var (
	ErrNeedConfig       = errors.New("please specify a configuration(LybicClient Config) for the MCP client initialization")
	ErrNeedMcpServerId  = errors.New("please specify a MCP server ID when DoNotUsingDefaultServer is true")
	ErrInvalidMcpClient = errors.New("invalid client type: UsingClient must be a client created by this SDK")
)

// NewMcpClient creates a new lybic MCP client with the specified options.
func NewMcpClient(ctx context.Context, opt McpOption) (Mcp, error) {
	if opt.UsingClientConfig == nil && opt.UsingClient == nil {
		return nil, ErrNeedConfig
	}

	var c *client
	var err error
	if opt.UsingClient == nil {
		c, err = newClient(opt.UsingClientConfig)
		if err != nil {
			return nil, err
		}
	} else {
		var ok bool
		c, ok = opt.UsingClient.(*client)
		if !ok {
			return nil, ErrInvalidMcpClient
		}
	}

	var mcpServerAddress *string
	if opt.DoNotUsingDefaultServer != nil && *opt.DoNotUsingDefaultServer {
		if opt.UsingSpecificMcpServerId == nil || strings.TrimSpace(*opt.UsingSpecificMcpServerId) == "" {
			return nil, ErrNeedMcpServerId
		} else {
			mcpServerAddress = opt.UsingSpecificMcpServerId
		}
	}
	return newMcpClient(ctx, c, mcpServerAddress)
}

// McpOption holds options for configuring the lybic MCP client.
type McpOption struct {
	UsingClientConfig *Config
	UsingClient       Client

	// DoNotUsingDefaultServer If this option is specified and is true, UsingSpecificMcpServerId must be specified
	DoNotUsingDefaultServer *bool
	// UsingSpecificMcpServerId If this option is specified, the MCP client will use the specified MCP server ID.
	UsingSpecificMcpServerId *string
}
