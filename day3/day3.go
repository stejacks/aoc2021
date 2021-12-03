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

  ones := make([]int, 12)
  zeros := make([]int, 12)
  gamma := make([]string, 12)
  epsilon := make([]string, 12)
  var j int

  for i := 0; i < len(lines); i++ {
    bits := strings.Split(lines[i], "")

    for j = 0; j < len(bits); j++ {
    if bits[j] == "1" {
      ones[j] += 1
    } else {
      zeros[j] += 1
    }
  } 
}

for j = 0; j < len(ones); j++ {
  if zeros[j] < ones[j] {
    gamma[j] = "0"
    epsilon[j] = "1"
  } else {
    gamma[j] = "1"
    epsilon[j] = "0"
   } 

}

gammastring := strings.Join(gamma[:], "")
epsilonstring := strings.Join(epsilon[:], "")
gammadecimal, err := strconv.ParseInt(gammastring, 2, 64)  
epsilondecimal, err := strconv.ParseInt(epsilonstring, 2, 64)  

power := gammadecimal * epsilondecimal
fmt.Println(power)
}
