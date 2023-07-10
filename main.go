package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Romans = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50,
	"C": 100}

func findOp(text string) (string, error) {
	switch {
	case strings.Contains(text, "+"):
		return "+", nil
	case strings.Contains(text, "-"):
		return "-", nil
	case strings.Contains(text, "*"):
		return "*", nil
	case strings.Contains(text, "/"):
		return "/", nil
	default:
		return "", fmt.Errorf("Оператор не найден")
	}
}

func calculate(n1, n2 int, oper string) (int, error) {
	switch oper {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		return n1 / n2, nil
	default:
		return 0, fmt.Errorf("%s not found", oper)
	}
}

func isRomans(num string) bool {
	if _, err := Romans[strings.Split(num, "")[0]]; err {
		return true
	}
	return false
}

func romToInt(text string) int {
	sum := 0
	t := len(text)
	for i := 0; i < t; i++ {
		if i != t-1 && Romans[string(text[i])] < Romans[string(text[i+1])] {
			sum += Romans[string(text[i+1])] - Romans[string(text[i])]
			i++
			continue
		}
		sum += Romans[string(text[i])]
	}
	return sum
}

func intToRom(num int) string {
	var srom = ""
	var numbs = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romns = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var idx = len(romns) - 1

	for num > 0 {
		for numbs[idx] <= num {
			srom += romns[idx]
			num -= numbs[idx]
		}
		idx -= 1
	}
	return srom
}

func numsAndTyp(text, oper string) (n1, n2 int, roms bool, err error) {
	nums := strings.Split(text, oper)
	if len(nums) > 2 {
		return n1, n2, roms, fmt.Errorf("Много операторов")
	}
	oneRom := isRomans(nums[0])
	twoRom := isRomans(nums[1])

	if oneRom != twoRom {
		return n1, n2, roms, fmt.Errorf("Разные форматы значений")
	}

	if oneRom && twoRom {
		roms = true
		n1 = romToInt(nums[0])
		n2 = romToInt(nums[1])
	} else {
		n1, err = strconv.Atoi(nums[0])
		if err != nil {
			return
		}
		n2, err = strconv.Atoi(nums[1])
		if err != nil {
			return
		}
	}

	switch {
	case n1 <= 0 || n1 > 10 && n2 <= 0 || n2 > 10:
		return n1, n2, roms, fmt.Errorf("%d или %d меньше 1 или больше 10", n1, n2)
	}
	return n1, n2, roms, nil
}

func main() {

	read := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите пример: ")
		text, _ := read.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")

		oper, err := findOp(text)
		if err != nil {
			panic(err)
		}
		n1, n2, roms, err := numsAndTyp(text, oper)
		if err != nil {
			panic(err)
		}
		result, err := calculate(n1, n2, oper)
		if err != nil {
			panic(err)
		}
		if roms {
			if result <= 0 {
				panic("Римские числа не могут быть меньше или равны  0")
			}

			res := intToRom(result)

			fmt.Println("Вывод: \n", res)
		} else {
			fmt.Println("Вывод: \n", result)
		}
	}
}
