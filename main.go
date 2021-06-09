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
