package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thalysonalexr/learning-go/pointers/dequeue"
)

func PrintMenu() {
	fmt.Println("*-- DEQUEUE --*")
	fmt.Println("1 - add head")
	fmt.Println("2 - add tail")
	fmt.Println("3 - remove head")
	fmt.Println("4 - remove tail")
	fmt.Println("5 - print head")
	fmt.Println("6 - print tail")
	fmt.Println("7 - show length")
	fmt.Println("8 - search binary")
	fmt.Println("0 - quit")
}

func userFactory(id, age int, name string) dequeue.User {
	return dequeue.User{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

func InputUser() dequeue.User {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID: ")
	id, _ := reader.ReadString('\n')
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Age: ")
	age, _ := reader.ReadString('\n')

	parsedId, _ := strconv.Atoi(strings.Replace(id, "\n", "", -1))
	parsedAge, _ := strconv.Atoi(strings.Replace(age, "\n", "", -1))

	return userFactory(parsedId, parsedAge, name)
}

func main() {
	d := dequeue.NewDequeue()
	reader := bufio.NewReader(os.Stdin)

	PrintMenu()
	for {
		fmt.Print("> ")
		opIn, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao realizar leitura")
			os.Exit(1)
		}

		o, _ := strconv.Atoi(strings.Replace(opIn, "\n", "", -1))

		switch o {
		case 1:
			fmt.Println("Add new user to head")
			d.AddHead(InputUser())
			break
		case 2:
			fmt.Println("Add new user to tail")
			d.AddTail(InputUser())
			break
		case 3:
			fmt.Println("Remove user from head")
			d.RemoveHead()
			break
		case 4:
			fmt.Println("Remove user from tail")
			d.RemoveTail()
			break
		case 5:
			d.PrintHead()
			break
		case 6:
			d.PrintTail()
			break
		case 7:
			fmt.Println(d.Length)
			break
		case 8:
			fmt.Println("Search binary")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			parsedId, _ := strconv.Atoi(strings.Replace(id, "\n", "", -1))
			user, err := d.SearchBinary(parsedId)
			if err == nil {
				fmt.Printf(dequeue.TEMPLATE_FORMAT, user.Id, user.Name, user.Age)
			} else {
				fmt.Println(err.Error())
			}
			break
		case 0:
			os.Exit(1)
			break

		default:
			fmt.Println("Operation not defined")
		}
	}
}
