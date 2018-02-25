package main

import (
	"fmt"
	"strconv"

	"github.com/makki0205/jobcounter"
)

func main() {
	j := jobcounter.NewJob()
	c := make(chan string)

	for i := 0; i <= 20; i++ {
		j.Add(1)
		go func(i int, c chan string, j jobcounter.JobCounter) {
			//エラーとする
			if i%7 == 0 {
				j.Done()
				return
			}
			c <- strconv.Itoa(i)
		}(i, c, j)
	}
	for j.IsEndJob() {
		select {
		case str := <-c:
			fmt.Println(str)
		default:
		}
	}
}
