package main

func main() {

	source := Scan(
		"Documents",
		SourceFiles,
	)

	backup := Scan(
		"Backup",
		BackupFiles,
	)

	plan := BuildPlan(
		source,
		backup,
	)

	Print(
		plan,
		source.Name,
		backup.Name,
	)

}
