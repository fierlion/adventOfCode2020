package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "strings"
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
  validCount := 0
  for scanner.Scan() {
    rawInput := scanner.Text()
    splitStrings := strings.Split(rawInput, " ")
    passRange := splitStrings[0]
    target := splitStrings[1]
    password := splitStrings[2]
    splitRange := strings.Split(passRange, "-")
    startNum, _ := strconv.Atoi(splitRange[0])
    endNum, _ := strconv.Atoi(splitRange[1])
    targetChar := []rune(target[:1])[0]
    if ok := passwordIsValid(startNum, endNum, targetChar, password); ok {
      validCount = validCount + 1 
    }
  }
  fmt.Printf("valid count: %d\n", validCount)
}

func passwordIsValid(startNum, endNum int, targetChar rune, password string) bool {
  totalRune := 0
  runes := []rune(password)
  for _, rune := range(runes) {
    if rune == targetChar {
      totalRune = totalRune + 1
    }
  }
  return (totalRune >= startNum && totalRune <= endNum)
}
