package dev

import (
	"fmt"
	"os"
)

func SpoofFiles() {
	fmt.Println("Spoofing files")
	os.WriteFile("/mnt/f/codeTesting/test.txt", []byte("test"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test2.txt", []byte("test2"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test3.txt", []byte("test3"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test4.txt", []byte("test4"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test5.txt", []byte("test5"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test6.txt", []byte("test6"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test7.txt", []byte("test7"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test8.txt", []byte("test8"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test9.txt", []byte("test9"), 0644)
	os.WriteFile("/mnt/f/codeTesting/test10.txt", []byte("test10"), 0644)
}
