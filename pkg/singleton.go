package pkg

import "fmt"

type Singleton struct {
	Type string
}

func (s *Singleton) Print() {
	fmt.Println("Singleton type: ", s.Type)
}

func NewSingleton(item *Singleton, typeName string) *Singleton {
	if item == nil {
		return &Singleton{Type: typeName}
	}

	fmt.Println("Singleton already exists")

	return item
}
