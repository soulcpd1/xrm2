package xray

import (
	"os"
)

func Transaction(tempFile string, data []byte) error {

	if err := Write(tempFile, data); err != nil {
		return err
	}

	if err := Validate(tempFile); err != nil {
		os.Remove(tempFile)
		return err
	}

	return nil
}
