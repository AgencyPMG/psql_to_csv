package config

import (
	"flag"
	"fmt"
	"strings"
)

type config struct {
	headers []string
	comma   rune
}

var data *config

func init() {
	data = &config{
		nil,
		',',
	}
}

func Load() {
	loadArgs()
	fmt.Println(*data)
}

func loadArgs() {
	headers := flag.String("headers", "", "comma separated list of output headers to override table values")
	comma := flag.Int("comma", int(data.comma), "field delimiter in output")
	flag.Parse()
	if *headers != "" {
		data.headers = strings.Split(*headers, ",")
	}
	data.comma = rune(*comma)
}

func HasHeaders() bool {
	return data.headers != nil
}

func Headers() []string {
	return data.headers
}

func Comma() rune {
	return data.comma
}
