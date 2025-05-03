package barcode

import "testing"

func TestDummTest(t *testing.T) {
	qr := NewQrCode()
	qr.GenQrCode()
}
