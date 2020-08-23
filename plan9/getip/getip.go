package main

import "fmt"

func main() {
	ip := getip()
	fmt.Println(ip)
	ip = getip()
	fmt.Println(ip)
	fmt.Println("some nothing")
	ip = getip()
	fmt.Println(ip)
	foo()
}

func foo() {
	fmt.Println(getip())
}

func getip() int
