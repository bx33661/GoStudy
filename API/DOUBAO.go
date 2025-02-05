package main

import (
	"context"
	"fmt"
	ark "github.com/sashabaranov/go-openai"
	"io"
)

func main() {
	config := ark.DefaultConfig("-----")
	config.BaseURL = "https://ark.cn-beijing.volces.com/api/v3"
	client := ark.NewClientWithConfig(config)

	fmt.Println("----- standard request -----")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		ark.ChatCompletionRequest{
			Model: "ep-20250205124041-wmtml",
			Messages: []ark.ChatCompletionMessage{
				{
					Role:    ark.ChatMessageRoleSystem,
					Content: "你是豆包，是由字节跳动开发的 AI 人工智能助手",
				},
				{
					Role:    ark.ChatMessageRoleUser,
					Content: "常见的十字花科植物有哪些？",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Message.Content)

	fmt.Println("----- streaming request -----")
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		ark.ChatCompletionRequest{
			Model: "ep-20250205124041-wmtml",
			Messages: []ark.ChatCompletionMessage{
				{
					Role:    ark.ChatMessageRoleSystem,
					Content: "你是豆包，是由字节跳动开发的 AI 人工智能助手",
				},
				{
					Role:    ark.ChatMessageRoleUser,
					Content: "常见的十字花科植物有哪些？",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("stream chat error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Stream chat error: %v\n", err)
			return
		}

		if len(recv.Choices) > 0 {
			fmt.Print(recv.Choices[0].Delta.Content)
		}
	}
}
