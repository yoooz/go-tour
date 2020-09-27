package main

import "fmt"

func fibonacci() func() int {
	before, after := 0, 0
	
	return func() (result int) {
		if after == 0 {
			after = 1
			return 0
		}
		
		result = before + after
		before = after
		after = result
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

