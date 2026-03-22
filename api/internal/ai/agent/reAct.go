package react

import (
	models "Echo/api/internal/ai/models/toolCallingChatModel"
	"Echo/api/internal/ai/tools"
	"Echo/api/internal/config"
	"context"
	"errors"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"io"

	"github.com/cloudwego/eino/flow/agent/react"
)

func NewAgent(ctx context.Context, config config.Config) *react.Agent {
	model := models.NewToolCallingChatModel(ctx, config)

	reactConfig := &react.AgentConfig{
		ToolCallingModel:   model,
		MaxStep:            25,
		ToolReturnDirectly: map[string]struct{}{},
		ToolsConfig: compose.ToolsNodeConfig{
			ExecuteSequentially: false,
		},
		StreamToolCallChecker: toolCallChecker,
	}
	reactConfig.ToolsConfig.Tools = []tool.BaseTool{}
	reactConfig.ToolsConfig.Tools = append(reactConfig.ToolsConfig.Tools, tools.NewQueryInternalDocsTool())
	reactConfig.ToolsConfig.Tools = append(reactConfig.ToolsConfig.Tools, tools.NewGetCurrentTimeTool())
	reactConfig.ToolsConfig.Tools = append(reactConfig.ToolsConfig.Tools, tools.NewModifyMdContentTool())

	agent, err := react.NewAgent(ctx, reactConfig)
	if err != nil {
		panic(err)
	}
	return agent
}

// toolCallChecker 用于检查从流中读取的消息是否包含工具调用。
// 它会持续从流中接收消息，直到找到一个包含工具调用的消息或流结束。
func toolCallChecker(ctx context.Context, sr *schema.StreamReader[*schema.Message]) (bool, error) {
	// 确保在函数退出时关闭流读取器，以防资源泄漏。
	defer sr.Close()
	// 无限循环，用于持续从流中读取消息。
	for {
		// 从流中接收一条消息。如果流中没有新消息，此调用会阻塞。
		msg, err := sr.Recv()
		// 检查接收过程中是否发生错误。
		if err != nil {
			// 如果错误是 io.EOF，表示流已正常结束，没有更多消息了。
			if errors.Is(err, io.EOF) {
				// 正常结束，跳出循环。
				break
			}

			// 如果是其他类型的错误，则直接返回错误。
			return false, err
		}

		// 检查收到的消息是否包含任何工具调用。
		if len(msg.ToolCalls) > 0 {
			// 如果找到工具调用，立即返回 true，表示检查成功。
			return true, nil
		}
	}
	// 如果完整遍历了流而没有找到任何工具调用，则返回 false。
	return false, nil
}
