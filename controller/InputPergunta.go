package controller

import (
	"net/http"

	"examplegpt.com/service"
	"github.com/gin-gonic/gin"
)

type PerguntaDTO struct {
	Pergunta string `json:"pergunta"`
	Resposta string `json:"resposta"`
}

func InputPergunta(ctx *gin.Context) {
	var pergunta PerguntaDTO

	if err := ctx.ShouldBindJSON(&pergunta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensagem_error": "A pergunta está vazia ou inválida. Por favor, fale algo.",
		})
		return
	}

	pergunta.Resposta = service.Gpt(pergunta.Pergunta)

	if err := pergunta.Resposta; err == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensagem_error": " A resposta veio vazia por favor pergunte de novo",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"resposta para o usuario": pergunta.Resposta,
	})
}
