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

	When("The desired LLM is hosted on Llama", func() {
		llm := llmmanager.NewLlamaClient("gemma3:1b")

		It("Should be able to send a request to the LLM", func() {
			// Create a prompt, enabling streaming responses
			// Stream is set to true because we don't want to wait for the entire response
			prompt := "What is the meaning of life?"
			promptAdapter := llm.ConstructPrompt(prompt, true)
			llamaPromptAdapter := promptAdapter.(llmmanager.LlamaPromptAdapter)

			resp, err := llm.SendQuery(ctx, llamaPromptAdapter)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
		})

		It("Should be able to stream a response from the LLM", func() {
			prompt := "Why is the meaning of life 42?"
			promptAdapter := llm.ConstructPrompt(prompt, true)

			resp, err := llm.SendQuery(ctx, promptAdapter)
			Expect(err).ToNot(HaveOccurred())

			// Use the response body as a parameter to stream the response
			outputString, err := llm.StreamResponse(ctx, resp)
			Expect(err).ToNot(HaveOccurred())
			Expect(outputString).ToNot(BeEmpty())
		})
	})
})
