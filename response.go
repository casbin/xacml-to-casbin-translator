package main

import "encoding/xml"

type Response struct {
	XMLName                xml.Name       `xml:"Response"`
	Result                 Result         `xml:"Result"`
}

type Result struct {
	XMLName                xml.Name       `xml:"Result"`
	Decision               string         `xml:"Decision"`
	Status                 Status         `xml:"Status"`
}

type Status struct {
	XMLName                xml.Name       `xml:"Status"`
	StatusCode             StatusCode     `xml:"StatusCode"`
}

type StatusCode struct {
	XMLName                xml.Name       `xml:"StatusCode"`
	Value                  string         `xml:"Value,attr"`
}
