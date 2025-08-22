package llmmanager_test

import (
	"context"
	"encoding/json"
	"fmt"
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
	It("Should be able to send a query", func() {
		ctx := context.Background()
		llm := llmmanager.NewLLMClient("ollamaModel", "")
		prompt := llmmanager.QueryAdapter{
			Model:  "gemma3:1b",
			Prompt: "How has redlining in the United States affected urban development in recent years?",
			Stream: false,
		}

		resp, err := llm.SendQuery(ctx, prompt)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(200))

		defer resp.Body.Close()
		var response llmmanager.ResponseAdapter
		err = json.NewDecoder(resp.Body).Decode(&response)
		Expect(err).ToNot(HaveOccurred())
		Expect(response.Response).ToNot(BeEmpty())

		fmt.Printf("Response: %s", response.Response)
	})
})
