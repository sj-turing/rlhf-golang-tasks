package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Define a recursive struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}

	type node struct {
		Val int `json:"val"`
	}

	data := node{n.Val}
	if n.Left != nil || n.Right != nil {
		data.Left = n.Left
		data.Right = n.Right
	}

	return json.Marshal(data)
}

func (n *TreeNode) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n = nil
		return nil
	}

	type node struct {
		Val   int       `json:"val"`
		Left  *TreeNode `json:"left"`
		Right *TreeNode `json:"right"`
	}

	var d node
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	n = &TreeNode{d.Val}
	n.Left = d.Left
	n.Right = d.Right
	return nil
}

func (n *TreeNode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if n == nil {
		return nil
	}

	start.Name.Local = "node"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if err := e.EncodeElement(&n.Val, xml.Name{Local: "val"}); err != nil {
		return err
	}

	if n.Left != nil {
		if err := e.EncodeElement(n.Left, xml.Name{Local: "left"}); err != nil {
			return err
		}
	}

	if n.Right != nil {
		if err := e.EncodeElement(n.Right, xml.Name{Local: "right"}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (n *TreeNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Local != "node" {
		return fmt.Errorf("expected <node>, got %s", start.Name.Local)
	}

	n = &TreeNode{}

	for {
		t, err := d.Token()
		if err != nil {
			return err
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "val":
				if err := d.DecodeElement(&n.Val, &se); err != nil {
					return err
				}
			case "left":
				if err := d.DecodeElement(&n.Left, &se); err != nil {
					return err
				}
			case "right":
				if err := d.DecodeElement(&n.Right, &se); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unexpected element: %s", se.Name.Local)
			}
		case xml.EndElement:
			if se.Name == start.Name {
				return nil
			}
			return fmt.Errorf("unexpected end element: %s", se.Name.Local)
		default:
			return fmt.Errorf("unexpected token: %T", t)
		}
	}
}

func main() {
	// Example data
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 5},
		},
	}

	// JSON serialization
	jsonData, err := json.Marshal(root)
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}
	fmt.Println("JSON data:", string(jsonData))

	// JSON deserialization
	var jsonRoot TreeNode
	if err := json.Unmarshal(jsonData, &jsonRoot); err != nil {
		log.Fatalf("Error unmarshaling from JSON: %v", err)
	}
	fmt.Println("JSON deserialized root:", jsonRoot)

	// XML serialization
	xmlData, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to XML: %v", err)
	}
	fmt.Println("XML data:", string(xmlData))

	// XML deserialization
	var xmlRoot TreeNode
	if err := xml.Unmarshal(xmlData, &xmlRoot); err != nil {
		log.Fatalf("Error unmarshaling from XML: %v", err)
	}
	fmt.Println("XML deserialized root:", xmlRoot)

	// Save to file and load back
	if err := ioutil.WriteFile("tree.json", jsonData, 0644); err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}

	jsonFromFile, err := ioutil.ReadFile("tree.json")
	if err != nil {
		log.Fatalf("Error reading JSON from file: %v", err)
	}

	var fileRoot TreeNode
	if err := json.Unmarshal(jsonFromFile, &fileRoot); err != nil {
		log.Fatalf("Error unmarshaling JSON from file: %v", err)
	}
	fmt.Println("JSON deserialized from file root:", fileRoot)

	if err := ioutil.WriteFile("tree.xml", xmlData, 0644); err != nil {
		log.Fatalf("Error writing XML to file: %v", err)
	}

	xmlFromFile, err := ioutil.ReadFile("tree.xml")
	if err != nil {
		log.Fatalf("Error reading XML from file: %v", err)
	}

	if err := xml.Unmarshal(xmlFromFile, &fileRoot); err != nil {
		log.Fatalf("Error unmarshaling XML from file: %v", err)
	}
	fmt.Println("XML deserialized from file root:", fileRoot)
}
