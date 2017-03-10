package sumtowords

import (
	"testing"
)

func TestSumToString(t *testing.T) {
	var Roubles = -1
	var Kopeks = -2

	var tests = []struct {
		inputAmount   string
		inputCurrency int
		inputGender   bool
		wantAmount    string
	}{
		{"12345678", Roubles, false, "двенадцать миллионов триста сорок пять тысяч шестьсот семьдесят восемь рублей"},
		{"12345678912300", Roubles, false, "двенадцать триллионов триста сорок пять миллиардов шестьсот семьдесят восемь миллионов девятьсот двенадцать тысяч триста рублей"},
		{"10000001000000", Roubles, false, "десять триллионов один миллион рублей"},
		{"10101010011001", Roubles, false, "десять триллионов сто один миллиард десять миллионов одиннадцать тысяч один рубль"},
		{"01", Kopeks, false, "одна копейка"},
		{"99", Kopeks, false, "девяносто девять копеек"},
	}

	for _, test := range tests {
		if got := SumToString(test.inputAmount, test.inputCurrency, test.inputGender); got != test.wantAmount {
			t.Errorf("SumToString(%q) = %v", test.inputAmount, got)
		}
	}

}
