package goxmlformat

import (
	"fmt"
	"strings"
)

func identStr(indent int) string {
	return strings.Repeat("   ", indent)
}

// FormatXML - Returns xmlStr formatted for Pretty Printing
func FormatXML(xmlStr string) string {
	var final strings.Builder
	var buffer strings.Builder

	prevFinished := false
	var i int

	for _, c := range xmlStr {
		buffer.WriteRune(c)
		if c == '>' {
			bufStr := buffer.String()
			n := strings.LastIndex(bufStr, "<") + 1

			if bufStr[n] == '/' {
				i--
				if prevFinished {
					fmt.Fprintf(&final, "%s%s\n", identStr(i), bufStr)
				} else {
					fmt.Fprintf(&final, "%s\n", bufStr)
				}
				prevFinished = true
			} else {
				bufStr = strings.TrimSpace(bufStr)
				if prevFinished {
					fmt.Fprintf(&final, "%s%s", identStr(i), bufStr)
				} else {
					fmt.Fprintf(&final, "\n%s%s", identStr(i), bufStr)
				}
				prevFinished = false
				i++
			}

			buffer.Reset()
		}
	}

	return strings.TrimSpace(final.String())
}
