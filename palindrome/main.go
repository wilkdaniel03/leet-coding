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

func OptimalTime1(input int) bool {
  str := strconv.Itoa(input)
  newString := []byte{}
  for i := len(str)-1; i >= 0; i-- {
    newString = append(newString, str[i])
  }
  return string(newString) == str
}
