package main

import (
	"os"
	"fmt"
)

func main() {
	write()
}

func write() {
	// write a simple string 
	in := "Hello, World!\n"
	fmt.Println(in)
	file, _ := os.Create("write.txt")
	file.Write([]byte(in))
	file.WriteAt([]byte("And offset!"), int64(len(in)))
	file.Close()
}

func read() {
	// list env
	// fmt.Println("Environment", os.Environ())

	fmt.Println("")
	// if go run, /tmp/go-build630897652/b001/exe/package-os
	// if go build then its the expected location
	str, _ := os.Executable()
	fmt.Println("loc", str)
	fmt.Println("hai")


	// os.Create("hei.txt")
	file, err := os.Open("xanadu.txt")
	if err != nil {
		os.Exit(1)
	}
	buf := make([]byte, 64)
	var eof error
	var read int = 1
	for eof == nil {
		read, err = file.Read(buf)
		// read = t
		eof = err
		fmt.Print(string(buf[:read]))
	}
	// move the pointer to the beginning of the file
	file.Seek(0, 0)
	read, _ = file.Read(buf)
	fmt.Print(string(buf[:read]))
	// fmt.Println(string(buf))
	// fmt.Println(string(buf))
	file.Close()
}