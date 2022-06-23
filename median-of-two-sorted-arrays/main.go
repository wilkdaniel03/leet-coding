package main

import "sort"

func BruthForceSolution(arr1, arr2 []int) float64 {
  array := append(arr1, arr2...)
  sort.Ints(array)
  median := float64(array[(len(array) -1) >> 1] + array[len(array) >> 1]) / 2
  return median
}
