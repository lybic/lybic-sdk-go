package covert

import (
	anthropic "github.com/anthropics/anthropic-sdk-go"
	param2 "github.com/anthropics/anthropic-sdk-go/packages/param"
	"github.com/lybic/lybic-sdk-go/pkg/json"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/shared"
	saopenai "github.com/sashabaranov/go-openai"
)

// McpTools2SaOpenAiTools converts a list of MCP tools to a list of OpenAI tools.
//
//	It maps the relevant fields from MCP Tool to OpenAI Tool format.
//	SDK: github.com/sashabaranov/go-openai
func McpTools2SaOpenAiTools(tools []*mcp.Tool) (retval []saopenai.Tool) {
	for _, mtool := range tools {
		openaitool := saopenai.Tool{
			Type: saopenai.ToolTypeFunction,
			Function: &saopenai.FunctionDefinition{
				Name:        mtool.Name,
				Description: mtool.Description,
				Parameters:  mtool.InputSchema,
			},
		}
		retval = append(retval, openaitool)
	}
	return
}

// McpTools2OpenAiTools converts a list of MCP tools to a list of OpenAI tools.
//
//	It maps the relevant fields from MCP Tool to OpenAI Tool format.
//	SDK: github.com/openai/openai-go
func McpTools2OpenAiTools(tools []*mcp.Tool) (retval []openai.ChatCompletionToolParam, err error) {
	for _, mtool := range tools {
		params := make(map[string]any)
		if mtool.InputSchema != nil {
			b, err := json.Marshal(mtool.InputSchema)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(b, &params)
			if err != nil {
				return nil, err
			}
		}
		openaitool := openai.ChatCompletionToolParam{
			Function: shared.FunctionDefinitionParam{
				Name:        mtool.Name,
				Description: param.Opt[string]{Value: mtool.Description},
				Parameters:  params,
			},
		}
		retval = append(retval, openaitool)
	}
	return
}

// McpTools2AnthropicTools converts a list of MCP tools to a list of Anthropic tools.
//
//	It maps the relevant fields from MCP Tool to Anthropic Tool format.
//	SDK: github.com/anthropics/anthropic-sdk-go
func McpTools2AnthropicTools(tools []*mcp.Tool) (retval []anthropic.ToolParam, err error) {
	for _, mtool := range tools {
		var schema anthropic.ToolInputSchemaParam
		if mtool.InputSchema != nil {
			b, err := json.Marshal(mtool.InputSchema)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(b, &schema)
			if err != nil {
				return nil, err
			}
		}
		tool := anthropic.ToolParam{
			Name:        mtool.Name,
			Description: param2.Opt[string]{Value: mtool.Description},
			InputSchema: schema,
		}
		retval = append(retval, tool)
	}
	return
}
