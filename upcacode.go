package barcode

import (
	"fmt"
)

const ManufacturerCodeLength int = 5
const ProductCodeLength int = 5

type UPCA struct {
	NumberSystem     int
	ManufacturerCode string
	ProductCode      string
}

func NewUPCA(NumberSystem int, ManufacturerCode string, ProductCode string) (*UPCA, error) {
	if NumberSystem < 0 || NumberSystem > 1 {
		return nil, fmt.Errorf("number system must be between 0-9")
	}
	if len(ManufacturerCode) != 5 || len(ProductCode) != 5 {
		return nil,
			fmt.Errorf("either manufacture code length is not %d or product code length is not %d",
				ManufacturerCodeLength,
				ProductCodeLength)
	}

	return &UPCA{NumberSystem: NumberSystem, ManufacturerCode: ManufacturerCode, ProductCode: ProductCode}, nil
}
