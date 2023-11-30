// main.go

package main

import (
	"fmt"
	"github.com/Radha-Sundar/godep-app/pkg/greeter"
)

func main() {
	message := greeter.GetGreeting()
	fmt.Println(message)
}
