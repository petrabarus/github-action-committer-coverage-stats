package coverage

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
)

func LoadCoverageFromCoberturaXML(path string) (*Coverage, error) {
	var coverage Coverage

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		err = fmt.Errorf("failed to read file %s: %w", path, err)
		return nil, err
	}

	root := doc.Root()
	// query /packages/package/classes/class
	for _, classTag := range root.FindElements("//class") {
		fileCoverage, fileErr := loadFileCoverage(classTag)
		if fileErr != nil {
			continue
		}
		coverage.FileCoverage = append(coverage.FileCoverage, fileCoverage)
	}

	return &coverage, nil
}

func loadFileCoverage(root *etree.Element) (FileCoverage, error) {
	var fileCoverage FileCoverage
	fileName := root.SelectAttrValue("filename", "")
	if fileName == "" {
		return fileCoverage, fmt.Errorf("file name not found for class %s", root.SelectAttrValue("name", ""))
	}
	fmt.Println(fileName)
	fileCoverage.FileName = fileName

	lineCoverage, err := loadLineCoverage(root)
	if err != nil {
		return fileCoverage, err
	}
	fileCoverage.LineCoverage = lineCoverage

	return fileCoverage, nil
}

func loadLineCoverage(root *etree.Element) ([]LineCoverage, error) {
	var lineCoverages []LineCoverage

	linesTag := root.SelectElement("lines")
	for _, lineTag := range linesTag.SelectElements("line") {
		number := lineTag.SelectAttrValue("number", "")
		// convert number to int
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		lineCoverages = append(lineCoverages, LineCoverage{
			LineNumber: numberInt,
			IsCovered:  lineTag.SelectAttrValue("hits", "") != "0",
		})
	}

	return lineCoverages, nil
}
