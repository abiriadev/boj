package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.DateOnly))
}
