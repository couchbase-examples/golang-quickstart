package services

import (
	"github.com/couchbase/gocb/v2"
)

// rowIterator is the part of *gocb.QueryResult that collectRows depends on.
// It exists so the row-collection logic can be tested without a cluster.
type rowIterator interface {
	Next() bool
	Row(valuePtr interface{}) error
	Err() error
	Close() error
}

// queryAll runs a SQL++ query against the scope with named parameters and
// decodes every row into a fresh T.
func queryAll[T any](scope *gocb.Scope, statement string, params map[string]interface{}) ([]T, error) {
	result, err := scope.Query(statement, &gocb.QueryOptions{NamedParameters: params})
	if err != nil {
		return nil, err
	}
	return collectRows[T](result)
}

// collectRows drains rows into a slice. Each row is decoded into a new T so
// fields missing from one row never inherit values from the previous row.
// Errors reported after streaming (rows.Err) are returned rather than
// silently truncating the result. An empty result is an empty, non-nil slice.
func collectRows[T any](rows rowIterator) ([]T, error) {
	defer rows.Close()
	documents := []T{}
	for rows.Next() {
		var document T
		if err := rows.Row(&document); err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return documents, nil
}
