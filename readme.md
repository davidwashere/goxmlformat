# Go XML Format
Simple package that formats a string of XML for pretty printing when the structure of the XML is not known ahead of time

At the time of writing this, the default `encoding/xml` package was not able to handle this [use case](https://github.com/golang/go/issues/26756)

## Example Usage

```go
package main

import (
    "fmt"

    "github.com/davidwashere/goxmlformat"
)

func main() {
    data := `<root><hello>world</hello><from name="David" /><meta><lang>go</lang><isfast>dunno</isfast></meta></root>`

    prettied := goxmlformat.FormatXML(data)

    fmt.Println(prettied)
}
```

Output:
```xml
<root>
   <hello>world</hello>
   <from name="David" />
   <meta>
      <lang>go</lang>
      <isfast>dunno</isfast>
   </meta>
</root>
```