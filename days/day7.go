package days

import (
	"bufio"
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed inputs/day7.txt
var input7 string

func Day7() string {
	input := input7

	fs := make(map[string][]File)

	fs["/"] = make([]File, 0)

	currentPath := "/"

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		if strings.HasPrefix(line, "$") {
			command := line[2:4]
			if command != "cd" {
				continue
			}

			path := strings.TrimSpace(line[5:])

			if path == ".." {
				currentPath = currentPath[:strings.LastIndex(currentPath, "/")]
			} else if path == "/" {
				currentPath = "/"
			} else {
				currentPath += "/" + path
			}
			continue
		}
		split := strings.Split(line, " ")
		sizeOrType := split[0]
		name := split[1]

		if sizeOrType == "dir" {
			fs[currentPath+"/"+name] = make([]File, 0)
		} else {
			files := fs[currentPath]
			size, _ := strconv.Atoi(sizeOrType)
			files = append(files, File{
				Name: name,
				Size: size,
			})
			fs[currentPath] = files
		}

	}
	overAll := 0
	for c := range fs {
		count := getSize(fs, c)
		if count <= 100000 {
			overAll += count
		}
	}
	return strconv.Itoa(overAll)
}

func getSize(list map[string][]File, path string) int {
	count := 0
	for _, a := range list[path] {
		count += a.Size
	}
	for b := range list {
		if strings.HasPrefix(b, path) && b != path && strings.Count(b, "/") == strings.Count(path, "/")+1 {
			count += getSize(list, b)
		}
	}
	return count
}

func Day7_2() string {
	input := input7

	fs := make(map[string][]File)

	fs["/"] = make([]File, 0)

	currentPath := "/"

	scannner := bufio.NewScanner(strings.NewReader(input))
	for scannner.Scan() {
		line := scannner.Text()
		if strings.HasPrefix(line, "$") {
			command := line[2:4]
			if command != "cd" {
				continue
			}

			path := strings.TrimSpace(line[5:])

			if path == ".." {
				currentPath = currentPath[:strings.LastIndex(currentPath, "/")]
			} else if path == "/" {
				currentPath = "/"
			} else {
				currentPath += "/" + path
			}
			continue
		}
		split := strings.Split(line, " ")
		sizeOrType := split[0]
		name := split[1]

		if sizeOrType == "dir" {
			fs[currentPath+"/"+name] = make([]File, 0)
		} else {
			files := fs[currentPath]
			size, _ := strconv.Atoi(sizeOrType)
			files = append(files, File{
				Name: name,
				Size: size,
			})
			fs[currentPath] = files
		}

	}
	overAll := getSize(fs, "/")
	free := 70000000 - overAll
	required := 30000000 - free
	smallest := 70000000
	for c := range fs {
		size := getSize(fs, c)
		if size >= required {
			smallest = int(math.Min(float64(smallest), float64(size)))
		}
	}
	return strconv.Itoa(smallest)
}

type StringPair struct {
	Path string
	Size int
}

type File struct {
	Name string
	Size int
}
