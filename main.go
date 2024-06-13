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
  if len(vals) < 3 {
    panic("Ошибка ввода! Допустимый фомат ввода `5 * 3` или `IX - VI`")
  }
  val1, val2, sign := vals[0], vals[2], vals[1]

  if num1, num2, ok := getValidRomanNums(val1, val2); ok {
    res := calculate(num1, num2, sign)
    if res < 1 {
      panic("Ошибка! Результат вычислений с римскими числами должен быть больше 1")
    }
    fmt.Println(intToRoman(res))
  } else {
    num1, num2 := getValidNum(val1), getValidNum(val2)
    res := calculate(num1, num2, sign)
    fmt.Println(res)
  }
}

func getValidNum(val string) int {
  if num, err := strconv.Atoi(val); num > 0 && num <= 10 && err == nil {
    return num
  }
  panic("Ошибка ввода! Введите только римские или только арабские числа от 1 до 10")
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
