package creator

import "os"

func directoryCreator(dirPath string) error {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}
