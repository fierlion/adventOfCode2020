package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
)

const INPUTFILE = "./input.txt"

func main() {
  inputFile, _ := filepath.Abs(INPUTFILE)
  file, err := os.Open(inputFile)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  
  slope := make([][]bool, 0)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    rawInput := scanner.Text()
    thisRow := make([]bool, 0)
    for _, val := range(rawInput) {
      if (val == '#') {
        thisRow = append(thisRow, true)  
      } else {
        thisRow = append(thisRow, false)
      }
    }
    slope = append(slope, thisRow) 
  }
  result := countSmashes(slope, 2, 1)
  fmt.Printf("result: %d\n", result)
}

func countSmashes(slope [][]bool, x, y int) int {
  rowPos := 0
  result := 0
  rowSize := len(slope)
  for i:=0;i<rowSize;i+=x { 
    row := slope[i]
    if row[rowPos] {
      result = result + 1
    }
    rowPos = (rowPos + y) % len(row)
  }
  return result
}
