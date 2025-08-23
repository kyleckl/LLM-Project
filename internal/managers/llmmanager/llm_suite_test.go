package llmmanager_test

import (
	"context"
	"llmApp/internal/managers/llmmanager"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAdapters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "")
}

var _ = Describe("HTTPClient", func() {
	ctx := context.Background()
	llm := llmmanager.NewLLMClient("ollamaModel", "")

	It("Should be able to send a request to the LLM", func() {
		// Create a prompt, enabling streaming responses
		// Stream is set to true because we don't want to wait for the entire response
		prompt := llmmanager.QueryAdapter{
			Model:  "gemma3:1b",
			Prompt: "How has redlining in the United States affected urban development in recent years?",
			Stream: true,
		}

		resp, err := llm.SendQuery(ctx, prompt)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(200))
	})

	It("Should be able to stream a response from the LLM", func() {
		prompt := llmmanager.QueryAdapter{
			Model:  "gemma3:1b",
			Prompt: "Explain the theory of relativity in simple terms.",
			Stream: true,
		}

		resp, err := llm.SendQuery(ctx, prompt)
		Expect(err).ToNot(HaveOccurred())

		// Use the response body as a parameter to stream the response
		outputString, err := llm.StreamResponse(ctx, resp)
		Expect(err).ToNot(HaveOccurred())
		Expect(outputString).ToNot(BeEmpty())
	})
})
