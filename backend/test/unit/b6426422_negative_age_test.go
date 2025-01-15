package unit

import (
	"testing"

	"t/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerNegativeAge(t *testing.T) {
	
	g := NewGomegaWithT(t)

	t.Run(`Age is negative`,func(t *testing.T) {
		customer := entity.Customers{
			Name:"aom",
			Age: "jjj",//ผิด
			CustomerID: "CM12345678",
		}
		ok,err := govalidator.ValidateStruct(customer)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Age is not interger"))
	})
			

}
