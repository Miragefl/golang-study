package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("a.txt")
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(f.Name())
}

func makeDir() {
	err := os.Mkdir("test", os.ModePerm)
	if nil != err {
		fmt.Println(err)
	}

	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if nil != err2 {
		fmt.Println(err2)
	}
}

func delFile() {
	err := os.Remove("a.txt")
	if nil != err {
		fmt.Println(err)
	}

	err2 := os.RemoveAll("a")
	if nil != err2 {
		fmt.Println(err2)
	}
	err3 := os.RemoveAll("test")
	if nil != err3 {
		fmt.Println(err3)
	}
}

func wd() {
	dir, _ := os.Getwd()
	fmt.Println("dir :", dir)

	os.Chdir("d:/")
	dir2, _ := os.Getwd()
	fmt.Println("dir2 :", dir2)

	dir3 := os.TempDir()
	fmt.Println("dir3 :", dir3)
}

func rename() {
	err := os.Rename("text.txt", "text2.txt")
	if nil != err {
		fmt.Println(err)
	}
}

func readFile() {
	b, _ := os.ReadFile("text2.txt")
	fmt.Println(string(b))
}

func writeFile() {
	s := "hello world"
	os.WriteFile("text2.txt", []byte(s), os.ModePerm)
}

func main() {
	//createFile()
	//makeDir()
	//delFile()
	//wd()
	//rename()
	//readFile()
	writeFile()
}
