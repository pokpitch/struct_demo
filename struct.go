package main

import "fmt"

type rectangle struct {
	w, l float64
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}
	rec.w = 6

	fmt.Println(rec.l)
	fmt.Println(rec.w)

}
