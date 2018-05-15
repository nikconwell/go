// test.go is a program that reads STDIN and changes the word "the" to the word "_the_"
package main

import (
    "bufio"
    "io"
    "fmt"
    "os"
    "strings"
    "github.com/mattn/go-isatty"
)

// Function main is even better.
func main() {

     if isatty.IsTerminal(os.Stdin.Fd()) {
     	fmt.Println("Expecting input on STDIN (^D at end...)")
     }

     reader := bufio.NewReader(os.Stdin)

     for {
          text, err := reader.ReadString('\n')
	  if err == io.EOF {
	    break
	  }

          text = strings.Replace(text, "the", "_the_",-1)

	  fmt.Print(text)
     }

}
