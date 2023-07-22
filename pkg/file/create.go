package file

import (
	"fmt"
	"os"
)

func Create(pathFileExtName string, content string) error {
	// The name of the file to create.
	fileName := pathFileExtName

	// Create the file.
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("File", fileName, "created and written successfully.")
	return nil
}
