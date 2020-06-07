package db

import "database/sql"

// DataStorage is inteface of db
type DataStorage interface {
	GetNews() ([]*Article, error)
}

// NewSQLDataStorage create new SQLDataStorage instance
func NewSQLDataStorage(sqlDB *sql.DB) SQLDataStorage {
	return SQLDataStorage{DB: sqlDB}
}

// SQLDataStorage is sql base data storage
type SQLDataStorage struct {
	DB *sql.DB
}
