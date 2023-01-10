package jwt

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJwt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jwt Suite")
}

var _ = Describe("Jwt", func() {

	authenticatedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjoidHJ1ZSJ9.V7NhQn8kJSmDkxBAuB-iU8PXf_UQfR0eWmPzRF3Lgr0"
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjoiZmFsc2UifQ.lpGdy2woonDCuX7OSoXvzr0mNTLEwL3Hz3sGzE5CMpA"

	Context("GetIsAuthenticated", func() {
		When("authenticated", func() {
			It("should return true", func() {
				Expect(NewJWT(authenticatedToken).GetIsAuthenticated()).To(BeTrue())
			})
		})

		When("not authenticated", func() {
			It("should return false", func() {
				Expect(NewJWT(invalidToken).GetIsAuthenticated()).To(BeFalse())
			})
		})
	})

	Context("DecodeToken", func() {
		It("should return the decoded token", func() {
			Expect(NewJWT(authenticatedToken).DecodeToken().DumpClaim["authenticated"]).To(Equal("true"))
		})
	})

	Context("GetToken", func() {
		It("should return the token", func() {
			Expect(NewJWT(authenticatedToken).GetToken()).To(Equal(authenticatedToken))
		})
	})

})
