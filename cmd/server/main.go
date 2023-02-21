package main

import "fmt"

// Run - responsible for the instantiation
// and startup of our application
func Run() error {
	fmt.Println("starting our application")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
