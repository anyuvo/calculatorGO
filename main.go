package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите операцию (или 'exit' для выхода): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения: ", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}
		result, err := calculate(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fmt.Println("Результат:", result)
	}
}

func calculate(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", fmt.Errorf("неправильный формат операции")
	}
	op1, op2, operator := parts[0], parts[2], parts[1]

	isRomanNumeral := isRoman(op1) // Проверяем, является ли первое число римским
	if isRomanNumeral != isRoman(op2) {
		return "", fmt.Errorf("используются разные системы счисления")
	}
	var num1, num2 int
	if isRomanNumeral {
		num1 = fromRoman(op1)
		num2 = fromRoman(op2)
	} else {
		var err error
		num1, err = strconv.Atoi(op1)
		if err != nil {
			return "", fmt.Errorf("неправильное число: %s", op1)
		}
		num2, err = strconv.Atoi(op2)
		if err != nil {
			return "", fmt.Errorf("неправильное число: %s", op2)
		}
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		return "", fmt.Errorf("числа должны быть в диапазоне от 1 до 10")
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		return "", fmt.Errorf("неподдерживаемый оператор: %s", operator)
	}

	if isRomanNumeral {
		if result < 1 {
			return "", fmt.Errorf("в римских числах нет нуля или отрицательных чисел")
		}
		return toRoman(result), nil
	}
	return strconv.Itoa(result), nil
}

// Определение структуры для хранения пар "значение - римская цифра"
type RomanPair struct {
	Value  int
	Symbol string
}

var romanPairs = []RomanPair{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// Конвертация из арабского числа в римское
func toRoman(num int) string {
	var result strings.Builder
	for _, pair := range romanPairs {
		for num >= pair.Value {
			result.WriteString(pair.Symbol)
			num -= pair.Value
		}
	}
	return result.String()
}

// Конвертация из римского числа в арабское
func fromRoman(roman string) int {
	// Создание мапы для быстрого поиска
	romanMap := make(map[string]int)
	for _, pair := range romanPairs {
		romanMap[pair.Symbol] = pair.Value
	}
	result := 0
	i := 0
	for i < len(roman) {
		if i+1 < len(roman) && romanMap[roman[i:i+2]] != 0 {
			result += romanMap[roman[i:i+2]]
			i += 2
		} else {
			result += romanMap[string(roman[i])]
			i++
		}
	}
	return result
}

// Проверяет, является ли строка римским числом.
func isRoman(s string) bool {
	validRomanChars := "IVXLCDM"
	for _, char := range s {
		if !strings.ContainsRune(validRomanChars, char) {
			return false
		}
	}
	return true
}
