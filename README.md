# nu
Go Library for scraping NU.nl RSS feed to JSON output


##Example

```go
package main

import (
	"fmt"
	"github.com/Rivalo/nu"
)

func main() {
  json, err := nu.Parse("http://www.nu.nl/rss/Algemeen")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}
  
  	fmt.Fprint(w, string(json))
}
```
