package orm

import (
	"database/sql"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/varutil"
)

// InsertContext is context for findByID function
type InsertContext struct {
	query string
}

// Insert create new record
func (q InsertContext) Insert(tx db.TX, entity interface{}) (int64, error) {
	var (
		res sql.Result
		err error
		id  int64
	)
	if res, err = tx.NamedExec(q.query, entity); err != nil {
		return -1, err
	}
	if id, err = res.LastInsertId(); err != nil {
		return -1, err
	}
	if err = varutil.SetField(entity, "ID", id); err != nil {
		return -1, err
	}
	return id, nil
}

// InsertContext create new dao function instance
func NewInsert(table db.Table, dsql db.DSQL) (db.Insert, error) {
	query, err := dsql.NewInsertSQL(table.Name(), table.Fields())
	if err != nil {
		return nil, err
	}
	context := &InsertContext{
		query: query,
	}
	return context.Insert, nil
}
