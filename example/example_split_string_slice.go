package main

import (
	"SplitSlice/SplitSlice"
	"fmt"
)

func main() {
	data := []string{"a", "2", "3"}
	fmt.Println(SplitSlice.SplitStringSlice(data, 2))
}
