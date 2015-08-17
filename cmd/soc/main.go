package main

import (
	soc "github.com/otobrglez/socialista"
	"fmt"
	"runtime"
)

func init() {
	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}
}

func main(){
	fmt.Println("\\m/ socialista \\m/")
	soc.GetStats("https://www.kickstarter.com/projects/elanlee/exploding-kittens");
}