package math

import "testing"

func TestGcd(t *testing.T) {
	gcd := Gcd(12, 10)
	if gcd != 2 {
		t.Errorf("Should be 2 is %v", gcd)
	}
	gcd = Gcd(27, 18)
	if gcd != 9 {
		t.Errorf("Should be 9 is %v", gcd)
	}
	gcd = Gcd(27, 27)
	if gcd != 27 {
		t.Errorf("Should be 27 is %v", gcd)
	}
}
