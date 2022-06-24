package main

import (
	"math"
	"sort"
)

func BruthForceSolution(arr1, arr2 []int) float64 {
  array := append(arr1, arr2...)
  sort.Ints(array)
  median := float64(array[(len(array) -1) >> 1] + array[len(array) >> 1]) / 2
  return median
}

func MedianOptimalApproach(arr1, arr2 []int) float64 {
  if len(arr1) > len(arr2) {
    arr1, arr2 = arr2, arr1
  }
  left, right := 0, len(arr1)-1
  totalLength := len(arr1) + len(arr2)
  half := totalLength / 2
  for {
    i := (left+right) / 2
    j := half - i - 2

    firstLeft := math.MinInt32
    if i > 0 {
      firstLeft = arr1[i]
    }
    firstRight := math.MaxInt32
    if i+1 < len(arr1) {
      firstRight = arr1[i+1]
    }
    secondLeft := math.MinInt32
    if j > 0 {
      secondLeft = arr2[j]
    }
    secondRight := math.MaxInt32
    if j+1 < len(arr2) {
      secondRight = arr2[j+1]
    }

    if firstLeft <= secondRight && secondLeft <= firstRight {
      if totalLength%2 != 0 {
        return float64(min(firstRight, secondRight))
      }
      return float64(max(firstLeft, secondLeft) + min(firstRight, secondRight)) / 2.0
    } else if firstLeft > secondRight {
      right = i-1
    } else {
      left = i+1
    }
  }
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func max(a, b int) int {
  if a < b {
    return b
  }
  return  a
}
