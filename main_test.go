package cex

import "testing"

func TestCurrency(t *testing.T) {
	c, err := ExchangeRate("USD", "ZAR")
	if err != nil {
		t.Error("Error testing USD/ZAR")
		t.Error(err)
	} else {
		t.Log("USD/ZAR =", c.Amount)
	}
}
