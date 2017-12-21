package sumtowords

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

var (
	// Тип преобразуемой строки: рубли
	Roubles = -1
	// Тип преобразуемой строки: копейки
	Kopeks  = -2
)

// ReSumm регулярное выражение для разбора текстового представления числа с плавающей запятой
var ReSumm = regexp.MustCompile(`^\d+([\.|\,]){0,1}\d*$`)

// SumToString переводит числовое представление суммы в рублях в текстовое
func SumToString(d string, currency int, gender bool) (res string, err error) {

	if ReSumm.MatchString(d) == false {
		return "", fmt.Errorf("summ pattern didn't match: %v", d)
	}

	if d == "0" || d == "00" {
		if currency == Roubles {
			return "ноль рублей", nil
		}
			return "ноль копеек", nil
	}

	s := strings.Split(d, "")
	l := len(d)

	for i := 0; i < l; i++ {
		z := l - i
		switch z {
		case 1, 4, 7, 10, 13:
			if l != z && s[i-1] == "1" && s[i] != "0" {
				res += numbers["1"+s[i]].a
			} else {
				if z == 1 {
					z = currency
				}
				if numeral[z].gender == false && (s[i] == "1" || s[i] == "2") {
					res += numbers[s[i]].b
				} else {
					res += numbers[s[i]].a
				}

			}

			if l >= z+3 && s[i] == "0" && s[i-1] == "0" && s[i-2] == "0" && z > 1 {
			} else {
				if z == 1 {
					z = currency
				}
				switch s[i] {
				case "1":
					if i == 0 {
						res += numeral[z].a
					} else if s[i-1] == "1" {
						res += numeral[z].b
					} else {
						res += numeral[z].a
					}
				case "2", "3", "4":
					if i == 0 {
						res += numeral[z].c
					} else if s[i-1] == "1" {
						res += numeral[z].b
					} else {
						res += numeral[z].c
					}

				default:
					res += numeral[z].b
				}
			}

		case 2, 5, 8, 11, 14:
			if s[i] == "1" && s[i+1] != "0" {
			} else {
				res += numbers[s[i]+"0"].a
			}

		case 3, 6, 9, 12, 15:
			res += numbers[s[i]+"00"].a

		default:
		}
	}
	return res, nil
}

type numeralTypes struct {
	a, b, c string
	gender  bool
}

var numeral = map[int]numeralTypes{
	13: numeralTypes{a: "триллион ", b: "триллионов ", c: "триллиона ", gender: true},
	10: numeralTypes{a: "миллиард ", b: "миллиардов ", c: "миллиарда ", gender: true},
	7:  numeralTypes{a: "миллион ", b: "миллионов ", c: "миллиона ", gender: true},
	4:  numeralTypes{a: "тысяча ", b: "тысяч ", c: "тысячи ", gender: false},
	-1: numeralTypes{a: "рубль", b: "рублей", c: "рубля", gender: true},
	-2: numeralTypes{a: "копейка", b: "копеек", c: "копейки", gender: false},
}

type numberTypes struct {
	a, b string
}

var numbers = map[string]numberTypes{
	"1":   numberTypes{a: "один ", b: "одна "},
	"2":   numberTypes{a: "два ", b: "две "},
	"3":   numberTypes{a: "три "},
	"4":   numberTypes{a: "четыре "},
	"5":   numberTypes{a: "пять "},
	"6":   numberTypes{a: "шесть "},
	"7":   numberTypes{a: "семь "},
	"8":   numberTypes{a: "восемь "},
	"9":   numberTypes{a: "девять "},
	"11":  numberTypes{a: "одиннадцать "},
	"12":  numberTypes{a: "двенадцать "},
	"13":  numberTypes{a: "тринадцать "},
	"14":  numberTypes{a: "четырнадцать "},
	"15":  numberTypes{a: "пятнадцать "},
	"16":  numberTypes{a: "шестнадцать "},
	"17":  numberTypes{a: "семнадцать "},
	"18":  numberTypes{a: "восемнадцать "},
	"19":  numberTypes{a: "девятнадцать "},
	"10":  numberTypes{a: "десять "},
	"20":  numberTypes{a: "двадцать "},
	"30":  numberTypes{a: "тридцать "},
	"40":  numberTypes{a: "сорок "},
	"50":  numberTypes{a: "пятьдесят "},
	"60":  numberTypes{a: "шестьдесят "},
	"70":  numberTypes{a: "семьдесят "},
	"80":  numberTypes{a: "восемьдесят "},
	"90":  numberTypes{a: "девяносто "},
	"100": numberTypes{a: "сто "},
	"200": numberTypes{a: "двести "},
	"300": numberTypes{a: "триста "},
	"400": numberTypes{a: "четыреста "},
	"500": numberTypes{a: "пятьсот "},
	"600": numberTypes{a: "шестьсот "},
	"700": numberTypes{a: "семьсот "},
	"800": numberTypes{a: "восемьсот "},
	"900": numberTypes{a: "девятьсот "},
}

// BigFloatToNotesAndCoins переводит число в формате big.Float в рубли и копейки
func BigFloatToNotesAndCoins(d *big.Float) (notes, coins string, err error) {

	s := strings.Split(d.SetMode(big.AwayFromZero).Text('f', 2), ".")
	notes = s[0]

	if len(notes) > 14 {
		err = fmt.Errorf("Summ more than 999 billions")
	}

	coins = "00"

	if len(s) > 1 {
		coins = s[1]
		if len(coins) == 1 {
			coins += "0"
		}
	}
	return
}

// StringToNotesAndCoins переводит строку содержащую число в рубли и копейки
func StringToNotesAndCoins(d string) (notes, coins string, err error) {

	if ReSumm.MatchString(d) == false {
		err = fmt.Errorf("Summ is not a digit: %s", d)
	} else {
		if n, _, err := big.ParseFloat(strings.Replace(d, ",", ".", -1), 10, 53, big.AwayFromZero); err == nil {

			notes, coins, err = BigFloatToNotesAndCoins(n)

		} else {
			err = fmt.Errorf("Summ is not a big.Float: %s", d)
		}
	}
	return
}
