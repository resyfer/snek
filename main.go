package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

var milli = time.Millisecond

var snek snake

func main() {

	// Getting terminal height
	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	snek.locationUpdate(width/2, height/2)
	snek.lengthUpdate(1)
	snek.durationUpdate(1000 * milli)

}