package dsql

import (
	"fmt"

	"github.com/goatcms/goat-core/db"
)

type DSQL struct{}

func NewDSQL() db.DSQL {
	return db.DSQL(DSQL{})
}

// NewSelectSQL create new select query
func (dsql DSQL) NewSelectSQL(table string, fields []string) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM " + table, nil
}

// NewSelectWhereSQL create new select query
func (dsql DSQL) NewSelectWhereSQL(table string, fields []string, where string) (string, error) {
	selectSQL, err := dsql.NewSelectSQL(table, fields)
	if err != nil {
		return "", err
	}
	return selectSQL + " WHERE " + where, nil
}

// NewInsertSQL create new insert sql
func (dsql DSQL) NewInsertSQL(table string, fields []string) (string, error) {
	sqlUpdate := "INSERT INTO " + table + " ("
	sqlValues := "VALUES ("
	for i, row := range fields {
		if i == 0 {
			sqlUpdate += "" + row
			sqlValues += ":" + row
		} else {
			sqlUpdate += ", " + row
			sqlValues += ", :" + row
		}
	}
	return sqlUpdate + ") " + sqlValues + ")", nil
}

// NewUpdateSQL create new update sql
func (dsql DSQL) NewUpdateSQL(table string, fields []string) (string, error) {
	sqlUpdate := "UPDATE " + table + " SET "
	for i, row := range fields {
		if i == 0 {
			sqlUpdate += row + " = :" + row
		} else {
			sqlUpdate += ", " + row + " = :" + row
		}
	}
	return sqlUpdate, nil
}

// NewUpdateWhereSQL create new select query
func (dsql DSQL) NewUpdateWhereSQL(table string, fields []string, where string) (string, error) {
	sql, err := dsql.NewUpdateSQL(table, fields)
	if err != nil {
		return "", err
	}
	return sql + " WHERE " + where, nil
}

// NewDeleteSQL create new delete query
func (dsql DSQL) NewDeleteSQL(table string) (string, error) {
	return "DELETE FROM " + table, nil
}

// NewDeleteWhereSQL create new delete query with where restrict
func (dsql DSQL) NewDeleteWhereSQL(table string, where string) (string, error) {
	deleteSQL, err := dsql.NewDeleteSQL(table)
	if err != nil {
		return "", err
	}
	return deleteSQL + " WHERE " + where, nil
}

// NewCreateSQL create new create table query
func (dsql DSQL) NewCreateSQL(table string, types map[string]string) (string, error) {
	sql := "CREATE TABLE " + table + " (\n"
	i := 0
	if len(types) == 0 {
		return "", fmt.Errorf("types can not be empty")
	}
	for name, typeDesc := range types {
		if i > 0 {
			sql += ",\n"
		}
		sql += name + " " + typeDesc
		i++
	}
	return sql + ")", nil
}
