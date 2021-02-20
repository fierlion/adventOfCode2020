package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "strconv"
  "sort"
)

const INPUTFILE = "./input.txt"

func main() {
  inputFile, _ := filepath.Abs(INPUTFILE)
  file, err := os.Open(inputFile)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  intInputArr := []int{}
  for scanner.Scan() {
    rawInput := scanner.Text()
    converted, _ := strconv.Atoi(rawInput)
    intInputArr = append(intInputArr, converted)
  }
  sort.Ints(intInputArr)
  
  for start, val := range(intInputArr) {
    target := 2020 - val
    if found, val1, val2 := twoPointers(intInputArr[start:len(intInputArr)], target); found {
      result := val * val1 * val2
      fmt.Printf("%d\n", result)
    }
  }
}

func twoPointers(ints []int, target int) (bool, int, int) {
  start := 0
  end := len(ints)-1
  for start < end {
    this_result := ints[start] + ints[end]
    if this_result == target {
      return true, ints[start], ints[end]
    } else if this_result < target {
      start = start+1
    } else if this_result > target {
      end = end-1
    }
  }  
  return false,-1,-1  
}
