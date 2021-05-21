package main

import (
	"fmt"

	"mdim/core"
)

func main() {
	// Deal with options.
	cliOptions := core.GetOptions()

	fmt.Println("Starting to scan markdown document..")

	// Scan docs in docFolder to fix image relative path.
	allRefImgsMap, aggErr := core.ScanToFixImgRelPath(cliOptions.DocFolder, cliOptions.ImgFolder, cliOptions.DoRelPathFix)
	if aggErr != nil {
		aggErr.PrintAggregateError()
		return
	}

	fmt.Println("Starting to scan images..")

	// Delete no reference images.
	if errs := core.DelNoRefImgs(cliOptions.ImgFolder, allRefImgsMap, cliOptions.DoImgDel); errs != nil {
		errs.PrintAggregateError()
		return
	}

	fmt.Println("All done.")
}
