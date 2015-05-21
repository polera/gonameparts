# gonameparts
gonameparts splits a human name into individual parts. This is useful when dealing with external data sources that provide names as a single value, but you need to store the discrete parts in a database for example.

[![GoDoc](https://godoc.org/github.com/polera/gonameparts?status.svg)](https://godoc.org/github.com/polera/gonameparts)  [![Build Status](https://travis-ci.org/polera/gonameparts.svg)](https://travis-ci.org/polera/gonameparts)

Author
==
James Polera <james@uncryptic.com>

Dependencies
==
No external dependencies.  Uses Go's standard packages

Example
==

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/polera/gonameparts"
)

func main() {

	// Parsing a name and printing its parts
	nameString := gonameparts.Parse("Thurston Howell III")
	fmt.Println("FirstName:", nameString.FirstName)
	fmt.Println("LastName:", nameString.LastName)
	fmt.Println("Generation:", nameString.Generation)
	// Output:
	// FirstName: Thurston
	// LastName: Howell
	// Generation: III

    // Parse a name with multiple "also known as" aliases, output JSON
	multipleAKA := gonameparts.Parse("Tony Stark a/k/a Ironman a/k/a Stark, Anthony a/k/a Anthony Edward \"Tony\" Stark")
	jsonParts, _ := json.Marshal(multipleAKA)
	fmt.Printf("%v\n", string(jsonParts))
	/* Output:
		{
	    "aliases": [
	        {
	            "aliases": null,
	            "first_name": "Ironman",
	            "full_name": "Ironman",
	            "generation": "",
	            "last_name": "",
	            "middle_name": "",
	            "nickname": "",
	            "provided_name": " Ironman ",
	            "salutation": "",
	            "suffix": ""
	        },
	        {
	            "aliases": null,
	            "first_name": "Anthony",
	            "full_name": "Anthony Stark",
	            "generation": "",
	            "last_name": "Stark",
	            "middle_name": "",
	            "nickname": "",
	            "provided_name": " Stark, Anthony ",
	            "salutation": "",
	            "suffix": ""
	        },
	        {
	            "aliases": null,
	            "first_name": "Anthony",
	            "full_name": "Anthony Edward Stark",
	            "generation": "",
	            "last_name": "Stark",
	            "middle_name": "Edward",
	            "nickname": "\"Tony\"",
	            "provided_name": " Anthony Edward \"Tony\" Stark",
	            "salutation": "",
	            "suffix": ""
	        }
	    ],
	    "first_name": "Tony",
	    "full_name": "Tony Stark",
	    "generation": "",
	    "last_name": "Stark",
	    "middle_name": "",
	    "nickname": "",
	    "provided_name": "Tony Stark a/k/a Ironman a/k/a Stark, Anthony a/k/a Anthony Edward \"Tony\" Stark",
	    "salutation": "",
	    "suffix": ""
		}*/

}
```
