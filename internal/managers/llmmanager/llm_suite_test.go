package llmmanager_test

import (
	"context"
	"fmt"
	"llmApp/internal/managers/llmmanager"
	"os"
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
			err := llm.ConstructPrompt(prompt)
			Expect(err).ToNot(HaveOccurred())

			resp, err := llm.SendQuery(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
		})

		It("Should be able to stream a response from the LLM", func() {
			prompt := "Why is the meaning of life 42?"
			err := llm.ConstructPrompt(prompt)
			Expect(err).ToNot(HaveOccurred())

			resp, err := llm.SendQuery(ctx)
			Expect(err).ToNot(HaveOccurred())

			// Use the response body as a parameter to stream the response
			outputString, err := llm.StreamResponse(ctx, resp)
			Expect(err).ToNot(HaveOccurred())
			Expect(outputString).ToNot(BeEmpty())
		})
	})

	When("The desired LLM is hosted on Google", func() {
		GoogleAPIKey := os.Getenv("GOOGLE_API_KEY")

		BeforeEach(func() {
			if GoogleAPIKey == "" {
				Skip("missing google api key, skipping tests that require it")
			}
		})

		llm, _ := llmmanager.NewGoogleAIManager("gemma-3-27b-it", GoogleAPIKey)

		It("Should be able to send and receive a response from the LLM", func() {
			prompt := "How do airplanes fly?"
			err := llm.ConstructPrompt(prompt)
			Expect(err).ToNot(HaveOccurred())

			outputString, err := llm.SendQuery(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(outputString).ToNot(BeEmpty())

			fmt.Print(outputString)
		})
	})
})
