package main
import (
	"fmt"
	"study/person/person"
)

func main() {
	fmt.Println("main func")
	var person person.Person
	person.SetPerson("xiaoming","19","man","178","65")
	person.SayHello()
	person.ShowInf()
}
