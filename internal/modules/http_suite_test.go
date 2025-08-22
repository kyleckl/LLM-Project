package modules_test

import (
	"context"
	"llmApp/internal/modules"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTTPClient", func() {
	// Instantiate the client and a context for requests
	ctx := context.Background()

	It("Should create a new HTTP client", func() {
		client := modules.NewHTTPClient("google.com")
		Expect(client).NotTo(BeNil())
	})

	It("Should send a POST request successfully", func() {
		client := modules.NewHTTPClient("https://httpbin.org")
		resp, err := client.Post(ctx, "/post", nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})

	It("Should send a GET request successfully", func() {
		client := modules.NewHTTPClient("https://httpbin.org")
		resp, err := client.Get(ctx, "/get", nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})
})
