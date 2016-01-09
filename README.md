# nu
Go Library for scraping NU.nl RSS feed to JSON output

> XML is crap. Really. There are no excuses. XML is nasty to parse for humans, and it's a disaster to parse even for computers. There's just no reason for that horrible crap to exist. - Linus Torvalds


##Example

```go
package main

import (
	"fmt"
	"github.com/Rivalo/nu"
)

func main() {
	json, err := nu.JSON("http://www.nu.nl/rss/Algemeen")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}
  
	fmt.Println(string(json))
}
```
