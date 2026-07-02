package check

import "fmt"

func Print(title string, results []Result) {

	fmt.Println("========================================")
	fmt.Printf("        %s\n", title)
	fmt.Println("========================================")
	fmt.Println()

	for _, r := range results {

		status := "FAIL"
		if r.OK {
			status = " OK "
		}

		fmt.Printf("[%s] %-20s %s\n", status, r.Name, r.Detail)
	}
}
