package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"text/template"
)

const maxDepth = 64 // Maximum allowed depth

// Experimental Decoding with depth limited
func customDecode(data interface{}, depth int) interface{} {
	if depth > maxDepth {
		panic(fmt.Sprintf("Maximum depth exceeded: %d", depth))
	}

	switch v := data.(type) {
	case nil:
		return nil
	case int:
		// Base case: return primitive types as is
		return v
	case *Node:
		// Handle Node struct with recursive Next pointer
		// Decode the Value field
		v.Value = customDecode(v.Value, depth+1).(int)

		// Recursively decode the Next field (if it exists)
		if v.Next != nil {
			v.Next = customDecode(v.Next, depth+1).(*Node)
		}

		return v
	default:
		// If it's an unsupported type, return as is
		return v
	}
}

// Define a recursive structure
type Node struct {
	Value int
	Next  *Node
}

// Handler for Home method
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

// Checks if a string is base64 encoded
func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

// Function to traverse and print the values in the structure
func getDepthAndDecodeData(node *Node) (int, string) {
	current := node
	var stringBuilder2 strings.Builder
	count := 0

	for current != nil && count < maxDepth {
		stringBuilder2.WriteString(fmt.Sprintf("[Node %d - Value %d] =>", count, current.Value))
		current = current.Next
		count++
	}
	stringBuilder2.WriteString("nil")
	return count, stringBuilder2.String()
}

// Decodes the string from base64 encoding and sanitises it to prevent XSS
func getSanitisedString(input string) string {
	base64DecodedInput, _ := base64.StdEncoding.DecodeString(input)
	base64DecodedString := html.EscapeString(string(base64DecodedInput))
	return base64DecodedString
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	userInput := r.FormValue("userInput")
	safeInput := html.EscapeString(userInput)
	var sb strings.Builder

	if !IsBase64(safeInput) {
		sb.WriteString("Wasn't it clear how much I love BASE64 absolutely EVERYWHERE! 😠")
		w.Write([]byte(sb.String()))
	} else {
		sanitisedString := getSanitisedString(safeInput)

		//TODO: REMOVE THIS API KEY
		sb.WriteString("fzb nhc rwjr sxkzagr nhcz qnl piq sgp? ycreiaj kpdm awic nhc EVA4zUeGAYujKssfA2cApfQgPOnxFvpuJSwLluL3GR0=")

		// This runs at the end. In the case of a panic, you get steps to the next part
		defer func() {
			if r := recover(); r != nil {
				sb.WriteString(`<script>console.log("/flagify");</script>`)
				sb.WriteString(`<img src="/assets/gophervignere.png" alt="Recovery Image" style="max-width:100%; height:30%;">`)
				w.Write([]byte(sb.String()))
			} else {
				w.Write([]byte(sb.String()))
			}
		}()

		//Standard Decoding (just to be safe)
		var decodedNode Node
		var buf bytes.Buffer
		buf.WriteString(sanitisedString)
		decoder := gob.NewDecoder(&buf)
		err := decoder.Decode(&decodedNode)

		if err != nil {
			fmt.Println(err)
			sb.Reset()
			sb.WriteString("Your Base 64 Decoded Input: ")
			sb.WriteString(string(sanitisedString))
		} else {
			decodedDataWithDepthLimit := customDecode(&decodedNode, 0)
			sb.Reset()
			current, ok := decodedDataWithDepthLimit.(*Node)
			if !ok {
				sb.WriteString("Not sure what you did, check the console maybe?")
				sb.WriteString(`<script>console.log("Error: Expected *Node")</script>`)
			} else {
				depth, outputString := getDepthAndDecodeData(current)
				if 0 < depth && depth < 16 {
					sb.WriteString("I can handle that without PANICking")
				} else if depth >= 16 && depth < 32 {
					sb.WriteString("Not even half PANICked")
				} else if depth >= 32 && depth < 48 {
					sb.WriteString("Bearable, just don't over do it")
				} else if depth >= 48 && depth < 64 {
					sb.WriteString("Okay that's it, don't nest it anymore")
				}
				sb.WriteString(fmt.Sprintf(". Anyways, Your Decoded Input: {%v}", outputString))
			}
		}
	}
}

func flagHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML page
	tmpl := template.Must(template.ParseFiles("flagServer.html"))
	tmpl.Execute(w, nil)
}

func flagAuth(w http.ResponseWriter, r *http.Request) {
	userInput := r.FormValue("userInput")
	safeInput := html.EscapeString(userInput)

	var sb strings.Builder
	if safeInput == "fzb nhc rwjr sxkzagr nhcz qnl piq sgp? ycreiaj kpdm awic nhc EVA4zUeGAYujKssfA2cApfQgPOnxFvpuJSwLluL3GR0=" {
		sb.WriteString("🏁 : hackcenter{FLAGE}")
	} else {
		sb.WriteString("🏁 : aHR0cHM6Ly93d3cueW91dHViZS5jb20vd2F0Y2g/dj1kUXc0dzlXZ1hjUQ==")
	}
	w.Write([]byte(sb.String()))
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/c3VibWl0", submitHandler)
	http.HandleFunc("/ZmxhZ2lmeQ==", flagHandler)
	http.HandleFunc("/ZmxhZ0F1dGg=", flagAuth)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
