package unit

import (
	"testing"
	"backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBestCase(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("All correct", func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0123456789",
			Password:    "1234567890",
			Url:         "http://example.com",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestPhoneNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("PhoneNumber length is not 10 digits", func(t *testing.T) {
        member := entity.Member{
            PhoneNumber: "12345",  // Invalid phone number (less than 10 digits)
            Password:    "1234567890",
            Url:         "http://example.com",
        }

        ok, err := govalidator.ValidateStruct(member)

        g.Expect(ok).To(BeFalse()) // Validation should fail
        g.Expect(err).NotTo(BeNil()) // Expect error
        g.Expect(err.Error()).To(Equal("Phone Number length is not 10 digits.")) // Correct error message for length validation
    })
}

func TestUrl(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Url is required", func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "1234567890",
			Password:    "1234567890",
			Url:         "google",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Url is invalid"))
	})
}