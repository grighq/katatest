package main

func getValidRomanNums(val1, val2 string) (int, int, bool) {
  validRomanNums := map[string]int{
    "I":   1,
    "II":  2,
    "III": 3,
    "IV":  4,
    "V":   5,
    "VI":  6,
    "VII": 7,
    "IIX": 8,
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
