package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写入

// 1. []byte方式写入
func writeFromByte(){
	fileObj,err := os.OpenFile("./error.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	defer fileObj.Close()
	if err != nil{
		fmt.Println("打开写入文件时发生错误")
		return
	}
	n,err:=fileObj.WriteString("abcdefg\n")
	if err !=nil{
		fmt.Println("写入字符错误")
		return
	}
	fmt.Println("写入的字节数为: ",n)
}

// 2. bufio方式写入
func writeFromBufio(){
	fileObj,err := os.OpenFile("./error.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	defer fileObj.Close()
	writer:=bufio.NewWriter(fileObj)
	n,err:=writer.WriteString("bufio write\n")
	if err !=nil{
		fmt.Println("写入字符错误")
		return
	}
	fmt.Println("写入的字节数为: ",n)
	writer.Flush()
}

// 3. ioutil方式写入, 似乎无法追加写入
func writeFromIoutil(){
	err:=ioutil.WriteFile("./error.txt",[]byte("write from ioutil"),0644)
	if err != nil{
		fmt.Println("写入字符错误")
		return
	}
}

func main(){
	//writeFromByte()
	//writeFromBufio()
	writeFromIoutil()
}