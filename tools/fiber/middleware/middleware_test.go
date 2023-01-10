package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"matcher/pkg"
	"matcher/test/mock"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFiberMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fiber Middleware Suite")
}

var _ = Describe("Fiber Middleware", func() {

	authenticatedToken := mock.AuthenticatedToken
	invalidToken := mock.InvalidToken
	var responseObject pkg.Response

	Context("GetIsAuthenticated", func() {

		getNearestDriverRequest := mock.Coordiantes

		When("Bearer token exists", func() {
			When("Token is authenticated", func() {
				It("shouldn't be error", func() {
					app := fiber.New()
					app.Get("/api/v1/drivers/nearest", IsAuthenticated)

					// create a request to our endpoint with authorization header
					req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/drivers/nearest?lat=%f&long=%f", getNearestDriverRequest.Latitude, getNearestDriverRequest.Longitude), nil)
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authenticatedToken))

					// send the request to the app
					resp, err := app.Test(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).ShouldNot(Equal(401))
				})
			})

			When("Token is not authenticated", func() {
				It("should be error", func() {
					app := fiber.New()
					app.Get("/api/v1/drivers/nearest", IsAuthenticated)

					// create a request to our endpoint with authorization header
					req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/drivers/nearest?lat=%f&long=%f", getNearestDriverRequest.Latitude, getNearestDriverRequest.Longitude), nil)
					req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", invalidToken))

					// send the request to the app
					resp, err := app.Test(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).Should(Equal(401))

					// read the response body
					body, err := io.ReadAll(resp.Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(json.Unmarshal(body, &responseObject)).NotTo(HaveOccurred())

					Expect(responseObject.Success).Should(BeFalse())
					Expect(responseObject.Message).Should(Equal("Unauthorized"))
					Expect(responseObject.Data).Should(BeNil())
				})
			})
		})

		When("Bearer token doesn't exist", func() {
			It("shouldn be error", func() {
				app := fiber.New()
				app.Get("/api/v1/drivers/nearest", IsAuthenticated)

				// create a request to our endpoint with authorization header
				req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/drivers/nearest?lat=%f&long=%f", getNearestDriverRequest.Latitude, getNearestDriverRequest.Longitude), nil)
				req.Header.Set("Authorization", fmt.Sprintf("Bearer "))

				// send the request to the app
				resp, err := app.Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(401))

				responseBody, err := io.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())

				err = json.Unmarshal(responseBody, &responseObject)
				Expect(err).To(BeNil())

				Expect(responseObject.Success).To(Equal(false))
				Expect(responseObject.Message).To(Equal("Unauthorized"))
				Expect(responseObject.Data).To(BeNil())
			})
		})
	})
})
