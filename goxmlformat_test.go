package goxmlformat

import (
	"testing"
)

// TODO: This is a BRITTLE test, break it apart to test formatting logic in smaller chunks
func TestFormatXML(t *testing.T) {
	input := `<?xml version="1.0" encoding="UTF-8"?><root><!-- comment --><hi>blah</hi><!-- comment --><a><![CDATA[<>>>>
<><>]]></a><attr hi="world">data</attr><list><selfend /><T>true</T><F>false</F><selfend /></list><next><inest><hi></hi></inest></nest></root>`

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<root>
   <!-- comment -->
   <hi>blah</hi>
   <!-- comment -->
   <a><![CDATA[<>>>>
<><>]]></a>
   <attr hi="world">data</attr>
   <list>
      <selfend />
      <T>true</T>
      <F>false</F>
      <selfend />
   </list>
   <next>
      <inest>
         <hi></hi>
      </inest>
   </nest>
</root>`

	actual := FormatXML(input)

	if actual != expected {
		t.Errorf("Actual [%s] does not match expected [%s]", actual, expected)
	}
}
