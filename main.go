package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID       uint
	Email    string
	Name     string
	Password string
}

type Task struct {
	ID         uint
	Name       string
	DueDate    string
	CategoryID uint
	IsDone     bool
	UserID     uint // Assuming tasks are associated with a user
}

type Categoty struct {
	ID     int
	Title  string
	Color  string
	UserID uint
}

var userStorage []User
var taskStorage []Task

var categoryStorage []Categoty

var authenticatedUser *User

func main() {
	fmt.Println(" (--- hello wellcome to twelfosoft todo ---) ")

	command := flag.String("command", "no command", "command to run ")
	flag.Parse()

	fmt.Println("")
	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter new command")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		return
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "list-tasks":
		listTasks()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid ->", command)
	}

}

func createTask() {
	fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	var name, duedate, category string

	fmt.Println(" please enter the task name: ")

	scanner.Scan()
	name = scanner.Text()

	fmt.Println(" please enter the due date: ")

	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println(" please enter the category: ")

	scanner.Scan()
	category = scanner.Text()

	categoryId, error := strconv.Atoi(category)
	if error != nil {
		fmt.Println("category-id isn't valid integer, %v\n", error)

		return
	}

	var isFound bool = false
	for _, c := range categoryStorage {
		if c.ID == categoryId && c.UserID == authenticatedUser.ID {
			isFound = true
			fmt.Print("break")
			break
		}
	}

	if !isFound {
		fmt.Printf("category-id is not valid \n")

		return
	}

	task := Task{
		ID:         uint(len(taskStorage) + 1),
		Name:       name,
		DueDate:    duedate,
		CategoryID: uint(categoryId),
		IsDone:     false,
		UserID:     authenticatedUser.ID, // Associate task with the authenticated user
	}

	taskStorage = append(taskStorage, task)

	fmt.Println("task informations => ", name, " ", duedate, " ", category)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)

	var title, color string
	fmt.Println(" please enter the category title: ")

	scanner.Scan()
	title = scanner.Text()

	fmt.Println(" please enter the category color: ")

	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category information => ", title, " ", color)

	c := Categoty{
		ID:     int(len(categoryStorage)),
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, c)

}

func registerUser() {
	fmt.Println("--- please register to create a new user ---")
	scanner := bufio.NewScanner(os.Stdin)

	var name, email, password string

	fmt.Println(" please enter the user name : ")

	scanner.Scan()
	name = scanner.Text()

	fmt.Println(" please enter the user email: ")

	scanner.Scan()
	email = scanner.Text()

	fmt.Println(" please enter the user password: ")

	scanner.Scan()
	password = scanner.Text()

	fmt.Println("create new user => ", " ", email, " ", password)

	user := User{
		ID:       uint(len(userStorage) + 1),
		Email:    email,
		Name:     name,
		Password: password,
	}
	userStorage = append(userStorage, user)
}

func login() {
	fmt.Println("--- please login to your account ---")
	scanner := bufio.NewScanner(os.Stdin)

	var email, password string

	fmt.Println(" please enter the user email: ")

	scanner.Scan()
	email = scanner.Text()

	fmt.Println(" please enter the user password: ")

	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			fmt.Println("You're Logged in.")
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("email or password is not correct")
		runCommand("login")
		return
	}

	fmt.Println("login successfully => ", email, " ", password)
}

func listTasks() {
	fmt.Println("\n--- your tasks ---\n--------------------------")
	foundCount := 0

	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			foundCount++
			output := fmt.Sprintf("Name: %s, Due Date: %s, Category: %f, Is Done: %t\n",
				task.Name, task.DueDate, task.CategoryID, task.IsDone)
			fmt.Println(output)
		}
	}
	if foundCount == 0 {
		fmt.Println("No tasks found ... [empty]")
	}
	fmt.Println("--------------------------")
	fmt.Println("Total tasks found: ", foundCount)
}
