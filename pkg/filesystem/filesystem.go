package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(folder, file, extension string) (*os.File, error) {
	if _, err := os.Stat(filepath.Join(folder, file+extension)); err == nil {
		return nil, fmt.Errorf("filesystem: '%s/%s' already exists", folder, file+extension)
	}

	f, err := os.Create(filepath.Join(folder, file+extension))
	if err != nil {
		panic(err)
	}

	return f, nil
}

func OpenFile(folder, file, extension string) (*os.File, error) {
	if _, err := os.Stat(filepath.Join(folder, file+extension)); err != nil {
		return nil, fmt.Errorf("filesystem: '%s/%s' already exists", folder, file+extension)
	}

	f, err := os.Open(filepath.Join(folder, file+extension))
	if err != nil {
		panic(err)
	}

	return f, nil
}

func OverrideFile(folder, file, extension string) (*os.File, error) {
	if _, err := os.Stat(filepath.Join(folder, file+extension)); err != nil {
		return nil, fmt.Errorf("filesystem: '%s' does not exists", folder+file+extension)
	}

	f, err := os.Create(filepath.Join(folder, file+extension))
	if err != nil {
		panic(err)
	}

	return f, nil
}

func DeleteFile(folder, file, extension string) error {
	if _, err := os.Stat(filepath.Join(folder, file+extension)); err != nil {
		return fmt.Errorf("filesystem: '%s' does not exists", folder+file+extension)
	}

	//TODO: Verify removes file in directory
	err := os.Remove(folder + file + extension)
	if err != nil {
		panic(err)
	}

	return nil
}
