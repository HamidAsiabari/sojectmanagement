package main

import (
	"fmt"

	"github.com/HamidAsiabari/sojectmanagement/webserver"
)

func main() {
	fmt.Printf("Soject managment %v startetd", "")
	webs := webserver.Instance()
	webs.Start()
}
