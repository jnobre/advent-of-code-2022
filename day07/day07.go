package main

import (
	"fmt"
	"strings"

	AOCUTILS "github.com/jnobre/advent-of-code-2022"
)

type FileType uint8

const (
	FILE FileType = iota
	DIR
)

type File struct {
	Type     FileType
	Name     string
	Size     int
	Children []*File
	Parent   *File
}

func (f *File) addChild(dir bool, size int, name string) {
	if dir {
		f.Children = append(f.Children, &File{
			Type:     DIR,
			Name:     name,
			Size:     size,
			Children: []*File{},
			Parent:   f,
		})
	} else {
		f.Children = append(f.Children, &File{
			Type:     FILE,
			Name:     name,
			Size:     size,
			Children: nil,
			Parent:   f,
		})
	}
}

func main() {
	lines := AOCUTILS.ReadLines("input.txt")
	Root := File{
		Type:     DIR,
		Name:     "/",
		Size:     0,
		Children: []*File{},
		Parent:   nil,
	}
	Walker := &Root
	for i := 0; i < len(lines); {
		if lines[i][0:4] == "$ ls" {
			i += 1
			for i < len(lines) && lines[i][0] != '$' {
				parts := strings.Split(lines[i], " ")
				if parts[0] == "dir" {
					Walker.addChild(true, 0, parts[1])
				} else {
					Walker.addChild(false, AOCUTILS.ToIntMust(parts[0]), parts[1])
				}
				i += 1
			}
			continue
		} else if lines[i][0:6] == "$ cd /" {
			Walker = &Root
		} else {
			dir := strings.Split(lines[i], " ")[2]
			if dir == ".." {
				if Walker.Parent != nil {
					Walker = Walker.Parent
				}
			} else {
				for _, f := range Walker.Children {
					if f.Name == dir {
						Walker = f
					}
				}
			}
		}
		i += 1
	}
	Walker = &Root
	calculateSizes(Walker)
	sizes := AOCUTILS.SortAscending(getSizes(Walker))
	sum := 0
	for _, s := range sizes {
		if s <= 100000 {
			sum += s
		}
	}
	fmt.Println(sum)
	needed := 30000000 - (70000000 - Root.Size)
	for _, s := range sizes {
		if s >= needed {
			fmt.Println(s)
			break
		}
	}

}

func calculateSizes(f *File) int {
	if f.Type == FILE {
		return f.Size
	}
	for _, ch := range f.Children {
		f.Size += calculateSizes(ch)
	}
	return f.Size
}
