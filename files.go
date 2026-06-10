package kv

import (
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

	//step 4: rename the temp file to original
	err = os.Rename(temp.Name(), filename)
	if err != nil {
		return err
	}
	//step 5: fysnc the directory
	dir, err := os.Open(filepath.Dir(filename))
	if err != nil {
		return err
	}
	err = dir.Sync()
	if err != nil {
		return err
	}
	err = dir.Close()
	if err != nil {
		return err
	}
	return nil
}
