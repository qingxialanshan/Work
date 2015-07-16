package main

import (
	"fmt"
	"global"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Println("Please supply a page id!\nUsage: PageAction [pageName] [args...]")
		fmt.Println("      		   [setup]")
		fmt.Println("      		   [suggestion]")
		fmt.Println("      		   [install_config] [location] [execlocation]")
		fmt.Println("      		   [install_dependency]")
		fmt.Println("      		   [installation] [nextpage|install]")
		fmt.Println("      		   [device_setup] [nextpage|install]")
		fmt.Println("      		   [deploy_samples] [nextpage|install]")
		fmt.Println("      		   [finished]")
		fmt.Println("      		   [query] [compName]")
		os.Exit(-1)
	}
	if os.Args[1] == "query" {
		//query whethter the component selected
		compName := os.Args[2]
		if global.ComponentSelect(compName) {
			fmt.Println("0")
		} else {
			fmt.Println("-1")
		}
		os.Exit(0)
	}
	pkgname := os.Args[1]
	if comp, ok := comps[pkgname]; ok {
		comp.PostAction(os.Args[2:]...)
	} else {
		fmt.Println("the pageName is invalidate")
	}
}
