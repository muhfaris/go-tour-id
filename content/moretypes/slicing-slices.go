// +build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// tanpa ada index kiri berarti 0
	fmt.Println("s[:3] ==", s[:3])

	// tanpa ada index di sisi kanan berarti len(s)
	fmt.Println("s[4:] ==", s[4:])
}
