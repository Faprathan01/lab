package unit

import (
	"testing"

	"t/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomer(t *testing.T) {
	// g := GomegaWithT(t)
	g := NewGomegaWithT(t)

	t.Run(`Customer is positive`,func(t *testing.T) {
		customer := entity.Customers{
			Name:"aom",
			Age: "15",
			CustomerID: "CM12345678",
		}
		ok,err := govalidator.ValidateStruct(customer)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
			

}

