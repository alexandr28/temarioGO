package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	// determina el numero de procesadores que usa  gruntime
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	numbers:= []uint32{1424,235,1246,822,64,22,16,2,4,10}
	wg.Add(len(numbers))
	fmt.Println("Comienza el proceso")
	for _, v := range numbers  {
		go func(a uint32) {
			defer wg.Done()
			fmt.Println(a,EsPrimo(v))
		}(v)
	}
	wg.Wait()
	fmt.Println("Termina el Proceso")
}
func EsPrimo(a uint32) bool  {
	c:=0
	var i uint32
	for i=1; i <= a; i++ {
		if a%i ==0{
			c++
		}
	}
	if c==2 {
		return true
	}
	return false
}
