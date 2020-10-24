package main

import (
	"fmt"
)

// inode
type iNode interface {
	print(string)
	clone() iNode
}

// exact proto
type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() iNode {
	cloneFile := &file{name: f.name + "_clone"}
	return cloneFile
}

// exact proto
type folder struct {
	name      string
	childrens []iNode
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.childrens {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() iNode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildrens []iNode
	for _, i := range f.childrens {
		copy := i.clone()
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.childrens = tempChildrens
	return cloneFolder
}

// main
func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		childrens: []iNode{file1},
		name:      "Folder1",
	}
	folder2 := &folder{
		childrens: []iNode{folder1, file2, file3},
		name:      "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
