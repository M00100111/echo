package main

//
//import (
//	"context"
//	"github.com/bytedance/sonic"
//	"github.com/cloudwego/eino-ext/components/model/agenticark"
//	"github.com/cloudwego/eino/components/model"
//	"github.com/cloudwego/eino/schema"
//	"github.com/joho/godotenv"
//	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
//	"log"
//)
//
//func chat() {
//	ctx := context.Background()
//
//	// 加载 .env 文件
//	err := godotenv.Load(".env")
//	if err != nil {
//		panic(err)
//	}
//
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
//
//	input := []*schema.AgenticMessage{
//		schema.UserAgenticMessage("今天发生了哪些事？"),
//	}
//
//	concatenated, err := schema.ConcatAgenticMessages(msgs)
//	if err != nil {
//		log.Fatalf("failed to concat agentic messages, err: %v", err)
//	}
//
//	meta := concatenated.ResponseMeta.Extension.(*agenticark.ResponseMetaExtension)
//	for _, block := range concatenated.ContentBlocks {
//		if block.ServerToolCall == nil {
//			continue
//		}
//
//		serverToolArgs := block.ServerToolCall.Arguments.(*agenticark.ServerToolCallArguments)
//
//		args, _ := sonic.MarshalIndent(serverToolArgs, "  ", "  ")
//		log.Printf("server_tool_args: %s", string(args))
//	}
//
//	log.Printf("request_id: %s", meta.ID)
//	respBody, _ := sonic.MarshalIndent(concatenated, "  ", "  ")
//	log.Printf("  body: %s", string(respBody))
//}
