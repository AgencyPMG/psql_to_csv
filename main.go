package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type columnSizes []int

func main() {
	inScan := bufio.NewScanner(os.Stdin)
	outCsv := csv.NewWriter(os.Stdout)
	cs := columnSizes(nil)
	if inScan.Scan() {
		headerLine := inScan.Text()
		cs = columnSizesFromHeaderLine(headerLine)
		outCsv.Write(convertRowLineToCsvRow(cs, headerLine))
		if outCsv.Error() != nil {
			fmt.Fprintln(os.Stderr, "could not write out header row")
			os.Exit(4)
		}
	}
	if len(cs) == 0 {
		fmt.Fprintln(os.Stderr, "could not read column sizes")
		os.Exit(1)
	} else {
	}
	if !inScan.Scan() {
		fmt.Fprintln(os.Stderr, "could not read second row of input")
		os.Exit(2)
	}
	for inScan.Scan() {
		line := inScan.Text()
		if line[0] == '(' {
			break
		}
		rowResult := convertRowLineToCsvRow(cs, line)
		outCsv.Write(rowResult)
		if outCsv.Error() != nil {
			break
		}
	}
	for inScan.Scan() {
	}
	outCsv.Flush()
	if outCsv.Error() != nil {
		fmt.Fprintln(os.Stderr, "could not write out result row")
		os.Exit(3)
	}
}

func columnSizesFromHeaderLine(line string) columnSizes {
	values := strings.Split(line, "|")
	result := []int{}
	for _, value := range values {
		result = append(result, len(value))
	}
	return result
}

func convertRowLineToCsvRow(cs columnSizes, line string) []string {
	result := []string{}
	start := 0
	for i, v := range cs {
		end := start + v
		if i == len(cs)-1 {
			end = len(line)
		}
		column := strings.TrimSpace(line[start:end])
		result = append(result, column)
		start += v + 1
	}
	return result
}
