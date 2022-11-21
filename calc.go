package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	massiveStr := strings.Fields(str)

	var strRom = "n I II III IV V VI VII VIII IX X XI XII XIII XIV XV XVI XVII XVIII XIX XX XXI XXII XXIII XXIV XXV XXVI XXVII XXVIII XXIX XXX XXXI XXXII XXXIII XXXIV XXXV XXXVI XXXVII XXXVIII XXXIX XL XLI XLII XLIII XLIV XLV XLVI	XLVII XLVIII XLIX L LI LII LIII LIV LV LVI LVII LVIII LIX LX LXI LXII LXIII LXIV LXV LXVI LXVII LXVIII LXIX LXX LXXI LXXII LXXIII LXXIV LXXV LXXVI LXXVII LXXVIII LXXIX LXXX LXXXI LXXXII LXXXIII LXXXIV LXXXV LXXXVI LXXXVII LXXXVIII LXXXIX XC XCI XCII XCIII XCIV XCV XCVI XCVII XCVIII XCIX C"
	massiveRom := strings.Fields(strRom)

	if len(massiveStr) < 2 {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	var aArabSim = massiveStr[0]
	var bArabSim = massiveStr[2]
	a, _ := strconv.ParseInt(aArabSim, 10, 32)
	b, _ := strconv.ParseInt(bArabSim, 10, 32)

	var aRomSim = massiveStr[0]
	var bRomSim = massiveStr[2]
	var aRom int
	var bRom int

	var cSim = massiveStr[1]
	var sign, _ = utf8.DecodeRuneInString(cSim)

	switch {

	case len(massiveStr) > 3:
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	case a == 0 && b == 0:
		fmt.Print()
	case a == 0 && b != 0:
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	case a != 0 && b == 0:
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	case a >= 11 || a <= 0:
		fmt.Println("Вывод ошибки: число не входит в диапазон от 1 до 10 включительно")
		return
	case b >= 11 || b <= 0:
		fmt.Println("Вывод ошибки: число не входит в диапазон от 1 до 10 включительно")
		return
	}

	if a != 0 && b != 0 {

		if sign == '+' {
			d := a + b
			fmt.Println(d)
		} else if sign == '-' {
			d := a - b
			fmt.Println(d)
		} else if sign == '*' {
			d := a * b
			fmt.Println(d)
		} else if sign == '/' {
			d := a / b
			fmt.Println(d)
		} else {
			fmt.Println("Вывод ошибки: введен недопустимый знак")
			return
		}

	} else {

		if aRomSim == "I" {
			aRom = 1
		} else if aRomSim == "II" {
			aRom = 2
		} else if aRomSim == "III" {
			aRom = 3
		} else if aRomSim == "IV" {
			aRom = 4
		} else if aRomSim == "V" {
			aRom = 5
		} else if aRomSim == "VI" {
			aRom = 6
		} else if aRomSim == "VII" {
			aRom = 7
		} else if aRomSim == "VIII" {
			aRom = 8
		} else if aRomSim == "IX" {
			aRom = 9
		} else if aRomSim == "X" {
			aRom = 10
		} else {
			fmt.Println("Вывод ошибки: число не входит в диапазон от I до X включительно")
			return
		}

		if bRomSim == "I" {
			bRom = 1
		} else if bRomSim == "II" {
			bRom = 2
		} else if bRomSim == "III" {
			bRom = 3
		} else if bRomSim == "IV" {
			bRom = 4
		} else if bRomSim == "V" {
			bRom = 5
		} else if bRomSim == "VI" {
			bRom = 6
		} else if bRomSim == "VII" {
			bRom = 7
		} else if bRomSim == "VIII" {
			bRom = 8
		} else if bRomSim == "IX" {
			bRom = 9
		} else if bRomSim == "X" {
			bRom = 10
		} else {
			fmt.Println("Вывод ошибки: число не входит в диапазон от I до X включительно")
			return
		}

		if sign == '+' {
			dRom := aRom + bRom
			fmt.Println(massiveRom[dRom])

		} else if sign == '-' {
			dRom := aRom - bRom
			if dRom < 0 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
				return
			} else if dRom == 0 {
				fmt.Println("Вывод ошибки, так как в римской системе нет ноля.")
			} else {
				fmt.Println(massiveRom[dRom])
			}
		} else if sign == '*' {
			dRom := aRom * bRom
			fmt.Println(massiveRom[dRom])

		} else if sign == '/' {
			dRom := aRom / bRom
			fmt.Println(massiveRom[dRom])
		} else {
			fmt.Println("Вывод ошибки: введен недопустимый знак")
			return
		}
	}
}
