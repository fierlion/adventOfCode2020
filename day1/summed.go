package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "strconv"
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
  var exists = struct{}{}
  intInputArr := []int{}
  intInputSet := make(map[int]struct{}) 
  for scanner.Scan() {
    rawInput := scanner.Text()
    converted, _ := strconv.Atoi(rawInput)
    intInputArr = append(intInputArr, converted)
    intInputSet[converted] = exists 
  }
  for _, val := range intInputArr {
    partner := 2020 - val
    if _, ok := intInputSet[partner]; ok {
      fmt.Printf("found pair: %d %d\n", val, partner)
      fmt.Printf("result: %d\n", val*partner)
    }
  }
}
