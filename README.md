# go-icndb

[![Build Status](https://travis-ci.org/saadullahsaeed/go-icndb.svg?branch=master)](https://travis-ci.org/saadullahsaeed/go-icndb)

GoLang client for the Internet Chuck Norris Database API (http://www.icndb.com/api/).

## Usage

#### Get a Random Joke

```go
package main

import (
	"fmt"
	"log"

	"github.com/saadullahsaeed/go-icndb"
)

func main() {
	joke, err := chucknorris.GetRandomJoke(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nID: %d ", joke.ID)
	fmt.Printf("\nText: %s", joke)
}
```

#### Specify Custom Name for the Joke

```go
  req := &chucknorris.CNDBRequest{
    FirstName: "John",
    LastName: "Doe",
  }
  joke, err := chucknorris.GetRandomJoke(req)
  if err != nil {
  	log.Fatal(err)
  }
  fmt.Printf("\nJoke: %s", joke)
```
