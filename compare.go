package main

func Missing(source Manifest, backup Manifest) []string {

	exists := make(map[string]bool)

	for _, file := range backup.Files {

		exists[file] = true

	}

	var result []string

	for _, file := range source.Files {

		if !exists[file] {

			result = append(result, file)

		}

	}

	return result

}
