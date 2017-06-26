package main

func main() {
	p := ParsePolicy("ConformanceTests/IIA001Policy.xml")
	PrintPolicy(p)
}