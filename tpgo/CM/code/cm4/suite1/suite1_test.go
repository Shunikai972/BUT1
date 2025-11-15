package suite1

import "testing"

func TestSuite1Arret(t *testing.T) {
	u0 := suite1(0)
	if u0 != 0 {
		t.Fatal("pb condition arret")
	}
}
func TestSuite1Recursif(t *testing.T) {
	u5 := suite1(5)
	if u5 != 31 {
		t.Fatal("pb u5")
	}
}
