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

	sub := ""
	for i, subject := range p.Rules[0].Target.Subjects.Subjects {
		sub += subject.SubjectMatch.AttributeValue.Value
		if i != len(p.Rules[0].Target.Subjects.Subjects) - 1 {
			sub += ", "
		}
	}
	fmt.Println(sub)

	obj := ""
	for i, object := range p.Rules[0].Target.Resources.Resources {
		obj += object.ResourceMatch.AttributeValue.Value
		if i != len(p.Rules[0].Target.Resources.Resources) - 1 {
			obj += ", "
		}
	}
	fmt.Println(obj)

	act := ""
	for i, subject := range p.Rules[0].Target.Actions.Actions {
		act += subject.ActionMatch.AttributeValue.Value
		if i != len(p.Rules[0].Target.Actions.Actions) - 1 {
			act += ", "
		}
	}
	fmt.Println(act)
}

func ParseRequest(path string) *Request {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()

	var p Request
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

func PrintRequest(r *Request) {
	sub := r.Subject.Attribute.AttributeValue
	obj := r.Resource.Attribute.AttributeValue
	act := r.Action.Attribute.AttributeValue

	fmt.Println(sub + ", " + obj + ", " + act)
}
