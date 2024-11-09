package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
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
	flag.Parse()

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
					Id:      generateUUID(),
					Tag:     args[1],
					Created: time.Now().Local().String(),
				}

				var todos []Todo
				fileData, err := os.ReadFile("todo.json")

				if err == nil && len(fileData) > 0 {
					// err := json.Unmarshal(fileData, &todos)
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
				log.Fatalf("Give the id of tag to delete. Use command `list` to display the todo with id.")
			} else {
				var todos []Todo
				decoder := json.NewDecoder(file)
				err := decoder.Decode(&todos)
				if err != nil {
					fmt.Println("Error here ", err)
					os.Exit(1)
				}
				fmt.Println(todos)
				fmt.Printf("Deleted todo with id %s", args[1])
				os.Exit(0)
			}
		}

	} else if args[0] == "list" {
		fmt.Println(args[0])
		file, err = os.OpenFile("todo.json", os.O_RDONLY|os.O_CREATE, 0744)
		if err != nil {
			log.Fatal("Error occurred", err)
			os.Exit(1)
		}

	} else {
		log.Fatalf("Unknown args %s", args[0])
	}
	defer file.Close()

}
