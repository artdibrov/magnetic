package main

import "fmt"

func Print(plan SyncResult, source string, backup string) {

	fmt.Println("Source :", source)

	fmt.Println()

	fmt.Println("Destination :", backup)

	fmt.Println()

	fmt.Println("Files to copy")

	fmt.Println()

	for _, file := range plan.Copy {

		fmt.Println("+", file)

	}

	fmt.Println()

	fmt.Println("Already synchronized")

	fmt.Println()

	for _, file := range plan.Ready {

		fmt.Println("✓", file)

	}

}
