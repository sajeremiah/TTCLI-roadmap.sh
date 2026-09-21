package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	ts "github.com/sajeremiah/TTCLI-roadmap.sh/internal/task"
)

var badArguments = errors.New("wrong arguments")

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env load error:", err)
		os.Exit(1)
	}
	storageFileName := os.Getenv("STORAGE_FILENAME")
	localStorage, err := ts.LoadStorage(storageFileName)
	if err != nil {
		fmt.Println("data load error:", err)
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		fmt.Println(badArguments)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		if len(os.Args) != 3 {
			fmt.Println(badArguments)
			os.Exit(1)
		}
		localStorage.AddTask(os.Args[2])
	case "update":
		if len(os.Args) != 4 {
			fmt.Println(badArguments)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id convertion error:", err)
			os.Exit(1)
		}
		err = localStorage.UpdateTask(id, os.Args[3])
		if err != nil {
			fmt.Println("update task error:", err)
			os.Exit(1)
		}
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println(badArguments)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id convertion error:", err)
			os.Exit(1)
		}
		err = localStorage.DeleteTask(id)
		if err != nil {
			fmt.Println("delete task error:", err)
			os.Exit(1)
		}
	case "list":
		if len(os.Args) > 3 {
			fmt.Println(badArguments)
		}
		if len(os.Args) == 2 {
			fmt.Println(localStorage.Tasks)
		} else {
			list, err := localStorage.ListTasks(os.Args[2])
			if err != nil {
				fmt.Println("list tasks error:", err)
			}
			fmt.Println(list)
		}
	case "mark-done":
		if len(os.Args) != 3 {
			fmt.Println(badArguments)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id convertion error:", err)
			os.Exit(1)
		}
		err = localStorage.MarkTask(id, "done")
	case "mark-in-progress":
		if len(os.Args) != 3 {
			fmt.Println(badArguments)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("id convertion error:", err)
			os.Exit(1)
		}
		err = localStorage.MarkTask(id, "in-progress")
	}

	if err = localStorage.SaveStorage(storageFileName); err != nil {
		fmt.Println("data save error:", err)
		os.Exit(1)
	}
}
