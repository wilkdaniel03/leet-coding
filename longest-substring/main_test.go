package main

import "testing"

func TestLongestSubstring(t *testing.T) {
  input1 := "abcabcbb"
  input2 := "bbbbb"
  input3 := "pwwkew"

  expected1 := 3
  expected2 := 1
  expected3 := 3

  result1 := Solution(input1)
  result2 := Solution(input2)
  result3 := Solution(input3)

  if result1 != expected1 {
    t.Errorf("Solution(%s) = %d\nExpected %d", input1, result1, expected1)
  }

  if result2 != expected2 {
    t.Errorf("Solution(%s) = %d\nExpected %d", input2, result2, expected2)
  }

  if result3 != expected3 {
    t.Errorf("Solution(%s) = %d\nExpected %d", input3, result3, expected3)
  }
}
