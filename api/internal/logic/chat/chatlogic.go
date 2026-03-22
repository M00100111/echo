// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package chat

import (
	react "Echo/api/internal/ai/agent"
	"Echo/api/internal/ai/callback"
	models2 "Echo/api/internal/ai/models/embeddingModel"
	"Echo/api/internal/ai/retriever"
	"Echo/api/internal/config"
	"Echo/pkg/mem"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"io"
	"log"

	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"

	"Echo/api/internal/svc"
	"Echo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.ChatReq, client chan<- *types.ChatResp) error {
	chatId := req.ChatId
	if chatId == "" {
		id, _ := uuid.NewUUID()
		chatId = id.String()
	}

	//var documents []*schema.Document
	//if req.Message != "" {
	//	var err error
	//	documents, err = l.queryInternalDocsTool(l.ctx, l.svcCtx.Config, req.Message)
	//	if err != nil {
	//		return err
	//	}
	//}

	//	serverTools := []*agenticark.ServerToolConfig{
	//		{
	//			WebSearch: &responses.ToolWebSearch{
	//				Type: responses.ToolType_web_search,
	//			},
	//		},
	//	}
	//
	//	allowedTools := []*schema.AllowedTool{
	//		{
	//			ServerTool: &schema.AllowedServerTool{
	//				Name: string(agenticark.ServerToolNameWebSearch),
	//			},
	//		},
	//	}
	//
	//	opts := []model.Option{
	//		agenticark.WithServerTools(serverTools),
	//		model.WithAgenticToolChoice(&schema.AgenticToolChoice{
	//			Type: schema.ToolChoiceForced,
	//			Forced: &schema.AgenticForcedToolChoice{
	//				Tools: allowedTools,
	//			},
	//		}),
	//		agenticark.WithThinking(&responses.ResponsesThinking{
	//			Type: responses.ThinkingType_enabled.Enum(),
	//		}),
	//	}

	// 添加检索到的文档作为系统消息
	//if len(documents) > 0 {
	//	docContent := "相关文档信息：\n"
	//	for _, doc := range documents {
	//		docContent += "- " + doc.Content + "\n"
	//	}
	//	input = append(input, schema.SystemAgenticMessage(docContent))
	//}
	//input := []*schema.AgenticMessage{}
	// 获取历史对话
	memory := mem.GetSimpleMemory(chatId)
	historyMessages := memory.GetMessages()
	//// 添加历史对话
	//for _, msg := range historyMessages {
	//	// 根据消息类型转换为AgenticMessage
	//	switch msg.Role {
	//	case schema.User:
	//		input = append(input, schema.UserAgenticMessage(msg.Content))
	//	case schema.Assistant:
	//		temp := schema.UserAgenticMessage(msg.Content)
	//		temp.Role = schema.AgenticRoleTypeAssistant
	//		input = append(input, temp)
	//	}
	//}

	userMessage := req.Message
	//input = append(input, schema.UserAgenticMessage(userMessage))

	// 添加用户消息
	historyMessages = append(historyMessages, &schema.Message{
		Role:    schema.User,
		Content: userMessage,
	})

	reactAgent := react.NewAgent(l.ctx, l.svcCtx.Config)
	resp, err := reactAgent.Stream(l.ctx, historyMessages,
		agent.WithComposeOptions(compose.WithCallbacks(&callback.LoggerCallback{})))
	if err != nil {
		fmt.Printf("failed to stream: %v", err)
		return err
	}
	reasonContent := ""
	finalContent := ""
	for {
		msg, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// finish
				// 保存用户消息到历史记录
				memory.SetMessages(&schema.Message{
					Role:    schema.User,
					Content: userMessage,
				})
				// 保存助手回复到历史记录
				memory.SetMessages(&schema.Message{
					Role:    schema.Assistant,
					Content: finalContent,
				})
				// 发送完整消息
				client <- &types.ChatResp{
					ChatId: chatId,
					Event:  "end",
					Data: &schema.Message{
						Role:             schema.Assistant,
						ReasoningContent: reasonContent,
						Content:          finalContent,
					},
				}
				break
			}
			fmt.Printf("failed to recv: %v", err)
			return err
		}
		reasonContent += msg.ReasoningContent
		finalContent += msg.Content
		// 发送流式消息
		client <- &types.ChatResp{
			ChatId: chatId,
			Event:  "message",
			Data:   msg,
		}
	}
	return nil
}

func (l *ChatLogic) queryInternalDocsTool(ctx context.Context, config config.Config, query string) ([]*schema.Document, error) {
	// 检索文档
	embedder := models2.NewEmbeddingModel(ctx)
	retriever := retriever.NewRetriever(ctx, embedder)
	documents, err := retriever.Retrieve(ctx, query)
	if err != nil {
		log.Fatalf("Failed to retrieve: %v", err)
		return nil, err
	}
	return documents, nil
}
