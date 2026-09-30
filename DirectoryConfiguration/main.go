package main

import (
	userdata "FileDirectoryCreation/UserData"
	parsing_data_from_json "FileDirectoryCreation/json/Parsing"
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	isDirectoryWritten = "true"
)

type DIRECTORY struct {
	Name        string
	isWrittenTo string
}

func (dir DIRECTORY) GetDirectory() string {
	val, err := strconv.ParseBool(dir.isWrittenTo)

	if err != nil {
		log.Fatal(err)
	}

	if val {
		return fmt.Sprintf("%s", dir.Name)
	}

	return fmt.Sprintf("%s", dir.Name)

}

type RouteInputLayer struct {
	GetValue        func() string
	isLayerComplete bool
}

func (route RouteInputLayer) Get() bool {
	if route.isLayerComplete {
		return true
	}

	return false
}

type DirectorySetup struct {
	JSDirectorySetup func() string
}

const (
	JS int = 0
	RS int = 1
	CS int = 2
)

func isNil(err error) bool {
	return err != nil
}

func main() {

	inputUser := ""

	for i := range 1 {
		i += 1

		if i == 1 {

			UserNameObject := userdata.UserData{
				GetUsername: func() string {
					return "User"
				},

				USERNAME: strings.ToUpper("User"),
			}

			fmt.Println(UserNameObject.Get())

			routeValueOptions := RouteInputLayer{
				isLayerComplete: true,

				GetValue: func() string {
					return strings.ToUpper("Enter Name | Route: ")
				},
			}

			content, err := os.ReadFile("../json/Config.json")

			if isNil(err) {
				fmt.Println("Error Has Not Found Json File to Parse Data")
			}

			var data parsing_data_from_json.DataTypedMesh
			err = json.Unmarshal(content, &data)

			if isNil(err) {
				fmt.Println("Error Has Not Found Json File to Parse Data")
			}

			if data.IsDataFinal == false {
				fmt.Println("Welcome to:", data.PrintElement())
				fmt.Println("Current Version:", data.GetVersion())
			}

			fmt.Println(routeValueOptions.GetValue())
			fmt.Scanln(&inputUser)

			if inputUser == "/O" || inputUser == "/o" {
				fmt.Printf("Language Index: %d, Language: %s\n", JS, "JS")
				fmt.Printf("Language Index: %d, Language: %s\n", RS, "RS")
				fmt.Printf("Language Index: %d, Language: %s\n", CS, "CS")

				fmt.Scanln()
			}

			if val, err := strconv.ParseBool(isDirectoryWritten); err != nil {
				fmt.Printf("%v", val)
			}

			routeValue := RouteInputLayer{
				isLayerComplete: true,

				GetValue: func() string {
					return strings.ToUpper("Enter Name of Directory: ")
				},
			}

			fmt.Println(routeValue.GetValue())
			fmt.Scanln(&inputUser)

			if inputUser == "DIR" {

				routeValue := RouteInputLayer{
					isLayerComplete: true,

					GetValue: func() string {
						return strings.ToUpper("Enter Directory Route | /o: ")
					},
				}

				fmt.Println(routeValue.GetValue())
				fmt.Scanln(&inputUser)

				switch inputUser {
				case "JS":

					l := list.New()
					l.PushFront(routeValue.Get())

					for el := l.Front(); el != nil; el = el.Next() {
						if el.Value == bool(true) {

							os.Mkdir("src", 0755)
							os.Mkdir("images", 0755)
							os.Mkdir("backend", 0755)
							fmt.Println("[DIRECTORY CREATED]: ", "src")
							fmt.Println("[DIRECTORY CREATED]: ", "images")
							fmt.Println("[DIRECTORY CREATED]: ", "images")

						}
					}

				case "RS":
					l := list.New()
					l.PushFront(routeValue.Get())

					for el := l.Front(); el != nil; el = el.Next() {
						if el.Value == bool(true) {

							routeValue := RouteInputLayer{
								isLayerComplete: true,

								GetValue: func() string {
									return strings.ToUpper("Enter Full Stack | /o: ")
								},
							}

							fmt.Println(routeValue.GetValue())
							fmt.Scanln(&inputUser)

							if inputUser == "Fullstack" {

								os.Mkdir("./Frontend/src", 0755)
								os.Mkdir("./img/images", 0755)
								os.Mkdir("./Backend/backend", 0755)
								fmt.Println("[DIRECTORY CREATED]: ", "JS")
							} else {
								os.Mkdir("./src", 0755)
								fmt.Println("[DIRECTORY CREATED]: ", "JS")
							}

						}

					}

				case "CS":
					l := list.New()
					l.PushFront(routeValue.Get())

					for el := l.Front(); el != nil; el = el.Next() {
						if el.Value == bool(true) {
							inputUser := ""

							routeValue := RouteInputLayer{
								isLayerComplete: true,

								GetValue: func() string {
									return strings.ToUpper("Enter Full Stack | /o: ")
								},
							}

							fmt.Println(routeValue.GetValue())
							fmt.Scanln(&inputUser)

							if inputUser == "Fullstack" {

								os.Mkdir("./Frontend/src", 0755)
								os.Mkdir("./img/images", 0755)
								os.Mkdir("./Backend/backend", 0755)
								fmt.Println("[DIRECTORY CREATED]: ", "CS")
							} else {
								os.Mkdir("./src", 0755)
								fmt.Println("[DIRECTORY CREATED]: ", "CS")
							}

						}
					}
				}

				time.Sleep(2 * time.Second)
				fmt.Println("Current Version:", data.GetMesh())

			}

		}
	}
}
