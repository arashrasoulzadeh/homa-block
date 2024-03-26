package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func SyncCollection(filePath string, collection interface{}) {
	// Check if the file already exists
	_, err := os.Stat(filePath)

	// If the file doesn't exist, create it
	if os.IsNotExist(err) {
		createFile(filePath, collection)

	} else if err != nil { // If there's another error (besides non-existence), handle it
		fmt.Println("Error:", err)
		return
	} else {
		updateFile(filePath, collection)
	}
}

// Function to create the file with initial data
func createFile(filePath string, data interface{}) {
	// Create a sample data structure

	// Marshal the data to JSON format
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("JSON file created:", filePath)
}

func updateFile(filePath string, data interface{}) {
	// Read the existing JSON data from the file
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Update the data

	// Marshal the updated data back to JSON format
	updatedJsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling updated JSON:", err)
		return
	}

	// Write the updated JSON data back to the file
	err = ioutil.WriteFile(filePath, updatedJsonData, 0644)
	if err != nil {
		fmt.Println("Error writing updated JSON data to file:", err)
		return
	}

	fmt.Println("JSON file updated:", filePath)
}

func ReadFile(filePath string) (interface{}, error) {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return nil, err
	}

	// Unmarshal the JSON data into a map
	var data interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return data, nil

}

func CreateDirectoryIfNotExists(directoryPath string) error {
	// Check if the directory already exists
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		// If the directory doesn't exist, create it
		err := os.MkdirAll(directoryPath, 0755) // 0755 is the permission mode
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
		fmt.Println("Directory created:", directoryPath)
	} else if err != nil {
		fmt.Println("Error:", err)
		return err
	} else {
		fmt.Println("Directory already exists:", directoryPath)
	}
	return nil
}

func ListDirectories(dirPath string) []string {

	// Get a list of all entries in the directory
	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Iterate over each entry and filter out directories
	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs
}
func ListFiles(dirPath string) []string {

	// Get a list of all entries in the directory
	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Iterate over each entry and filter out files
	var dirs []string
	for _, entry := range entries {
		if !entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs
}
