package main

import (
	"fmt"

	"github.com/h-masano/go-workspace-sample/workspace-1/module-b"
)

func main() {
	b := modb.ModB()
	fmt.Println("run moda, import ", b)
}