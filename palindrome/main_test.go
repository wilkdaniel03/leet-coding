package main

import "testing"

func TestBruthForceSolution(t *testing.T) {
  input1 := 121
  input2 := -121
  input3 := 10

  expected1 := true
  expected2 := false
  expected3 := false

  result1 := BruthForce(input1)
  result2 := BruthForce(input2)
  result3 := BruthForce(input3)

  if result1 != expected1 {
    t.Errorf("BruthForce(%v) = %v\nExpected %v", input1, result1, expected1)
  }

  if result2 != expected2 {
    t.Errorf("BruthForce(%v) = %v\nExpected %v", input2, result2, expected2)
  }

  if result3 != expected3 {
    t.Errorf("BruthForce(%v) = %v\nExpected %v", input3, result3, expected3)
  }
}

func TestOptimalTime1(t *testing.T) {
  input1 := 121
  input2 := -121
  input3 := 10

  expected1 := true
  expected2 := false
  expected3 := false

  result1 := OptimalTime1(input1)
  result2 := OptimalTime1(input2)
  result3 := OptimalTime1(input3)

  if result1 != expected1 {
    t.Errorf("OptimalTime1(%v) = %v\nExpected %v", input1, result1, expected1)
  }

  if result2 != expected2 {
    t.Errorf("OptimalTime1(%v) = %v\nExpected %v", input2, result2, expected2)
  }

  if result3 != expected3 {
    t.Errorf("OptimalTime1(%v) = %v\nExpected %v", input3, result3, expected3)
  }
}

func TestOptimalTime2(t *testing.T) {
  input1 := 121
  input2 := -121
  input3 := 10

  expected1 := true
  expected2 := false
  expected3 := false

  result1 := OptimalTime2(input1)
  result2 := OptimalTime2(input2)
  result3 := OptimalTime2(input3)

  if result1 != expected1 {
    t.Errorf("OptimalTime2(%v) = %v\nExpected %v", input1, result1, expected1)
  }

  if result2 != expected2 {
    t.Errorf("OptimalTime2(%v) = %v\nExpected %v", input2, result2, expected2)
  }

  if result3 != expected3 {
    t.Errorf("OptimalTime2(%v) = %v\nExpected %v", input3, result3, expected3)
  }
}
