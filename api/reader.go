package api

import "strings"

func GetStringReader(s string) *strings.Reader{
	return strings.NewReader(s)
}
