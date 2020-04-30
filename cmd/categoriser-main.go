package main

import (
	"fmt"
	"github.com/sheikhusmanshakeel/hello-go2/pkg/categoriser"
	)

func main()  {
	fmt.Println("main called")
	categoriser.Categorise("http:google.co.uk")

}
