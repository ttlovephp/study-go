package lib

import (
	"fmt"
)

// Person ...
type Person struct {
	Name string
	Age  int
}

// Add ...
func Add(a, b int) int {
	return a + b
}

// GetName ...
func (p Person) GetName() string {
	return fmt.Sprintf("my name is %s", p.Name)
}

// ModifyName1 ...
func (p Person) ModifyName1(name string) {
	p.Name = name
	fmt.Println("Modify name 1 is", p.Name)
}

// ModifyName2 ...
func (p *Person) ModifyName2(name string) {
	p.Name = name
	fmt.Println("Modify name 2 is", p.Name)
}
