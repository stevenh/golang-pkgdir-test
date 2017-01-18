package main

import (
	"fmt"

	"github.com/stevenh/golang-pkgdir-test/pkg1"
)

func main() {
	t := pkg1.TODO()
	fmt.Println(t)
}
