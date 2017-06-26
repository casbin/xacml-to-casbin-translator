package main

import (
	"fmt"
	"strings"
)

var Cases []string

func InitCases() {
	for i := 1; i < 22; i++ {
		Cases = append(Cases, fmt.Sprintf("IIA%03d", i))
	}

	for i := 1; i < 54; i++ {
		Cases = append(Cases, fmt.Sprintf("IIB%03d", i))
	}

	for i := 1; i < 233; i++ {
		if i != 23 && i != 54 && i != 55 && i != 88 && i != 89 && i != 92 && i != 93 && i != 98 && i != 99 {
			Cases = append(Cases, fmt.Sprintf("IIC%03d", i))
		}
	}

	for i := 1; i < 31; i++ {
		Cases = append(Cases, fmt.Sprintf("IID%03d", i))
	}

	for i := 1; i < 4; i++ {
		Cases = append(Cases, fmt.Sprintf("IIE%03d", i))
	}

	for i := 1; i < 29; i++ {
		Cases = append(Cases, fmt.Sprintf("IIIA%03d", i))
	}

	for i := 1; i < 4; i++ {
		Cases = append(Cases, fmt.Sprintf("IIIC%03d", i))
	}

	for i := 1; i < 8; i++ {
		Cases = append(Cases, fmt.Sprintf("IIIF%03d", i))
	}

	for i := 1; i < 7; i++ {
		Cases = append(Cases, fmt.Sprintf("IIIG%03d", i))
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

	fmt.Printf("\nNumber of cases: %d\n\n", len(Cases))
}

func main() {
	InitCases()
	PrintCases()

	for _, c := range Cases {
		path := fmt.Sprintf("ConformanceTests/%sPolicy.xml", c)
		if res, _ := PathExists(path); !res {
			path = strings.Replace(path, ".xml", "1.xml", -1)
		}
		p := ParsePolicy(path)
		fmt.Print(c + ", ")
		PrintPolicy(p)
		fmt.Print("\n")
	}

	//for _, c := range Cases {
	//	r := ParseRequest(fmt.Sprintf("ConformanceTests/%sRequest.xml", c))
	//	fmt.Print(c + ", ")
	//	PrintRequest(r)
	//
	//	res := ParseResponse(fmt.Sprintf("ConformanceTests/%sResponse.xml", c))
	//	fmt.Print(" ----> ")
	//	PrintResponse(res)
	//	fmt.Print("\n")
	//}
}