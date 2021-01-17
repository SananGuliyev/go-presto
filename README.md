Presto & Trino for Go
=============

This is a tiny golang client for Facebook's [Presto][1] & [Trino][2] query engine.

```go
import (
  "fmt"
  "github.com/SananGuliyev/go-presto"
)

// Host, user, source, catalog, schema, query
sql := "SELECT * FROM sys.node"
query, _ := presto.NewQuery("http://presto-or-trino-coordinator:8080", "", "", "", "", sql)

if row, _ := query.Next(); row != nil {
  fmt.Println(row...)
}
```

[1]: https://github.com/facebook/presto
[2]: https://trino.io/
