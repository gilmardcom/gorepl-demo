package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gilmardcom/gorepl"
)

func main() {
	fmt.Println()
	fmt.Println("get it now! ---> ", "https://github.com/gilmardcom/gorepl")
	fmt.Println()
	fmt.Println("GoRepl demo started...")
	demo()
	fmt.Println("...GoRepl demo finished.")
}

type SampleJson struct {
	Name     string  `json:"Name"`
	BoolT    bool    `json:"BoolT"`
	BoolF    bool    `json:"BoolF"`
	IntNum   int     `json:"Integer"`
	FloatNum float32 `json:"Flaot"`
	Nulldata string  `json:"Null"`
}

var value = 0
var sampleJson = SampleJson{"gilamrd.com", true, false, 12345, 88.234, "something"}

func demo() {

	gorepl.Config.ConStatus = gorepl.CsDisonnected

	gorepl.AddCommand("quit", "q", "quits the repl", func(args []string) (string, error) {
		gorepl.StopRepl()
		return "", nil
	})

	gorepl.AddCommand("clear", "c", "clears the screen", func(args []string) (string, error) {
		gorepl.Clear()
		return "", nil
	})

	gorepl.AddCommand("help", "h", "shows default help", func(args []string) (string, error) {
		gorepl.DefaultHelp()
		return "", nil
	})

	gorepl.AddCommand("print", "p", "prints the current value", func(args []string) (string, error) {
		fmt.Println("Current value:", value)
		return "", nil
	})

	gorepl.AddCommand("reset", "r", "resets the current value", func(args []string) (string, error) {
		value = 0
		return "", nil
	})

	gorepl.AddCommand("add", "a", "adds the given value to the current value, e.g. add 10", func(args []string) (string, error) {
		if len(args) < 1 {
			return "At least one arg needed", errors.New("invalid number of args")
		}
		for _, item := range args {
			givenValue, _ := strconv.Atoi(item)
			value += givenValue
		}
		return "", nil
	})

	gorepl.AddCommand("sub", "s", "substracts the given value from the current value, e.g. sub 10", func(args []string) (string, error) {
		if len(args) < 1 {
			return "At least one arg needed", errors.New("invalid number of args")
		}
		for _, item := range args {
			givenValue, _ := strconv.Atoi(item)
			value -= givenValue
		}
		return "", nil
	})

	gorepl.AddCommand("json", "j", "prints a sample json", func(args []string) (string, error) {
		gorepl.Json(sampleJson)
		return "", nil
	})

	gorepl.GoRepl()
}
