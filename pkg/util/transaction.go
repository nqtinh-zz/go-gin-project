package util

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// ContextKey type
type ContextKey string

// ContextKeyTx key
const ContextKeyTx ContextKey = "tx"

// TxLogger is sqlx.Tx wrapper with injected logger
type TxLogger struct {
	Tx     *sqlx.Tx
	Logger *logrus.Logger
}

/* Implementation of sqlx.Tx wrap logging */
// Get within a transaction.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (tx *TxLogger) Get(dest interface{}, query string, args ...interface{}) error {
	tx.Logger.Print(tx.bindArgs(query, args...))
	return tx.Tx.Get(dest, query, args...)
}

// Select within a transaction.
// Any placeholder parameters are replaced with supplied args.
func (tx *TxLogger) Select(dest interface{}, query string, args ...interface{}) error {
	tx.Logger.Print(tx.bindArgs(query, args...))
	return tx.Tx.Select(dest, query, args...)
}

// Exec executes a query that doesn't return rows.
// For example: an INSERT and UPDATE.
func (tx *TxLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	tx.Logger.Print(tx.bindArgs(query, args...))
	return tx.Tx.Exec(query, args...)
}

// QueryRowx executes a query that return rows.
// For example: an INSERT and UPDATE.
func (tx *TxLogger) QueryRowx(query string, dest interface{}, args ...interface{}) error {
	tx.Logger.Print(tx.bindArgs(query, args...))
	return tx.Tx.QueryRowx(query, args...).Scan(dest)
}

// QueryRowxStruct executes a query that return rows.
func (tx *TxLogger) QueryRowxStruct(query string, dest interface{}, args ...interface{}) error {
	tx.Logger.Print(tx.bindArgs(query, args...))
	return tx.Tx.QueryRowx(query, args...).StructScan(dest)
}

func (tx *TxLogger) bindArgs(query string, args ...interface{}) string {
	bindedQuery := query
	for i := range args {
		bindedQuery = strings.Replace(bindedQuery, fmt.Sprintf("$%d", i+1), fmt.Sprintf(`'%v'`, args[i]), 1)
	}

	return bindedQuery
}
