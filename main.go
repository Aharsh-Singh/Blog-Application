package main

import (
	"fmt"
	"myapp/Routes"
)

func main() {
	r := Routes.SetupRouter()
    fmt.Println("Server is running on port 8080")
    r.Run(":8080")
}
