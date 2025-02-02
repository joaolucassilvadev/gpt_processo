package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func Gpt(b string) string {
    var output string

    // Defina a chave da API diretamente no código
    apikey := "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"
    if apikey == "" {
        log.Fatal("Chave da API não encontrada")
    }

    // Cria um cliente OpenAI com a chave da API
    client := openai.NewClient(apikey)

    // Lê o conteúdo do arquivo de entrada
    inputFile := "./input_with_code.txt"
    filebytes, err := os.ReadFile(inputFile)
    if err != nil {
        log.Fatalf("Falha ao ler arquivo: %v", err)
    }

    // Define o contexto para a solicitação
    contexto := fmt.Sprintf("Explique sobre o texto:\n\n%s\n\n", string(filebytes))

    // Cria um contexto de execução
    ctx := context.Background()

    // Cria a solicitação de chat com o modelo
    req := openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo0125,
        Messages: []openai.ChatCompletionMessage{
            {Role: openai.ChatMessageRoleSystem, Content: contexto},
            {Role: openai.ChatMessageRoleUser, Content: b},
        },
        MaxTokens:   50,
        Temperature: 0.5,
    }

    // Envia a solicitação de chat ao modelo
    resp, err := client.CreateChatCompletion(ctx, req)
    if err != nil {
        log.Fatalf("Erro ao obter resposta do modelo: %v", err)
        return "Erro ao obter resposta"
    }

    // Exibe a resposta ao usuário
    output = resp.Choices[0].Message.Content
    // fmt.Printf("Resposta: %s\n", output)

    return output
}
