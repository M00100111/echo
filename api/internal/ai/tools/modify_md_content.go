package tools

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

// ModifyMdContentInput 修改md文件内容的输入参数
type ModifyMdContentInput struct {
	FilePath    string `json:"file_path" jsonschema:"description=The path to the markdown file that needs to be modified"`
	OldContent  string `json:"old_content" jsonschema:"description=The original content in the markdown file that needs to be replaced"`
	NewContent  string `json:"new_content" jsonschema:"description=The new content to replace the old content with"`
}

// ModifyMdContentOutput 修改md文件内容的输出结果
type ModifyMdContentOutput struct {
	Success   bool   `json:"success" jsonschema:"description=Indicates whether the file modification was successful"`
	Message   string `json:"message" jsonschema:"description=Status message describing the operation result"`
	FilePath  string `json:"file_path" jsonschema:"description=The path of the modified markdown file"`
}

// NewModifyMdContentTool 创建修改md文件内容的工具
func NewModifyMdContentTool() tool.InvokableTool {
	t, err := utils.InferOptionableTool(
		"modify_md_content",
		"Modify the content of a markdown file by replacing specified text. This tool reads the markdown file, replaces the old content with new content, and saves the changes. Use this tool when you need to update documentation or modify markdown content programmatically.",
		func(ctx context.Context, input *ModifyMdContentInput, opts ...tool.Option) (output string, err error) {
			// 读取原始文件内容
			file, err := os.Open(input.FilePath)
			if err != nil {
				log.Printf("Error opening file: %v", err)
				result := ModifyMdContentOutput{
					Success:  false,
					Message:  "Failed to open file",
					FilePath: input.FilePath,
				}
				jsonBytes, _ := json.Marshal(result)
				return string(jsonBytes), nil
			}
			defer file.Close()

			// 读取文件全部内容
			content, err := io.ReadAll(file)
			if err != nil {
				log.Printf("Error reading file: %v", err)
				result := ModifyMdContentOutput{
					Success:  false,
					Message:  "Failed to read file content",
					FilePath: input.FilePath,
				}
				jsonBytes, _ := json.Marshal(result)
				return string(jsonBytes), nil
			}

			// 替换内容
			updatedContent := []byte(content)
			oldBytes := []byte(input.OldContent)
			newBytes := []byte(input.NewContent)
			updatedContent = []byte(string(updatedContent)) // 确保是字符串类型

			// 检查是否包含要替换的内容
			if !containsBytes(updatedContent, oldBytes) {
				log.Printf("Old content not found in file: %s", input.FilePath)
				result := ModifyMdContentOutput{
					Success:  false,
					Message:  "Old content not found in the file",
					FilePath: input.FilePath,
				}
				jsonBytes, _ := json.Marshal(result)
				return string(jsonBytes), nil
			}

			// 执行替换
			updatedContent = replaceBytes(updatedContent, oldBytes, newBytes)

			// 写回文件
			err = os.WriteFile(input.FilePath, updatedContent, 0644)
			if err != nil {
				log.Printf("Error writing file: %v", err)
				result := ModifyMdContentOutput{
					Success:  false,
					Message:  "Failed to write file",
					FilePath: input.FilePath,
				}
				jsonBytes, _ := json.Marshal(result)
				return string(jsonBytes), nil
			}

			// 构建成功结果
			result := ModifyMdContentOutput{
				Success:  true,
				Message:  "Markdown file updated successfully",
				FilePath: input.FilePath,
			}

			// 转换为JSON
			jsonBytes, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				log.Printf("Error marshaling result to JSON: %v", err)
				return "", err
			}

			log.Printf("Successfully updated markdown file: %s", input.FilePath)
			return string(jsonBytes), nil
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return t
}

// containsBytes 检查字节切片中是否包含指定的字节序列
func containsBytes(data, sub []byte) bool {
	for i := 0; i <= len(data)-len(sub); i++ {
		if equalBytes(data[i:i+len(sub)], sub) {
			return true
		}
	}
	return false
}

// replaceBytes 替换字节切片中的指定字节序列
func replaceBytes(data, old, new []byte) []byte {
	result := make([]byte, 0, len(data))
	for i := 0; i < len(data); {
		if i <= len(data)-len(old) && equalBytes(data[i:i+len(old)], old) {
			result = append(result, new...)
			i += len(old)
		} else {
			result = append(result, data[i])
			i++
		}
	}
	return result
}

// equalBytes 比较两个字节切片是否相等
func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
