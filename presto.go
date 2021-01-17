/*
Package presto-trino provides a standard database/sql driver for
Facebook's Presto and Trino query engine.
*/
package presto

const (
	version = "0.1.5"

	// Presto headers
	userHeaderPresto    = "X-Presto-User"
	sourceHeaderPresto  = "X-Presto-Source"
	catalogHeaderPresto = "X-Presto-Catalog"
	schemaHeaderPresto  = "X-Presto-Schema"

	// Trino headers
	userHeaderTrino    = "X-Trino-User"
	sourceHeaderTrino  = "X-Trino-Source"
	catalogHeaderTrino = "X-Trino-Catalog"
	schemaHeaderTrino  = "X-Trino-Schema"

	// General headers
	userAgentHeader = "User-Agent"
	userAgent       = "go-presto-trino/" + version
)
