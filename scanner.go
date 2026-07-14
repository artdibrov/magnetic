package main

func Scan(name string, files []string) Manifest {

	return Manifest{
		Name: name,
		Files: files,
	}

}
