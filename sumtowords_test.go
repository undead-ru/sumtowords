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
		wantErr       error
	}{
		{"12345678", Roubles, false, "двенадцать миллионов триста сорок пять тысяч шестьсот семьдесят восемь рублей", nil},
		{"12345678912300", Roubles, false, "двенадцать триллионов триста сорок пять миллиардов шестьсот семьдесят восемь миллионов девятьсот двенадцать тысяч триста рублей", nil},
		{"10000001000000", Roubles, false, "десять триллионов один миллион рублей", nil},
		{"10101010011001", Roubles, false, "десять триллионов сто один миллиард десять миллионов одиннадцать тысяч один рубль", nil},
		{"01", Kopeks, false, "одна копейка", nil},
		{"99", Kopeks, false, "девяносто девять копеек", nil},
	}

	for _, test := range tests {
		if got, got2 := SumToString(test.inputAmount, test.inputCurrency, test.inputGender); got != test.wantAmount || got2 != test.wantErr {
			t.Errorf("SumToString(%q) = %v", test.inputAmount, got)
		}
	}

}

func TestStringToNotesAndCoins(t *testing.T) {
	var tests = []struct {
		input     string
		wantNotes string
		wantCoins string
		wantErr   error
	}{
		{"123456789.10", "123456789", "10", nil},
		{"1234567890123.004", "1234567890123", "00", nil},
		{"1.0059", "1", "01", nil},
	}

	for _, test := range tests {
		if got, got2, got3 := StringToNotesAndCoins(test.input); got != test.wantNotes || got2 != test.wantCoins || got3 != test.wantErr {
			t.Errorf("StringToNotesAndCoins(%q) = %v.%v", test.input, got, got2)
		}
	}

}
