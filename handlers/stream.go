package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Constante que contém a frase a ser transmitida em partes.
const streamPhrase = "Streaming responses are fun and useful"

// StreamPhrase implementa o streaming de palavras usando o pacote padrão net/http.
// Ele envia cada palavra da frase como um JSON separado, simulando um comportamento de streaming.
func (h *Handler) StreamPhrase(w http.ResponseWriter, r *http.Request) {
	// Configura os cabeçalhos para indicar que a resposta será transmitida em partes.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Transfer-Encoding", "chunked")

	// Verifica se o ResponseWriter suporta o flush (necessário para streaming).
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Divide a frase em palavras para transmitir uma por vez.
	words := strings.Split(streamPhrase, " ")
	for _, word := range words {
		// Formata a palavra atual como JSON delimitado por linha (NDJSON).
		// resp := fmt.Sprintf(`{"word":"%s"}`+"\n", word)
		resp := fmt.Sprintf(`{"word":"%s"}`, word)

		w.Write([]byte(resp))

		// Envia o conteúdo imediatamente para o cliente.
		flusher.Flush()

		// Aguarda 500ms antes de enviar a próxima palavra.
		time.Sleep(500 * time.Millisecond)
	}
}

// StreamPhraseGin implementa o streaming de palavras usando o framework Gin.
// Ele envia cada palavra da frase como um JSON separado, simulando um comportamento de streaming.
func (h *Handler) StreamPhraseGin(c *gin.Context) {
	// Configura os cabeçalhos para indicar que a resposta será transmitida em partes.
	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Encoding", "chunked")

	// Divide a frase em palavras para transmitir uma por vez.
	words := strings.Split(streamPhrase, " ")

	// Usa o método Stream do Gin para enviar os dados de forma incremental.
	c.Stream(func(w io.Writer) bool {
		for _, word := range words {
			// Formata a palavra atual como JSON e escreve no Writer.
			// resp := fmt.Sprintf(`{"word":"%s"}`+"\n", word)
			resp := fmt.Sprintf(`{"word":"%s"}`, word)
			w.Write([]byte(resp))

			fmt.Printf("%s ", word)

			// Envia o conteúdo imediatamente para o cliente.
			c.Writer.Flush()

			// Aguarda 500ms antes de enviar a próxima palavra.
			time.Sleep(500 * time.Millisecond)
		}
		// Retorna false para indicar que o streaming foi concluído.
		return false
	})
}
