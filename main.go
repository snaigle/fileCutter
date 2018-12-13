package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		_usage()
		return
	}
	fileName := os.Args[1]
	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("index must be number,and match the condition: >0 and < file.size", err)
		return
	}
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file", fileName, "error", err)
		return
	}
	fs, err := f.Stat()
	if err != nil {
		fmt.Println("read file", fileName, "error", err)
		return
	}
	if index <= 0 || int64(index) >= fs.Size() {
		fmt.Println("index must be >0 and < file.size")
		return
	}
	f1, err := os.OpenFile(fileName+"-1", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("create file", fileName+"-1", "failed", err)
	}
	f2, err := os.OpenFile(fileName+"-2", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("create file", fileName+"-2", "failed", err)
	}
	buf := make([]byte, 4196)
	for {
		count, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read file", fileName, "error", err)
		}
		if index > 0 {
			if count <= index {
				f1.Write(buf)
				index -= count
			} else {
				f1.Write(buf[0:index])
				f2.Write(buf[index:count])
				index = 0
			}
		} else {
			f2.Write(buf[0:count])
		}
		if err == io.EOF || count == 0 {
			break
		}
	}
	f1.Sync()
	f2.Sync()
	fmt.Println("success")
}

func _usage() {
	fmt.Println("usage: cutter fileName index")
}
