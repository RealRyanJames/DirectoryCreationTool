package main

import (
	"container/list"
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

func (route RouteInputLayer) Get() bool {
	if route.isLayerComplete {
		return true
	}

	return false
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

	l := list.New()
	l.PushFront(routeValue.Get())

	for el := l.Front(); el != nil; el = el.Next() {

		if el.Value == bool(true) {

			os.Mkdir(dir.GetDirectory(), 0755)

			fmt.Println("[DIRECTORY CREATED]: ", dir.GetDirectory())
		}
	}

}
