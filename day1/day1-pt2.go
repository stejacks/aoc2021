package main

import (  
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
)

func main() {
 file, err := os.Open("input.txt")
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    var depths []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      depth, err := strconv.Atoi(scanner.Text())
  		if err != nil {
     			log.Fatal(err)
		  }
	  	depths = append(depths, depth)
	}

  var count int
  fmt.Println(len(depths))
  for i := 3; i < len(depths); i++ {
      window1 := depths[i-3]+depths[i-2]+depths[i-1]
      window2 := depths[i-2]+depths[i-1]+depths[i]
      if window1 < window2 {
      count++
    }
  }

  fmt.Println(count, " number of increases")

}
