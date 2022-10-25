package helper

import (
	"fmt"
	"io"
	"os"
)

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		err := inputFile.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)
	_, err = io.Copy(outputFile, inputFile)
	err = inputFile.Close()
	if err != nil {
		return err
	}
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}
