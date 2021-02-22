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
    // these aren't 0-indexed so fix that in the call below
    startNum, _ := strconv.Atoi(splitRange[0])
    endNum, _ := strconv.Atoi(splitRange[1])
    targetChar := []rune(target[:1])[0]
    if ok := passwordIsValid(startNum-1, endNum-1, targetChar, password); ok {
      validCount = validCount + 1 
    }
  }
  fmt.Printf("valid count: %d\n", validCount)
}

func passwordIsValid(startNum, endNum int, targetChar rune, password string) bool {
  runes := []rune(password)
  startChar := runes[startNum] == targetChar
  endChar := runes[endNum] == targetChar
  // return xor 
  return (startChar || endChar) && !(startChar && endChar)
}
