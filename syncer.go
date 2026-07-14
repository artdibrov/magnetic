package main

type SyncResult struct {
	Copy []string
	Ready []string
}

func BuildPlan(source Manifest, backup Manifest) SyncResult {

	var ready []string

	lookup := make(map[string]bool)

	for _, file := range backup.Files {

		lookup[file] = true

	}

	for _, file := range source.Files {

		if lookup[file] {

			ready = append(ready, file)

		}

	}

	return SyncResult{
		Copy: Missing(source, backup),
		Ready: ready,
	}

}
