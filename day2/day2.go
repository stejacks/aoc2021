package main

import (  
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func main() {
 file, err := os.Open("input.txt")
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
	  	lines = append(lines, scanner.Text())
	}

  var horizontal int
  var vertical int

  for i := 0; i < len(lines); i++ {
    words := strings.Fields(lines[i])
    amount, err := strconv.Atoi(words[1])
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    if words[0] == "forward" {
      horizontal += amount
    } else if words[0] == "down"  {
      vertical += amount
    } else {
      vertical -= amount
    }
}
    

  fmt.Println(horizontal)
  fmt.Println(vertical)
  fmt.Println(horizontal * vertical)
}
