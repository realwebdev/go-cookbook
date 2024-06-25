package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the JSON input file
	inputFile, err := os.Open("foundationDownload.json")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}

	defer inputFile.Close()

	// Read the infput file content

	byteValue, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Declare a map to stroe the decoded JSON data
	var data map[string]interface{}

	// Unmarshal JSON content into the data map
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Marshal the data backe to an indented JSON byte slice
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write the formatted JSON to the output file
	outputFile, err := os.Create("indented.json")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.Write(prettyJSON)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Formatted JSON has been written to output.json")
}
