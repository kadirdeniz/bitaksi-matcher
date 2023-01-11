package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang/mock/gomock"
	"io"
	"matcher/pkg"
	"matcher/test/mock"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testingObj *testing.T

func TestFiber(t *testing.T) {
	testingObj = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fiber Suite")
}

var _ = Describe("Fiber", Ordered, func() {
	ctrl := gomock.NewController(testingObj)
	defer ctrl.Finish()

	var mockRepository *mock.MockIRepository

	BeforeEach(func() {
		mockRepository = mock.NewMockIRepository(ctrl)
	})

	Context("Create a new handler", func() {
		It("Should return a new handler", func() {
			handler := NewHandler(mockRepository)
			Expect(handler).NotTo(BeNil())
		})
	})

	Context("Get nearest driver", func() {

		var responseObject pkg.Response
		mockLocation := mock.Location
		getNearestDriverRequest := mock.Coordiantes

		When("Driver is found", func() {
			It("Should return a driver", func() {

				mockRepository.EXPECT().GetNearestDriver(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockLocation, nil).Times(1)
				handler := NewHandler(mockRepository)
				Expect(handler).NotTo(BeNil())

				app := fiber.New()
				app.Get("/api/v1/drivers/nearest", handler.GetNearestDriver)

				req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/drivers/nearest?lat=%f&long=%f", getNearestDriverRequest.Latitude, getNearestDriverRequest.Longitude), nil)

				resp, err := app.Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(200))

				responseBody, err := io.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())

				err = json.Unmarshal(responseBody, &responseObject)
				Expect(err).To(BeNil())

				Expect(responseObject.Success).To(BeTrue())
				Expect(responseObject.Message).To(Equal("Driver found"))
				Expect(responseObject.Data).NotTo(BeNil())
			})
		})

		When("Driver is not found", func() {
			It("Should return an error", func() {

				mockRepository.EXPECT().GetNearestDriver(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, pkg.ErrDriverNotFound).Times(1)
				handler := NewHandler(mockRepository)
				Expect(handler).NotTo(BeNil())

				app := fiber.New()
				app.Get("/api/v1/drivers/nearest", handler.GetNearestDriver)

				req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/drivers/nearest?lat=%f&long=%f", getNearestDriverRequest.Latitude, getNearestDriverRequest.Longitude), nil)

				resp, err := app.Test(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(404))

				responseBody, err := io.ReadAll(resp.Body)
				Expect(err).NotTo(HaveOccurred())

				err = json.Unmarshal(responseBody, &responseObject)
				Expect(err).To(BeNil())

				Expect(responseObject.Success).To(BeFalse())
				Expect(responseObject.Message).To(Equal(pkg.ErrDriverNotFound.Error()))
				Expect(responseObject.Data).To(BeNil())
			})
		})

		When("Invalid request", func() {
			When("Fields are missing", func() {
				It("Should return an error", func() {
					handler := NewHandler(mockRepository)
					Expect(handler).NotTo(BeNil())

					app := fiber.New()
					app.Get("/api/v1/drivers/nearest", handler.GetNearestDriver)

					req := httptest.NewRequest("GET", "/api/v1/drivers/nearest", nil)

					resp, err := app.Test(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(400))

					responseBody, err := io.ReadAll(resp.Body)
					Expect(err).NotTo(HaveOccurred())

					err = json.Unmarshal(responseBody, &responseObject)
					Expect(err).To(BeNil())

					Expect(responseObject.Success).To(BeFalse())
					Expect(responseObject.Message).To(Equal(pkg.ErrInvalidRequest.Error()))
					Expect(responseObject.Data).To(BeNil())
				})
			})

			When("Fields are invalid", func() {
				It("Should return an error", func() {
					handler := NewHandler(mockRepository)
					Expect(handler).NotTo(BeNil())

					app := fiber.New()
					app.Get("/api/v1/drivers/nearest", handler.GetNearestDriver)

					req := httptest.NewRequest("GET", "/api/v1/drivers/nearest?lat=abc&long=def", nil)

					resp, err := app.Test(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(400))

					responseBody, err := io.ReadAll(resp.Body)
					Expect(err).NotTo(HaveOccurred())

					err = json.Unmarshal(responseBody, &responseObject)
					Expect(err).To(BeNil())

					Expect(responseObject.Success).To(BeFalse())
					Expect(responseObject.Message).To(Equal(pkg.ErrInvalidRequest.Error()))
					Expect(responseObject.Data).To(BeNil())
				})
			})

			When("Fields are out of range", func() {
				It("Should return an error", func() {
					handler := NewHandler(mockRepository)
					Expect(handler).NotTo(BeNil())

					app := fiber.New()
					app.Get("/api/v1/drivers/nearest", handler.GetNearestDriver)

					req := httptest.NewRequest("GET", "/api/v1/drivers/nearest?lat=1000&long=2000", nil)

					resp, err := app.Test(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(400))

					responseBody, err := io.ReadAll(resp.Body)
					Expect(err).NotTo(HaveOccurred())

					err = json.Unmarshal(responseBody, &responseObject)
					Expect(err).To(BeNil())

					Expect(responseObject.Success).To(BeFalse())
					Expect(responseObject.Message).To(Equal(pkg.ErrInvalidRequest.Error()))
					Expect(responseObject.Data).To(BeNil())
				})
			})
		})
	})

	Context("ErrorHandler", func() {
		var responseObject pkg.Response

		It("Should return an error", func() {
			handler := NewHandler(mockRepository)
			Expect(handler).NotTo(BeNil())

			app := fiber.New(fiber.Config{
				ErrorHandler: handler.ErrorHandler,
			})
			app.Use(recover.New())

			app.Get("/api/v1/error", func(c *fiber.Ctx) error {
				panic("error")
			})

			req := httptest.NewRequest("GET", "/api/v1/error", nil)

			resp, err := app.Test(req)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(500))

			responseBody, err := io.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			err = json.Unmarshal(responseBody, &responseObject)
			Expect(err).To(BeNil())

			Expect(responseObject.Success).To(BeFalse())
			Expect(responseObject.Message).To(Equal(pkg.ErrInternalServer.Error()))
			Expect(responseObject.Data).To(BeNil())
		})
	})
})
