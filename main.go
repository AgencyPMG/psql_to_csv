package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"psql_to_csv/config"
	"strings"
	"unicode/utf8"
)

type columnSizes []int

func main() {
	config.Load()
	inScan := bufio.NewScanner(os.Stdin)
	outCsv := csv.NewWriter(os.Stdout)
	outCsv.Comma = config.Comma()
	cs := columnSizes(nil)
	if inScan.Scan() {
		headerLine := inScan.Text()
		cs = columnSizesFromHeaderLine(headerLine)
		if config.HasHeaders() {
			headers := config.Headers()
			if len(headers) != len(cs) {
				fmt.Fprintln(os.Stderr, "supplied headers and columns in table are not equal")
				os.Exit(5)
			}
			outCsv.Write(headers)
		} else {
			outCsv.Write(convertRowLineToCsvRow(cs, headerLine))
		}
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
	for _, v := range cs {
		temp := getSubstringRuneCount(start, v, line)
		column := strings.TrimSpace(temp)
		result = append(result, column)
		start += len(temp) + 1
	}
	return result
}

func getSubstringRuneCount(index, count int, line string) string {
	line = line[index:]
	result := []rune{}
	for i := 0; i < count && len(line) > 0; i++ {
		r, size := utf8.DecodeRuneInString(line)
		result = append(result, r)
		line = line[size:]
	}
	return string(result)
}
