package model

import "encoding/xml"

type Node struct {
	XMLName     xml.Name `xml:"student"`
	StudentName string   `xml:"studentName"`
	StudentId   string   `xml:"studentId"`
}

type Table struct {
	XMLName     xml.Name  `xml:"table"`
	Node   []Node `xml:"node"`
	Description string    `xml:",innerxml"`
}