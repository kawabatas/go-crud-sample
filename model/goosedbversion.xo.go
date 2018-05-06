// Package model contains the types for schema 'go_crud_sample'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// GooseDbVersion represents a row from 'go_crud_sample.goose_db_version'.
type GooseDbVersion struct {
	ID        uint64         `json:"id"`         // id
	VersionID int64          `json:"version_id"` // version_id
	IsApplied bool           `json:"is_applied"` // is_applied
	Tstamp    mysql.NullTime `json:"tstamp"`     // tstamp

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GooseDbVersion exists in the database.
func (gdv *GooseDbVersion) Exists() bool {
	return gdv._exists
}

// Deleted provides information if the GooseDbVersion has been deleted from the database.
func (gdv *GooseDbVersion) Deleted() bool {
	return gdv._deleted
}

// Insert inserts the GooseDbVersion to the database.
func (gdv *GooseDbVersion) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if gdv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO go_crud_sample.goose_db_version (` +
		`version_id, is_applied, tstamp` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	res, err := db.Exec(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	gdv.ID = uint64(id)
	gdv._exists = true

	return nil
}

// Update updates the GooseDbVersion in the database.
func (gdv *GooseDbVersion) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !gdv._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if gdv._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE go_crud_sample.goose_db_version SET ` +
		`version_id = ?, is_applied = ?, tstamp = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp, gdv.ID)
	_, err = db.Exec(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp, gdv.ID)
	return err
}

// Save saves the GooseDbVersion to the database.
func (gdv *GooseDbVersion) Save(db XODB) error {
	if gdv.Exists() {
		return gdv.Update(db)
	}

	return gdv.Insert(db)
}

// Delete deletes the GooseDbVersion from the database.
func (gdv *GooseDbVersion) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !gdv._exists {
		return nil
	}

	// if deleted, bail
	if gdv._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM go_crud_sample.goose_db_version WHERE id = ?`

	// run query
	XOLog(sqlstr, gdv.ID)
	_, err = db.Exec(sqlstr, gdv.ID)
	if err != nil {
		return err
	}

	// set deleted
	gdv._deleted = true

	return nil
}

// GooseDbVersionByID retrieves a row from 'go_crud_sample.goose_db_version' as a GooseDbVersion.
//
// Generated from index 'goose_db_version_id_pkey'.
func GooseDbVersionByID(db XODB, id uint64) (*GooseDbVersion, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, version_id, is_applied, tstamp ` +
		`FROM go_crud_sample.goose_db_version ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	gdv := GooseDbVersion{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
	if err != nil {
		return nil, err
	}

	return &gdv, nil
}

// GooseDbVersionByID retrieves a row from 'go_crud_sample.goose_db_version' as a GooseDbVersion.
//
// Generated from index 'id'.
func GooseDbVersionByID(db XODB, id uint64) (*GooseDbVersion, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, version_id, is_applied, tstamp ` +
		`FROM go_crud_sample.goose_db_version ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	gdv := GooseDbVersion{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
	if err != nil {
		return nil, err
	}

	return &gdv, nil
}
