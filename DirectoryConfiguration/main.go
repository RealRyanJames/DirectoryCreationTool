package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func main() {

	routeValue := RouteInputLayer{
		isLayerComplete: true,

		GetValue: func() string {
			return strings.ToUpper("Enter Name of Directory: ")
		},
	}

	if val, err := strconv.ParseBool(isDirectoryWritten); err != nil {
		fmt.Printf("%v", val)
	}

	inputUser := ""
	fmt.Println(routeValue.GetValue())
	fmt.Scanln(&inputUser)

	dir := DIRECTORY{
		Name:        inputUser,
		isWrittenTo: string("true"),
	}

	strconv.ParseBool(dir.isWrittenTo)
	err := os.MkdirAll(dir.GetDirectory(), 0755)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[DIRECTORY CREATED]: ", dir.GetDirectory())
}
