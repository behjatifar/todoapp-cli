package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Category struct {
	CategoryTitle string
	CategoryColor string
}
type User struct {
	id       int
	email    string
	password string
}

var storage = make([]Category, 10)
var userStorage = make([]User, 1, 10)

func main() {
	// CLI Command line Application Interface
	fmt.Println("Hello Todo App")
	command := flag.String("command", "new task", "Create a new task")
	flag.Parse()
	runCommand(*command)

}

func runCommand(command string) {
	switch command {
	case "create task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	default:
		{
			println("command is not valid")
		}
	}
}

func createTask() {
	scanner := bufio.NewScanner((os.Stdin))
	var name, duedate, category string
	fmt.Println("please enter task title")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("please enter due-date title")
	scanner.Scan()
	duedate = scanner.Text()
	fmt.Println("please enter category ")
	scanner.Scan()
	category = scanner.Text()
	fmt.Println("task", name, duedate, category)
}

func createCategory() {
	scanner := bufio.NewScanner((os.Stdin))
	var categoryTitle, categoryColor string
	fmt.Println("enter category title")
	scanner.Scan()
	categoryTitle = scanner.Text()
	fmt.Println("enter category color")
	scanner.Scan()
	categoryColor = scanner.Text()
	// fmt.Printf(" category created: categoryTitle %s categoryColor %s ", categoryTitle, categoryColor)
	category := Category{CategoryTitle: categoryTitle, CategoryColor: categoryColor}
	storage = append(storage, category)
	fmt.Println(storage)
}

func registerUser() {
	scanner := bufio.NewScanner((os.Stdin))
	var email, password string

	fmt.Println("welcome to app let's take you to our club you are a new legend")
	fmt.Println("at first enter your email:")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("now we need a password for next logins")
	scanner.Scan()
	password = scanner.Text()
	user := User{len(userStorage) + 1, email, password}
	userStorage = append(userStorage, user)
}
