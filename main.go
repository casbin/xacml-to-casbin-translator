package main

import "fmt"

var Cases []string

func InitCases() {
	for i := 1; i < 2; i++ {
		Cases = append(Cases, fmt.Sprintf("IIA%03d", i))
	}
}

func PrintCases() {
	fmt.Println("Cases:")

	for i, c := range Cases {
		fmt.Print(c)
		if i != len(Cases) - 1 {
			fmt.Print(", ")
		}
	}

	fmt.Print("\n\n")
}

func main() {
	InitCases()
	PrintCases()

	//p := ParsePolicy("ConformanceTests/IIA001Policy.xml")
	//PrintPolicy(p)

	for _, c := range Cases {
		r := ParseRequest(fmt.Sprintf("ConformanceTests/%sRequest.xml", c))
		fmt.Print(c + ", ")
		PrintRequest(r)
	}
}