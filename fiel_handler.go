package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id      string `json:"id"`
	Tag     string `json:"tag"`
	Created string `json:"created"`
}

func generateUUID() string {
	return time.Now().Local().String()
}

func file_io() {
	help := flag.Bool("help", false, "Help with commands")
	flag.Parse()

	if *help {
		fmt.Println(" -------- Todo Command Line -----------")
		fmt.Println("Commands :- ")
		fmt.Println("1. add `task` -- To add a todo\n2. delete [id] -- Pass id to delete a specific todo or all will be deleted\n3. list -- To list all the todos")
		os.Exit(0)
	}

	args := flag.Args()

	var file *os.File
	var err error

	if args[0] == "delete" || args[0] == "add" {
		file, err = os.OpenFile("todo.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal("Error occurred", err)
			os.Exit(1)
		}

		if args[0] == "add" {
			if args[1] == "" {
				log.Fatalf("Give a todo to add")
			} else {
				todo := Todo{
					Id:      uuid.NewString(),
					Tag:     args[1],
					Created: time.Now().Local().String(),
				}

				var todos []Todo
				fileData, err := os.ReadFile("todo.json")

				if err == nil && len(fileData) > 0 {
					err := json.Unmarshal(fileData, &todos)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					// if err != nil {
					// 	err := json.Unmarshal(fileData, &singleTodo)
					// 	if err != nil {
					// 		fmt.Println("here : ", err)
					// 		os.Exit(1)
					// 	}
					// 	todos = append(todos, singleTodo)
					// 	jsonData, err := json.MarshalIndent(todos, "", " ")
					// 	if err != nil {
					// 		fmt.Println(err)
					// 		os.Exit(1)
					// 	}
					// 	_, err = file.Write(jsonData)
					// 	if err != nil {
					// 		log.Fatal(err)
					// 		os.Exit(1)
					// 	}
					// } else {
					todos = append(todos, todo)
					jsonData, err := json.MarshalIndent(todos, "", " ")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					err = os.WriteFile("todo.json", jsonData, fs.FileMode(os.O_WRONLY))
					if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}
					// }
				} else {

					todos = append(todos, todo)
					jsonData, err := json.MarshalIndent(todos, "", " ")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					} else {
						_, err := file.Write(jsonData)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						} else {
							fmt.Println("Todo added successfully: ", todos)
						}
					}

				}

				fmt.Println("Added todo")
				os.Exit(0)
			}
		} else {
			if len(args) < 2 {
				// delete all the todo
				var perm rune
				fmt.Print("Do you want to delete all the todos ? (y/n) ")
				fmt.Scanf("%c", &perm)
				if perm == 'y' || perm == 'Y' {
					var fileData []byte
					err := os.WriteFile("todo.json", fileData, fs.FileMode(os.O_WRONLY))
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					fmt.Println("Deleted all the todos")
					os.Exit(0)
				}
			} else {
				var todos []Todo
				fileData, err := os.ReadFile("todo.json")
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				err = json.Unmarshal(fileData, &todos)
				if err != nil {
					fmt.Println("Error here ", err)
					os.Exit(1)
				}

				var newTodos []Todo
				for _, todo := range todos {
					if todo.Id != args[1] {
						newTodos = append(newTodos, todo)
					}
				}

				fileData, err = json.MarshalIndent(newTodos, "", " ")
				if err != nil {
					fmt.Println("Error here ", err)
					os.Exit(1)
				}
				err = os.WriteFile("todo.json", fileData, fs.FileMode(os.O_WRONLY))
				if err != nil {
					fmt.Println("Error here ", err)
					os.Exit(1)
				}
				fmt.Println("Deleted todo with id ", args[1])
				os.Exit(0)
			}
		}

	} else if args[0] == "list" {
		file, err = os.OpenFile("todo.json", os.O_RDONLY|os.O_CREATE, 0744)
		if err != nil {
			log.Fatal("Error occurred", err)
			os.Exit(1)
		}

		fileData, err := os.ReadFile("todo.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			var todos []Todo
			err := json.Unmarshal(fileData, &todos)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("File Data : ")
			fmt.Println("Created On ------- Task : Id")
			for _, todo := range todos {
				fmt.Println(todo.Created, "------", todo.Tag, ":", todo.Id)
			}
		}

	} else {
		log.Fatalf("Unknown command %s\nFor commands type --help", args[0])
	}
	defer file.Close()

}
