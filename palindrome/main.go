package main

import "strconv"

func BruthForce(input int) bool {
  str := strconv.Itoa(input)
  newString := ""
  for i := len(str)-1; i >= 0; i-- {
    newString += string(str[i])
  }
  return newString == str
}
