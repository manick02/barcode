package barcode

import (
	"testing"
)

func TestUpcaCodeConstruction(t *testing.T) {
	t.Run("ProductCodeValidation", func(t *testing.T) {

		_, err := NewUPCA(1, "1111", "12345")
		if err == nil { // Invalid product code length
			t.Fail()
		}
	})
	t.Run("NumericValidationCheck", func(t *testing.T) {
		t.Run("NumericValidationCheck", func(t *testing.T) {

			_, err := NewUPCA(10, "11111", "12345")
			if err == nil { // Invalid product code length
				t.Fail()
			}
		})
	})
	t.Run("ManufacturerCodeValidationCheck", func(t *testing.T) {
		_, err := NewUPCA(10, "11111", "12345")
		if err == nil { // Invalid product code length
			t.Fail()
		}

	})
}
