package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RandA(srezA []string, srezR []string, str string) (bool, int) {

	for _, atype := range srezA {
		if atype == str {
			return true, 1
		}
	}

	for _, rtype := range srezR {
		if rtype == str {
			return true, 2
		}
	}
	return false, 0
	}

func arabCalculation(x int, y string, z int) int {
	
	switch {
	case y == "+":
	   res := x + z
	   return res
	case y == "-":
	   res := x - z
	   return res
	case y == "*":
	   res := x * z
	   return res
	case y == "/":
	   res := x / z
	   return res
	default:
	   panic("Неверный оператор. Доступные операторы: +, -, *, /")
} 
}

func romeCalculation(x string, y string, z string) string {
	
	arabToRome := map[int]string {
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 
		11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17:"XVII", 18: "XVIII", 19: "XIX", 20: "XX", 
		21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX", 
		31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL", 
		41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L", 
		51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX", 
		61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
		71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
		81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
		91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
		}
		
		romeToArab := make(map[string]int)
		
		for k, v := range arabToRome {
		romeToArab[v] = k
		}

		x1, z1 := romeToArab[x], romeToArab[z]

		switch {
		case y == "+":
		   resA := x1 + z1
		   resR := arabToRome[resA]
		   return resR
		case y == "-":
			if x1 <= z1 {
				panic("Результатом операции с римскими цифрами может быть только положительное число.")
			} else {
		   resA := x1 - z1
		   resR := arabToRome[resA]
		   return resR
			}
		case y == "*":
		   resA := x1 * z1
		   resR := arabToRome[resA]
		   return resR
		case y == "/":
			if x1 / z1 < 1 {
				panic("Результатом операции с римскими цифрами может быть только положительное число.")
			} else {
		   resA := x1 / z1
		   resR := arabToRome[resA]
		   return resR
			}
		default:
		   panic("Неверный оператор. Доступные операторы: +, -, *, /")

		
}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	//Выводим приветсвенный текст, получаем выражение, разделяем его на массив и присваем значения массива по переменным, где: a - первый операнд, b - оператор, с - второй операнд.

		fmt.Println("Привет, я - \"Katalator\"\nУмею считать арабские цифры(от 1 до 10) или римские цифры(от I до X). Только смешивать арабские и римские цифры не получится.\nСтавь пробелы между цифрами и оператором(+, -, *, /). Например: 4 - 2 или V + VIII.\nА теперь, введи значение:")
		text, _ := reader.ReadString('\n')
		vmass := strings.Fields(text)
		a := vmass[0]
		b := vmass[1]
		c := vmass[2]

	// Проверяем, выражение на соответствие правилу формата математической операции.

		switch {
		case len(vmass) != 3:
			panic("Пример введен неверно. Только 2 операнда от 1 до 10, либо от I до X и 1 оператор(+, -, *, /)")
		}

	// Создаем 2 среза, чтобы проверить операнды на величину.	

		srezA := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		srezR := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	// Загоняем срезы в функцию вместе с значением, на выходе получаем true - если значение есть в одном из срезов и цифру. 1 - арабское число, 2 - римское.

		aStatus, AorR1 := RandA(srezA, srezR, a)
		cStatus, AorR2 := RandA(srezA, srezR, c)

	// Если оба операнда подходят по правилу величины и являются арабскими, сразу же переводим операнды в int и проводим операцию с помощью функции "arabCalculation", которая определяет оператор и возвращает итог операции в переменную "total".
	// Если же числа римские, то в функции "romeCalculation" мы определяем полную карту "arabToRome" и переварачиваем ее ключи и значения, чтобы определить карту "romeToArab". Сразу переводим римские числа string в арабские числа int и, определяя оператор, получаем результат арабским числом. После чего переводим результат int в римское число string и возвращаем в переменную "totalrome".
	// Любое другое "Если" не подходит и выдается паника о неверном значении.
		if aStatus == true && cStatus == true && AorR1 == 1 && AorR2 == 1 {
			readyA, _ := strconv.Atoi(a)
			readyC, _ := strconv.Atoi(c)
			totalarab := arabCalculation(readyA, b, readyC) 
			fmt.Println(totalarab)
		} else if aStatus == true && cStatus == true && AorR1 == 2 && AorR2 == 2 {
			totalrome := romeCalculation(a, b, c)
			fmt.Println(totalrome)
		} else {
			panic("Неверное значение")
		}
	}