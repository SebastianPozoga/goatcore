package orm

import (
	"database/sql"
	"fmt"

	"github.com/goatcms/goat-core/db"
	"github.com/jmoiron/sqlx"
)

// BaseDAO is default dao interface
type BaseDAO struct {
	Table *BaseTable
}

// NewBaseDAO create new base DAO
func NewBaseDAO(bt *BaseTable) *BaseDAO {
	return &BaseDAO{
		Table: bt,
	}
}

// FindAll obtain all articles from database
func (dao *BaseDAO) FindAll(tx db.TX) (*sqlx.Rows, error) {
	return tx.Queryx(dao.Table.selectSQL)
}

// FindByID obtain article of given ID from database
func (dao *BaseDAO) FindByID(tx db.TX, id int64) *sqlx.Row {
	return tx.QueryRowx(dao.Table.selectByIDSQL, id)
}

// Insert store given articles to database
func (dao *BaseDAO) Insert(tx db.TX, entity interface{}) (int64, error) {
	var (
		res sql.Result
		err error
	)
	if res, err = tx.NamedExec(dao.Table.insertSQL, entity); err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertWithID store given articles to database (It persist with id from entity)
func (dao *BaseDAO) InsertWithID(tx db.TX, entity interface{}) error {
	if _, err := tx.NamedExec(dao.Table.insertWithIDSQL, entity); err != nil {
		return err
	}
	return nil
}

// Update data of article
func (dao *BaseDAO) Update(tx db.TX, entity interface{}) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = tx.NamedExec(dao.Table.updateByIDSQL, entity); err != nil {
		return err
	}
	if count, err = res.RowsAffected(); err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("Update modified more then one record (%v records modyfieds)", count)
	}
	return nil
}

// Delete remove specyfic record
func (dao *BaseDAO) Delete(tx db.TX, id int64) error {
	var (
		res   sql.Result
		err   error
		count int64
	)
	if res, err = tx.NamedExec(dao.Table.deleteByIDSQL, &IDContainer{id}); err != nil {
		return err
	}
	if count, err = res.RowsAffected(); err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("Delete more than one record (%v records deleted)", count)
	}
	return nil
}

// CreateTable add new table to a database
func (dao *BaseDAO) CreateTable(tx db.TX) error {
	tx.MustExec(dao.Table.createSQL)
	return nil
}
