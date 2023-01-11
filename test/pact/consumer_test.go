package pact

import (
	"github.com/pact-foundation/pact-go/dsl"
	"matcher/internal"
	"matcher/pkg"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPact(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pact Suite")
}

var _ = Describe("Pact", func() {

	var pact *dsl.Pact

	repository := internal.NewRepository()

	BeforeEach(func() {
		pact = &dsl.Pact{
			Host:                     "localhost",
			Consumer:                 "MatcherService",
			Provider:                 "DriverService",
			DisableToolValidityCheck: true,
			PactFileWriteMode:        "merge",
			LogDir:                   "./logs",
		}
	})

	AfterEach(func() {
		pact.Teardown()
	})

	// Context for finding a driver
	Context("Find a nearest driver", func() {
		When("a driver exists", func() {
			It("should return nearest driver's coordinates and the distance info", func() {
				pact.
					AddInteraction().
					Given("a driver exists").
					UponReceiving("a request to find a nearest driver").
					WithRequest(dsl.Request{
						Method:  "GET",
						Path:    dsl.String("/api/v1/drivers/nearest"),
						Query:   dsl.MapMatcher{"lat": dsl.String("51.507351"), "long": dsl.String("-0.127758"), "api_key": dsl.String("123")},
						Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
					}).
					WillRespondWith(dsl.Response{
						Status:  200,
						Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
						Body:    dsl.MapMatcher{"lat": dsl.String("51.507351"), "long": dsl.String("-0.127758"), "distance": dsl.String("0")},
					})

				err := pact.Verify(func() error {
					_, err := repository.GetNearestDriver(pact.Server.Port, 51.507351, -0.127758, "123")
					return err
				})
				Expect(err).To(BeNil())

			})
		})

		When("a driver does not exist", func() {
			It("should return an error", func() {
				pact.
					AddInteraction().
					Given("a driver does not exist").
					UponReceiving("a request to find a nearest driver").
					WithRequest(dsl.Request{
						Method:  "GET",
						Path:    dsl.String("/api/v1/drivers/nearest"),
						Query:   dsl.MapMatcher{"lat": dsl.String("51.507351"), "long": dsl.String("-0.127758"), "api_key": dsl.String("123")},
						Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
					}).
					WillRespondWith(dsl.Response{
						Status:  404,
						Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
						Body:    dsl.MapMatcher{"error": dsl.String("driver not found")},
					})

				err := pact.Verify(func() error {
					_, err := repository.GetNearestDriver(pact.Server.Port, 51.507351, -0.127758, "123")
					return err
				})
				Expect(err).To(Equal(pkg.ErrDriverNotFound))
			})
		})

		When("api key is invalid", func() {
			It("should return an error", func() {
				pact.
					AddInteraction().
					Given("a driver exists").
					UponReceiving("a request to find a nearest driver").
					WithRequest(dsl.Request{
						Method:  "GET",
						Path:    dsl.String("/api/v1/drivers/nearest"),
						Query:   dsl.MapMatcher{"lat": dsl.String("51.507351"), "long": dsl.String("-0.127758"), "api_key": dsl.String("1234")},
						Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
					}).
					WillRespondWith(dsl.Response{
						Status:  401,
						Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
						Body:    dsl.MapMatcher{"error": dsl.String("invalid api key")},
					})

				err := pact.Verify(func() error {
					_, err := repository.GetNearestDriver(pact.Server.Port, 51.507351, -0.127758, "1234")
					return err
				})

				Expect(err.Error()).To(Equal(pkg.ErrInvalidAPIKey.Error()))

			})
		})
	})
})
