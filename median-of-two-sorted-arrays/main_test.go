package main

import "testing"

type Input struct {
  arr1 []int
  arr2 []int
}

func TestMedianBruthForce(t *testing.T) {
  input1 := Input{arr1: []int{1,3}, arr2: []int{2}}
  input2 := Input{arr1: []int{1,2}, arr2: []int{3,4}}

  expected1 := float32(2)
  expected2 := float32(2.5)

  result1 := BruthForceSolution(input1.arr1, input1.arr2)
  result2 := BruthForceSolution(input2.arr1, input2.arr2)

  if result1 != expected1 {
    t.Errorf("BruthForceSolution(%v, %v) = %f\nExpexted %f", input1.arr1, input1.arr2, result1, expected1)
  }

  if result2 != expected2 {
    t.Errorf("BruthForceSolution(%v, %v) = %f\nExpexted %f", input2.arr1, input2.arr2, result2, expected2)
  }
}
