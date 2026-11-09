package pages

import "fmt"

type Page struct {
	data [8192]byte
	count int
}

func Demo() {
	var p Page
	fmt.Println(p.count, len(p.data))
}