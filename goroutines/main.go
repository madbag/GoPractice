package main

import (
	"fmt"
	"time"
)

func printNumbers(){
	for i:=0; i<10; i++{
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters(){
    for i:='a'; i<='e'; i++ {
        fmt.Printf("%c\n", i)
        time.Sleep(500 * time.Millisecond)
    }
}


func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
