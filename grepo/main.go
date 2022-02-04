package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func includeDotFiles() {
	//TODO
}

func maxDepth() {
	//TODO
}

func regex() {
	//TODO
}

func exclude() {
	//TODO

}

func contains(s string, slice []string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, term := range slice {
		set[term] = struct{}{}
	}
	_, ok := set[s]
	return ok
}

func traverse(directory string) []string {
	files := []string{}
	items, err := ioutil.ReadDir(directory)
	//TODO: Implement ExcludeItems
	excludedItems := []string{""}
	if err != nil {
		fmt.Printf("Error fetching files or directories inside %s : \n %s", items, err)
	}
	for _, item := range items {
		if contains(item.Name(), excludedItems) {
			continue
		}
		if item.IsDir() {
			nextStep := directory + "/" + item.Name()
			subFiles := traverse(nextStep)
			files = append(files, subFiles...)
		} else {
			file := directory + "/" + item.Name()
			files = append(files, file)
		}
	}
	return files
}

func findString(file string, searchString string) ([]string, error) {
	re, _ := regexp.Compile(searchString)
	output := []string{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if re.MatchString(line) {
			output = append(output, file+": "+line)
		} else {
			continue
		}
	}

	return output, nil
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	var r bool
	flag.BoolVar(&r, "r", false, "Enable recursively search.")
	//excludedItems := flag.String("exclude", "", "Items to be excluded from the search.")
	flag.Parse()
	values := flag.Args()
	if len(values) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	dirLocation := values[len(values)-1]
	if dirLocation == "" {
		dirLocation = "."
	}

	searchString := values[len(values)-2]
	if searchString == "" {
		fmt.Println("No string defined for search")
	}

	if isFlagPassed("excludedItems") {
		exclude()
	}
	//filesToSearch := []string{}
	files := traverse(dirLocation)
	for _, file := range files {
		output, err := findString(file, searchString)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(output)
		}
	}
}
