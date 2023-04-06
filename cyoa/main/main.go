package main

import (
	"fmt"
	"os"

	"github.com/nekidb/gophercises/cyoa/stories"
)

func main() {
	storiesList, err := stories.GetStoriesFromFile(os.DirFS("."), "stories.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(storiesList)
}
