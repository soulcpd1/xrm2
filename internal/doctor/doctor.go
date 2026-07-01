package doctor

import (
	"fmt"

	"github.com/fatih/color"
)

type Result struct {
	Name   string
	OK     bool
	Detail string
}

type Checker interface {
	Run() Result
}

func Print(results []Result) {

	fmt.Println("========================================")
	fmt.Println("            XRM2 Doctor")
	fmt.Println("========================================")
	fmt.Println()

	for _, r := range results {

		if r.OK {
			color.Green("[ OK ] %-18s %s", r.Name, r.Detail)
		} else {
			color.Red("[FAIL] %-18s %s", r.Name, r.Detail)
		}
	}
}
