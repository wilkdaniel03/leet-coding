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

func OptimalTime2(input int) bool {
  str := strconv.Itoa(input)
  return helper(str, 0)
}

func helper(str string, i int) bool {
  if i >= len(str) -1 -i {
    return true
  }

  if str[i] != str[len(str) -1 -i] {
    return false
  }

  return helper(str, i+1)
}
