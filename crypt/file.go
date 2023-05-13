package crypt

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func fileByteSize(file *os.File) (int64, error) {
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return stat.Size(), nil
}

func FileToBytes(file *os.File) ([]byte, error) {
	fileSize, err := fileByteSize(file)
	if err != nil {
		return nil, err
	}
	byteSlice := make([]byte, fileSize)
	_, err = bufio.NewReader(file).Read(byteSlice)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return byteSlice, nil
}

func BytesToFile(file *os.File, bs []byte) error {
	_, err := file.Write(bs)
	if err != nil {
		return err
	}
	return nil

}
