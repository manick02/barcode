package barcode

import "testing"

func TestQrGen(t *testing.T) {
	qr := NewQrCode()
	qr.GenQrCode()
}

func TestDumTestDecode(t *testing.T) {
	qr := NewQrCode()
	x := qr.Decode()
	if x == nil {
		t.Fail()
	}
}
