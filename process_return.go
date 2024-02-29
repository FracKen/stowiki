package stowiki

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// processAPICall is a function that takes in data of type pointer to STOWiki and a slice of strings (tables).
// It processes an API call for each table in the tables slice.
func ProcessAPICall(data *STOWiki, tables []string) {
	// Loop through each table in the tables slice.
	for _, table := range tables {
		// Initialize an empty byte slice for the body of the HTTP response.
		body := []byte{}

		// Create a query URL for the current table.
		url := createQuery(table)

		// Make an HTTP request to the query URL.
		// If an error occurs, print the error and continue to the next iteration of the loop.
		res, err := makeRequest(url)
		if err != nil {
			fmt.Printf("Error making the request for table %s: %v\n", table, err)
			continue
		}

		// Ensure the response body is closed after the function returns.
		defer res.Body.Close()

		// Read the entire response body.
		// If an error occurs, print the error and continue to the next iteration of the loop.
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading the response body for table %s: %v\n", table, err)
			continue
		}

		// If the current table is "Infobox", unmarshal the response body into the Equipment field of data.
		// If an error occurs, print the error, create a file with the error and the table name,
		// and write the response body to the file.
		if table == "Infobox" {
			err = json.Unmarshal(body, &data.Equipment)
			if err != nil {
				handleUnmarshalError(err, table, res)
			}
		}

		// Similar handling for the "Traits" table, but unmarshal into the PersonalTraits field of data.
		if table == "Traits" {
			err = json.Unmarshal(body, &data.PersonalTraits)
			if err != nil {
				handleUnmarshalError(err, table, res)
			}
		}

		// Similar handling for the "StarshipTraits" table, but unmarshal into the StarshipTraits field of data.
		if table == "StarshipTraits" {
			err = json.Unmarshal(body, &data.StarshipTraits)
			if err != nil {
				handleUnmarshalError(err, table, res)
			}
		}
	}
}

// handleUnmarshalError is a helper function to handle errors during unmarshalling.
// It creates a file with the error and the table name, and writes the response body to the file.
func handleUnmarshalError(err error, table string, res *http.Response) {
	fmt.Printf("failed with error: %v\n", err)
	file, err := os.Create(fmt.Sprintf("error-%s.json", table))
	if err != nil {
		fmt.Printf("Error creating file for table %s: %v\n", table, err)
		return
	}
	defer file.Close()

	// Write the response body to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Printf("Error writing the response body to a file for table %s: %v\n", table, err)
	}
}

// processAPICall is a function that processes API calls.
// It takes in two parameters:
// 'data' which is a pointer to a STOWiki struct
// 'tables' which is a slice of strings.
// The function doesn't return any value. Only the 'data' parameter is modified via he pointer.
// func processAPICall(data *STOWiki, tables []string) {
// 	for _, table := range tables {
// 		body := []byte{}
// 		url := createQuery(table)
// 		res, err := makeRequest(url)
// 		if err != nil {
// 			fmt.Printf("Error making the request for table %s: %v\n", table, err)
// 			continue
// 		}
// 		defer res.Body.Close()
// 		body, err = io.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Printf("Error reading the response body for table %s: %v\n", table, err)
// 			continue
// 		}
// 		if table == "Infobox" {
// 			err = json.Unmarshal(body, &data.Equipment)
// 			if err != nil {
// 				fmt.Printf("failed with error: %v\n", err)
// 				file, err := os.Create(fmt.Sprintf("error-%s.json", table))
// 				if err != nil {
// 					fmt.Printf("Error creating file for table %s: %v\n", table, err)
// 					continue
// 				}
// 				defer file.Close()

// 				// Write the response body to the file
// 				_, err = io.Copy(file, res.Body)
// 				if err != nil {
// 					fmt.Printf("Error writing the response body to a file for table %s: %v\n", table, err)
// 					continue
// 				}
// 			}
// 		}

// 		if table == "Traits" {
// 			err = json.Unmarshal(body, &data.PersonalTraits)
// 			if err != nil {
// 				fmt.Printf("failed with error: %v\n", err)
// 				file, err := os.Create(fmt.Sprintf("error-%s.json", table))
// 				if err != nil {
// 					fmt.Printf("Error creating file for table %s: %v\n", table, err)
// 					continue
// 				}
// 				defer file.Close()

// 				// Write the response body to the file
// 				_, err = io.Copy(file, res.Body)
// 				if err != nil {
// 					fmt.Printf("Error writing the response body to a file for table %s: %v\n", table, err)
// 					continue
// 				}
// 			}
// 		}
// 		if table == "StarshipTraits" {
// 			err = json.Unmarshal(body, &data.StarshipTraits)
// 			if err != nil {
// 				fmt.Printf("failed with error: %v\n", err)
// 				file, err := os.Create(fmt.Sprintf("error-%s.json", table))
// 				if err != nil {
// 					fmt.Printf("Error creating file for table %s: %v\n", table, err)
// 					continue
// 				}
// 				defer file.Close()

// 				// Write the response body to the file
// 				_, err = io.Copy(file, res.Body)
// 				if err != nil {
// 					fmt.Printf("Error writing the response body to a file for table %s: %v\n", table, err)
// 					continue
// 				}
// 			}
// 		}
// 	}
// }
