package goxmlformat

import (
	"fmt"
	"strings"
)

func identStr(indent int) string {
	return strings.Repeat("   ", indent)
}

// This is a test string the validates all the desired functionality
// TODO: ADD TO TEST
/*
var data = `<?xml version="1.0" encoding="UTF-8"?><root><!-- comment --><hi>blah</hi><!-- comment --><a><![CDATA[<>>>>
<><>]]></a><attr hi="world">data</attr><list><selfend /><T>true</T><F>false</F><selfend /></list></root>`
*/

// FormatXML - Returns xmlStr formatted for Pretty Printing
func FormatXML(xmlStr string) string {
	var final strings.Builder
	var buffer strings.Builder
	var rolling string

	prevFinished := false
	hitNiner := false
	inCDATA := false

	var i int // indent level

	for _, c := range xmlStr {
		buffer.WriteRune(c)

		if hitNiner {
			rolling = rolling[1:] + string(c)
		} else {
			rolling += string(c)
			if len(rolling) == 9 {
				hitNiner = true
			}
		}

		if inCDATA {
			if strings.HasSuffix(rolling, "]]>") {
				inCDATA = false
			}
			continue
		}

		if rolling == "<![CDATA[" {
			inCDATA = true
			continue
		}

		if c == '>' {
			bufStr := buffer.String()
			first := strings.LastIndex(bufStr, "<") + 1
			last := len(bufStr) - 2

			// fmt.Printf("First[%d] Last[%d]\n", first, last)

			if bufStr[first] == '/' {
				i--
				if prevFinished {
					bufStr = strings.TrimSpace(bufStr)
					fmt.Fprintf(&final, "%s%s\n", identStr(i), bufStr)
				} else {
					fmt.Fprintf(&final, "%s\n", bufStr)
				}
				prevFinished = true
			} else if bufStr[first] == '?' || bufStr[last] == '/' || bufStr[first] == '!' {
				if prevFinished {
					bufStr = strings.TrimSpace(bufStr)
					fmt.Fprintf(&final, "%s%s\n", identStr(i), bufStr)
				} else {
					bufStr = strings.TrimSpace(bufStr)
					fmt.Fprintf(&final, "\n%s%s\n", identStr(i), bufStr)
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
