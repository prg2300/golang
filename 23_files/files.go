package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("exmple.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fileinfo, err := f.Stat()
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Println("file name", fileinfo.Name())
	// fmt.Println("file or folder", fileinfo.IsDir())
	// fmt.Println("file size", fileinfo.Size())
	// fmt.Println("file permission", fileinfo.Mode())
	// fmt.Println("file modified at", fileinfo.ModTime())

	//read file
	// f, err := os.Open("exmple.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// buf := make([]byte, 12)

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	//fmt.Println("data", d, buf)

	// data, err := os.ReadFile("exmple.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())

	// }

	// f, err := os.Create("exmple2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	//  f.WriteString("hi go")
	// f.WriteString("pranal")

	// bytes := []byte("hello golang")  //replace

	// f.Write(bytes)

	//read and write to another file

	// sourceFile, err := os.Open("exmple.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("exmple2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	// fmt.Println("written to new file successfully")

	//delete file

	err := os.Remove("exmple2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted successfully")

}
