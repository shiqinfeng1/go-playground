package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var age int

type Person struct {
	sync.Mutex
}

func (p *Person) AddAge() {
	defer wg.Done()
	p.Lock()
	age++
	defer p.Unlock()

}

func main() {
	p1 := Person{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go p1.AddAge()
	}
	wg.Wait()
	fmt.Println(age)
}
