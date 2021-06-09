package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// lists all the yaml or tmpl files in .deployments folder
func listYamlFiles(root string) []string {
	var files []string
	regss := regexp.MustCompile(".*\\.(yaml|yml|tmpl)$")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() && regss.MatchString(info.Name()) && strings.Contains(path, "deployments") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// List the docker images in Yaml file
func listDockerImages(files []string) {
	var dockerImages []string
	reg := regexp.MustCompile(`image:\s*(docker.appdirect.tools.*)\s*$`)

	for _, yamlFile := range files {
		file, err := os.Open(yamlFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			image := strings.TrimSpace(scanner.Text())
			if reg.MatchString(image) == true {

				if exists(dockerImages, image) == false {
					dockerImages = append(dockerImages, image)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	for _, dockerimage := range dockerImages {
		fmt.Println(stringsReplacer(dockerimage))
	}
}

// checks if the same image image exists in .deployments/kubernetes
func exists(dockerImages []string, docker_image string) bool {
	for _, i := range dockerImages {
		if i == docker_image {
			return true
		}
	}
	return false
}

// string replacer
func stringsReplacer(imagedetails string) string {
	replacer := strings.NewReplacer("image: ", "", "{{ .Values.imageVersion }}", os.Args[1])
	return replacer.Replace(imagedetails)
}

func main() {
	files := listYamlFiles(".")
	listDockerImages(files)
}
