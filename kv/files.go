package kv

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveData(filename string, data []byte) error {
	//step 1: create a temp file
	temp, err := os.CreateTemp("", "temp-*.db")
	if err != nil {
		return err

	}
	defer os.Remove(temp.Name())
	//step 2: write data into the temp file
	_, err = temp.Write(data)
	if err != nil {
		return err
	}

	//step 3: fsync the temp file
	err = temp.Sync()
	if err != nil {
		return err
	}
	err = temp.Close()
	if err != nil {
		return err
	}

	//step 4: rename the temp file to original
	err = os.Rename(temp.Name(), filename)
	if err != nil {
		return err
	}

	//step 5: fysnc the directory
	dirName := filepath.Dir(filename)
	if dirName == "" {
		dirName = "."
	}
	dirName, err = (filepath.Abs(dirName))
	if err != nil {
		return err
	}
	fmt.Println("dirName:", dirName)
	dir, err := os.Open(dirName)
	fmt.Println("open error:", err)
	if err != nil {
		return err
	}

	err = dir.Close()
	if err != nil {
		return err
	}
	return nil
}
