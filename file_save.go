package stowiki

import (
	"encoding/json"
	"fmt"
	"os"
)

// func saveToFile(filename string, data STOWiki) error {

// 	newBody, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshalling the JSON:", err)
// 		return err
// 	}

// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(file, strings.NewReader(string(newBody)))
// 	return err
// }

// saveToFile is a function that takes a filename and data of type STOWiki.
// It marshals the data to JSON and writes it to a file with the given filename.
// If an error occurs during this process, it returns the error.
func SaveToFile(filename string, data STOWiki) error {
	// Marshal the data into JSON, with indents for readability.
	// If an error occurs, print an error message and return the error.
	newBody, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error marshalling the JSON:", err)
		return err
	}

	// Create a new file with the given filename.
	// If an error occurs, return the error.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Ensure the file is closed after the function returns.
	defer file.Close()

	// Write the JSON data to the file.
	// If an error occurs, it will be returned.
	_, err = file.Write(newBody)
	return err
}
