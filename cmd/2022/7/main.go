package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

var sizes map[string]int

type Node interface {
	Name() string
}

type Dir struct {
	parent *Dir
	name   string
	files  map[string]Node
}

type File struct {
	name string
	size int
}

func (d *Dir) Name() string {
	return d.name
}

func (f File) Name() string {
	return f.name
}

func (d *Dir) AddNode(name string, n Node) {
	d.files[name] = n
}

func (d Dir) Size() int {
	total := 0
	for k, v := range d.files {
		switch nt := v.(type) {
		case File:
			total += nt.size
		case *Dir:
			s := nt.Size()
			sizes[k] = s
			total += s
		default:
			log.Fatalf("unknown type: %v", nt)
		}
	}
	return total
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	_, filename, _, _ := runtime.Caller(0)
	input := aoc.ReadFile(f, filename)

	var cmds []string
	for i, v := range strings.Split(input, "\n$") {
		if v == "" || i == 0 {
			continue
		}
		cmds = append(cmds, v)
	}
	root := Dir{nil, "/", make(map[string]Node, 0)}
	var curr *Dir = &root
	for _, v := range cmds {
		arr := strings.Split(v, "\n")
		cmd := strings.TrimSpace(arr[0])
		if cmd == "ls" {
			for _, f := range arr[1:] {
				if f == "" {
					continue
				}
				dirent := strings.Split(f, " ")
				var node Node
				filename := curr.name + dirent[1]
				if dirent[0] == "dir" {
					filename = filename + "/"
					node = &Dir{curr, filename, make(map[string]Node, 0)}
				} else {
					s := aoc.ParseInt(dirent[0])
					node = File{filename, s}
				}
				curr.AddNode(filename, node)
			}
		} else {
			arr := strings.Split(cmd, " ")
			if arr[1] == ".." {
				curr = curr.parent
			} else {
				curr = curr.files[curr.name+arr[1]+"/"].(*Dir)
			}
		}
	}

	totalSize := 70000000
	need := 30000000

	sizes = map[string]int{}
	used := root.Size()
	unused := totalSize - used
	atmost := 0
	minSize := used
	for _, v := range sizes {
		if v <= 100000 {
			atmost += v
		}
		if unused+v > need && v < minSize {
			minSize = v
		}
	}
	fmt.Printf("atmost: %d\n", atmost)
	fmt.Printf("min: %d\n", minSize)
}
