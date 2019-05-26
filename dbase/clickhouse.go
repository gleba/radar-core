package dbase

import (
	"database/sql"
	"fmt"
	"github.com/gleba/radar-core/ux"
	"github.com/jmoiron/sqlx"
	"github.com/kshvakov/clickhouse"
)


type ClickHouseInstance struct{
	X *sqlx.DB
}

func ConnectClickHouse(path string) *ClickHouseInstance {
	var err error
	chi := ClickHouseInstance{}
	chi.X, err = sqlx.Open("clickhouse", path)
	ux.Safe(err)
	if err := chi.X.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
	}
	for _, table := range tablies {
		_, err = chi.X.Exec(table)
		ux.Safe(err)
	}
	return &chi
}

type ClickHouseWriter struct {
	X     *sql.Tx
	stmt  *sql.Stmt
	xx    *sql.DB
	Count int
}

func (chi *ClickHouseInstance) Writer() *ClickHouseWriter  {
	writer := ClickHouseWriter{
		Count: 0,
	}
	var err error
	writer.X, err = chi.X.Begin()
	ux.Safe(err)
	return &writer
}

func (w *ClickHouseWriter) Statement(stmt *sql.Stmt, err error) {
	if ux.Safe(err) {
		w.stmt = stmt
	}
}

func (w *ClickHouseWriter) Add(data ...interface{}) {
	_, err := w.stmt.Exec(data...)
	ux.Safe(err)
	w.Count = 1 + w.Count
}
func (w *ClickHouseWriter) Commit() {
	if w.Count >= 1 {
		ux.Safe(w.X.Commit())
		w.Count = 0
	}
}