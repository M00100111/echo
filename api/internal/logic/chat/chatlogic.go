// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package chat

import (
	models "Echo/api/internal/ai/models/chatModel"
	models2 "Echo/api/internal/ai/models/embeddingModel"
	"Echo/api/internal/ai/retriever"
	"Echo/api/internal/config"
	"Echo/pkg/mem"
	"context"
	"errors"
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

	// 获取历史对话
	memory := mem.GetSimpleMemory(chatId)
	historyMessages := memory.GetMessages()

	var documents []*schema.Document
	if req.Message != "" {
		var err error
		documents, err = l.queryInternalDocsTool(l.ctx, l.svcCtx.Config, req.Message)
		if err != nil {
			return err
		}
	}

	chatModel := models.NewChatModel(l.ctx, l.svcCtx.Config)

	input := []*schema.AgenticMessage{}

	// 添加检索到的文档作为系统消息
	if len(documents) > 0 {
		docContent := "相关文档信息：\n"
		for _, doc := range documents {
			docContent += "- " + doc.Content + "\n"
		}
		input = append(input, schema.SystemAgenticMessage(docContent))
	}

	// 添加历史对话
	for _, msg := range historyMessages {
		// 根据消息类型转换为AgenticMessage
		switch msg.Role {
		case schema.User:
			input = append(input, schema.UserAgenticMessage(msg.Content))
		case schema.Assistant:
			temp := schema.UserAgenticMessage(msg.Content)
			temp.Role = schema.AgenticRoleTypeAssistant
			input = append(input, temp)
		}
	}

	// 添加用户消息
	userMessage := req.Message
	input = append(input, schema.UserAgenticMessage(userMessage))

	// 保存用户消息到历史记录
	memory.SetMessages(&schema.Message{
		Role:    schema.User,
		Content: userMessage,
	})

	resp, err := chatModel.Stream(l.ctx, input)
	if err != nil {
		log.Fatalf("failed to stream, err: %v", err)
	}

	var msgs []*schema.AgenticMessage
	var assistantResponse string
	for {
		msg, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// 保存助手回复到历史记录
				memory.SetMessages(&schema.Message{
					Role:    schema.Assistant,
					Content: assistantResponse,
				})

				// 发送结束事件
				client <- &types.ChatResp{
					ChatId: chatId,
					Event:  "end",
					Data:   msgs,
				}
				break
			}
			log.Fatalf("failed to receive stream response, err: %v", err)
		}
		msgs = append(msgs, msg)

		// 累加助手回复内容
		if msg.ContentBlocks != nil {
			for _, content := range msg.ContentBlocks {
				if content.AssistantGenText != nil && content.AssistantGenText.Text != "" {
					assistantResponse += content.AssistantGenText.Text
				}
			}
		}

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
	embedder := models2.NewEmbeddingModel(ctx, config)
	retriever := retriever.NewRetriever(ctx, config, embedder)
	documents, err := retriever.Retrieve(ctx, query)
	if err != nil {
		log.Fatalf("Failed to retrieve: %v", err)
		return nil, err
	}
	return documents, nil
}
