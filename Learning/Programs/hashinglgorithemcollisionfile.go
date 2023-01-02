package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	file1, err := os.Open("erase.exe")

	if err != nil {
		panic(err)
	}

	defer file1.Close()

	hash1 := md5.New()
	_, err = io.Copy(hash1, file1)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s MD5 checksum is %x \n", file1.Name(), hash1.Sum(nil))

	file2, err := os.Open("hello.exe")

	if err != nil {
		panic(err)
	}

	defer file2.Close()

	hash2 := md5.New()
	_, err = io.Copy(hash2, file2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s MD5 checksum is %x \n", file2.Name(), hash2.Sum(nil))
}

/*
go run hashinglgorithemcollisionfile.go
erase.exe MD5 checksum is cdc47d670159eef60916ca03a9d4a007
hello.exe MD5 checksum is cdc47d670159eef60916ca03a9d4a007

// Download the files from the https://www.mscs.dal.ca/~selinger/md5collision/
*/
