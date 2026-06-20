package main

import (
	"fmt"
	"os"
	"strconv"
)

// CommandHandler reads os.Args and runs the right action on Todos.
func CommandHandler(todos *Todos, storage *Storage[Todos]) {

	args := os.Args

	if len(args) < 2 {
		printUsage()
		return
	}

	command := args[1]

	switch command {

	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a title. Example: todo-cli add \"buy milk\"")
			return
		}
		todos.add(args[2])
		storage.Save(*todos)
		fmt.Println("Todo added!")

	case "list":
		todos.print()

	case "done":
		if len(args) < 3 {
			fmt.Println("Please provide an index. Example: todo-cli done 0")
			return
		}
		index, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Index must be a number.")
			return
		}
		if err := todos.markCompleted(index); err != nil {
			return
		}
		storage.Save(*todos)
		fmt.Println("Todo marked as completed!")

	case "edit":
		if len(args) < 4 {
			fmt.Println("Please provide an index and new title. Example: todo-cli edit 0 \"buy eggs\"")
			return
		}
		index, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Index must be a number.")
			return
		}
		if err := todos.edit(index, args[3]); err != nil {
			return
		}
		storage.Save(*todos)
		fmt.Println("Todo updated!")

	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide an index. Example: todo-cli delete 0")
			return
		}
		index, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Index must be a number.")
			return
		}
		if err := todos.delete(index); err != nil {
			return
		}
		storage.Save(*todos)
		fmt.Println("Todo deleted!")

	default:
		fmt.Println("Unknown command:", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println(`
Usage: todo-cli <command> [arguments]

Commands:
  add "<title>"          Add a new todo
  list                   List all todos
  done <index>           Mark a todo as completed
  edit <index> "<title>" Edit a todo's title
  delete <index>         Delete a todo
`)
}
