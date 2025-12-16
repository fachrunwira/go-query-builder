# Go Query Builder

Go Query Builder is not an ORM, it just simplified your queries.
This query builder is still supported for MySQL only.
It will also support for Postgre in future updates.

## Install
Go 1.24 or higher

```bash
go get github.com/fachrunwira/go-query-builder
```

## Usage
Basic example of query builder 
```go
package main

import (
	"fmt"
	"log"

	"github.com/fachrunwira/go-query-builder/builder"
	"github.com/fachrunwira/go-query-builder/examples"
)

func main() {
  // Initialzie database
	db, err := examples.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

  // SELECT * FROM users;
	result, err := builder.Make(db).Table("users").Get()
	if err != nil {
		log.Fatalln(err)
	}

  // results is in builder.Rows type or []map[string]any
	for _, m := range result {
		for col, val := range m {
			fmt.Printf("%s: %v\n", col, val)
		}
	}
}
```

### Fetching Single Row
```go
  // SELECT id, name FROM users LIMIT 1;
  result, err := builder.Make(db).
		Table("users").
		Select("id", "name").
		First()

  if err != nil {
		log.Fatalln(err)
	}

	// the query result is in builder.Row or map[string]any
  for col, val := range result {
		fmt.Printf("%s: %v\n", col, val)
	}
```

### Affecting Rows
When affecting rows, it already wrapped on transaction

#### Insert
```go
  // using map[string]any
	usersData := map[string]any{
	 	"id": 1,
	 	"name": "Example Name",
	}

  // or using builder.Row
	usersData := builder.Row{
		"id": 1,
		"name": "Example Name",
	}

  // INSERT INTO users (id, name) VALUES (?, ?);
  err = builder.Make(db).Table("users").Insert(usersData).Save()
```

#### Insert Multiple Rows
```go
  // using []map[string]any
	usersData := []map[string]any{
		{
			"id": 1,
			"name": "Example Name 1",
		},
		{
			"id": 2,
			"name": "Example Name 2",
		},
	}

	// or using builder.Rows
	usersData := builder.Rows{
		{
			"id":   1,
			"name": "Example Name 1",
		},
		{
			"id":   2,
			"name": "Example Name 2",
		},
	}

	// INSERT INTO users (id, name) VALUES (?, ?), (?, ?);
	err = builder.Make(db).Table("users").Insert(usersData).Save()
```

#### Insert Using Raw Method
```go
  err = builder.Make(db).InsertRaw("INSERT INTO users (id, name) VALUES (?, ?);", 1, "Example Name")
```

#### Update Row
```go
	// using map[string]any
	userData := map[string]any{
		"name": "Update Name",
	}

	// or using builder.Row
	userData := builder.Row{
		"name": "Update Name",
	}

	// UPDATE users SET name = ? WHERE id = ?;
	err = builder.Make(db).Table("users").Where("id", 1).Update(userData).Save()
```

#### Update Row Using Raw Method
```go
  err = builder.Make(db).UpdateRaw("UPDATE users SET name = ? WHERE id = ?;", "New Name", 1)
```

#### Delete Row
```go
	// DELETE FROM users WHERE id = ?;
	err = builder.Make(db).Table("users").Where("id", 1).Delete().Save()
```

#### Delete Row Using Raw Method
```go
  err = builder.Make(db).DeleteRaw("DELETE FROM users WHERE id = ?;", 1)
```