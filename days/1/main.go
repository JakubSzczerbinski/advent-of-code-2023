package main

import (
	"bufio"
	"fmt"
	"os"
)

type unit = struct{};
type set = map[string]unit

func match_forward(s string, i int, digits map[string]int) int {
  fmt.Println("forw", s, i)
  keys := set {}
  for d := range(digits) {
    keys[d] = unit {}
  }
  for j := 0; i + j < len(s) && len(keys) > 0; j++ {
    for k := range(keys) {
      if k[j] != s[i + j] {
        delete(keys, k)
        continue;
      }

      if len(k) - 1 == j {
        fmt.Println(k);
        return digits[k];
      }
    }
  }
  return -1
}

func match_backwards(s string, i int, digits map[string]int) int {
  keys := set {}
  for d := range(digits) {
    keys[d] = unit {}
  }
  for j := 0; i - j >= 0 && len(keys) > 0; j++ {
    for k := range(keys) {
      if k[len(k) - 1 - j] != s[i - j] {
        delete(keys, k)
        continue;
      }

      if len(k) - 1 == j {
        return digits[k];
      }
    }
  }
  return -1
}

func first(s string, digits map[string]int) int {
  for i := 0; i < len(s); i++ {
    value := match_forward(s, i, digits) 
    if value != -1 {
      return value
    }
  }
  return 0
}

func last(s string, digits map[string]int) int {
  for i := len(s) - 1; i >= 0; i-- {
    value := match_backwards(s, i, digits) 
    if value != -1 {
      return value
    }
  }
  fmt.Println(s)
  return 0
}

func calibration_value(line string, digits map[string]int) int {
  res := first(line, digits) * 10 + last(line, digits)
  return res
}

func main() {

  digits := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
    "four":  4,
    "five":  5,
    "six":   6,
    "seven": 7,
    "eight": 8,
    "nine":  9,
    "1":     1,
    "2":     2,
    "3":     3,
    "4":     4,
    "5":     5,
    "6":     6,
    "7":     7,
    "8":     8,
    "9":     9,
  }

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

  sum := 0
	for scanner.Scan() {
		line := scanner.Text()
    sum = sum + calibration_value(line, digits)
	}
  fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

