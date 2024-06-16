package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  fmt.Println("Калькулятор, введите римские или арабские числа от 1 до 10:")
  in := bufio.NewReader(os.Stdin)
  text, _ := in.ReadString('\n')
  vals := strings.Split(strings.Trim(text, "\n"), " ")
  if len(vals) != 3 {
    panic("Ошибка ввода! Допустимый фомат ввода `5 * 3` или `IX - VI`")
  }
  val1, val2, sign := vals[0], vals[2], vals[1]

  if num1, num2, ok := getValidRomanNums(val1, val2); ok {
    res := calculate(num1, num2, sign)
    if res < 1 {
      panic("Ошибка! В римской системе нет отрицательных чисел")
    }
    fmt.Println(intToRoman(res))
  } else {
    num1, num2 := getValidNum(val1), getValidNum(val2)
    fmt.Println(calculate(num1, num2, sign))
  }
}

func calculate(num1, num2 int, sign string) int {
  switch sign {
  case "+":
    return num1 + num2
  case "-":
    return num1 - num2
  case "*":
    return num1 * num2
  case "/":
    return num1 / num2
  default:
    panic("Ошибка ввода! Доступны только операции +,-,*,/")
  }
}

func getValidNum(val string) int {
  if num, err := strconv.Atoi(val); num > 0 && num <= 10 && err == nil {
    return num
  }
  panic("Ошибка ввода! Введите только римские или только арабские числа от 1 до 10")
}

func getValidRomanNums(val1, val2 string) (int, int, bool) {
  validRomanNums := map[string]int{
    "I":   1,
    "II":  2,
    "III": 3,
    "IV":  4,
    "V":   5,
    "VI":  6,
    "VII": 7,
    "VIII": 8,
    "IX":  9,
    "X":   10,
  }
  num1, ok1 := validRomanNums[val1]
  num2, ok2 := validRomanNums[val2]
  return num1, num2, ok1 && ok2
}

func intToRoman(num int) string {
  intVals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
  romanVals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
  romanNum := ""
  for i, val := range intVals {
    for num >= val {
      romanNum += romanVals[i]
      num -= val
    }
  }
  return romanNum
}
