package main

import (
	"fmt"
	"strings"
)

func main () {
	println(fmt.Sprintf("result: %v", romanToInt("MCMXCIV")))
}

func romanToInt(s string) int {
  romanArab := map[string]int{
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
  }

  var res int
  skip := false
  split := strings.Split(s, "")
  for i, l := range split {
    if skip {
      skip = false
      continue
    } else if i < (len(split) - 1) && romanArab[l] < romanArab[split[i+1]] {
      res = res + romanArab[split[i+1]] - romanArab[l]
      skip = true
    } else {
      res = res + romanArab[l]
    }
  }
    
  return res
}
