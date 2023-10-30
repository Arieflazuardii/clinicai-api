package helpers

import (
	"context"

	"github.com/sashabaranov/go-openai"
)


func DiagnosticAI(userInput, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Pasien datang ke klinik dengan keluhan, Pasien ingin mencari informasi lebih lanjut tentang keluhan ini. Tolong berikan penjelasan yang relevan, termasuk penyebab, gejala, diagnosis, dan opsi pengobatan yang mungkin. Juga berikan saran tentang langkah-langkah pertama yang sebaiknya diambil oleh pasien untuk mengatasi keluhan ini, jika user tidak memberikan keluhan atau keluhan tidak spesifik maka berikan tips hidup sehat",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}