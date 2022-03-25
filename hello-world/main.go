package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	for _, each := range []string{"wick", "josee", "steve"} {
    go func(who string) {
      var data = fmt.Sprintf("hello %s", who)
      message <- data
    }(each)
  }

  for i := 0; i < 3; i++ {
    printMessage(message)
  }
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
