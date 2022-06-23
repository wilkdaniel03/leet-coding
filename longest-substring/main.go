package main

// create map and pointers
// loop entire string
// on each iteration move window if needed

func Solution(s string) int {
  chars := make(map[byte]bool)
  left, right, res := 0, 0, 0
  for right < len(s) {
    if _, found := chars[s[right]]; !found {
      chars[s[right]] = true
      right += 1
      res = max(len(chars), res)
    } else {
      delete(chars, s[left])
      left += 1
    }
  }
  return res
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
