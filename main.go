package main

import (
	"fmt"
)

func main() {
	variadF()

}

func variadF() {
	fs := []int{9, 7, 4, 2, 1}
	ts := []int{9, 7, 4, 2, 1}
	fs = append(fs, ts...)
	m := (*[10]int)(fs)
	Shaw(m)
}
func Shaw(v *[10]int) {
	for _, v := range *v {
		defer fmt.Println(v)
	}
}
