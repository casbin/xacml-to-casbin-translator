package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Hasdcorona/go-xacml/pdp"
)

func ParsePolicy(path string) *pdp.Policy {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()

	var p pdp.Policy
	s, err := xmlFile.Stat()
	if err != nil {
		fmt.Println("Error stating file:", err)
		return nil
	}
	size := s.Size()
	b := make([]byte, size)
	xmlFile.Read(b)
	err = xml.Unmarshal([]byte(b), &p)
	return &p
}

func PrintPolicy(p *pdp.Policy) {
	// fmt.Printf("%+v\n", p)
	fmt.Print(p.Rules[0].Target.Subjects.Subjects[0].SubjectMatch.AttributeValue.Value)
}