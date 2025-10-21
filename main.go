package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Data-Types Structures
type Category struct {
	Title string
	Color string
}
type User struct {
	id       int
	email    string
	password string
	tasks    []Tasks
}

type Tasks struct {
	id           int
	taskCategory Category
	taskName     string
}

// In Memory System Storage
var categoryStorage = make([]Category, 10)
var userStorage = make([]User, 0)
var taskStorage = make([]Tasks, 0)
var AuthUser User

func main() {
	// CLI Command line Application Interface
	fmt.Println("Hello Todo App")
	command := flag.String("command", "new task", "Create a new task")
	// runCommand(*command)
	flag.Parse()
	for {
		runCommand((*command))
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {
	switch command {
	case "create task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		Login()
	case "show-users":
		fmt.Printf("%v \n", userStorage)
	case "exit":
		os.Exit(0)
	default:
		{
			println("command is not valid")
		}
	}
}

func createTask() {
	if len(AuthUser.password) > 1 {

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
		//  Note , in this case i enter a test category later we need to fix categories and in this part user should choose by exist category not creating one here
		newTask := Tasks{id: len(taskStorage) + 1, taskName: name, taskCategory: Category{Title: category, Color: "green"}}

		AuthUser.tasks = append(AuthUser.tasks, newTask)
	} else {
		fmt.Println("Please Login First")
	}
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
	category := Category{Title: categoryTitle, Color: categoryColor}
	categoryStorage = append(categoryStorage, category)
	fmt.Println(categoryStorage)
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
	user := User{
		id:       len(userStorage) + 1,
		email:    email,
		password: password}
	userStorage = append(userStorage, user)
}

func Login() {
	scanner := bufio.NewScanner((os.Stdin))
	var inputEmail, inputPassword string
	fmt.Println("PLEASE ENTER YOUR EMAIL")
	scanner.Scan()
	inputEmail = scanner.Text()
	fmt.Println("PLEASE ENTER YOUR PASSWORD")
	scanner.Scan()
	inputPassword = scanner.Text()
	// fmt.Println("user enterd ", email, password)

	for _, v := range userStorage {
		if v.email == inputEmail {
			var Auth bool = inputPassword == v.password
			if Auth {
				fmt.Println("login sucess")
				AuthUser = v
				fmt.Println(AuthUser)
			}
		} else {
			fmt.Println("Login Failed..")
		}
	}

}
