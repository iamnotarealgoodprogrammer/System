package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	//ErrBadArgs ... Err parsing args of the command line
	ErrBadArgs = errors.New("Err parsing arguments of the cmd")
)

func main() {

	args, err := checkArgs()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	str, err := getInfo(args[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Println(str)

}

func checkArgs() ([]string, error) {
	args := os.Args
	if len(args) == 1 {
		return nil, ErrBadArgs
	}
	return args[1:], nil
}

func getInfo(path string) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		log.Printf("%s is a symbolic link", path)
		realPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return "", err
		}
		log.Printf("RealPath: %s", realPath)
		return fmt.Sprintf("RealPath: %s", realPath), nil
	}
	return "It is not a symbolic link", nil
}
