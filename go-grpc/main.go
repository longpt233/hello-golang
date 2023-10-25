package main

import (
	"fmt"
	"sync"
)

func main() {
	c := &Customer{id: "long", age: 23}
	fmt.Println(c)
	c.UpdateAge(-10)
	// fmt.Println(c)
}

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

func (c *Customer) UpdateAge(age int) error {
	c.mutex.Lock()
	defer func() {
		fmt.Println("dong khoa")
		c.mutex.Unlock()
	}()

	var er error

	if age < 0 {
		er = fmt.Errorf("age should be positive for customer %v", c) // deadlock here: c call String() has lock
	}

	c.age = age

	return er
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}
