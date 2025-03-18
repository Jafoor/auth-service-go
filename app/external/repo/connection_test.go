package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_Success(t *testing.T) {
	readDB, readMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer readDB.Close()

	writeDB, writeMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer writeDB.Close()

	// Expect Ping to return no error
	readMock.ExpectPing().WillReturnError(nil)
	writeMock.ExpectPing().WillReturnError(nil)

	// Wrap mock database in DB struct
	db := &DB{
		ReadDB:  sqlx.NewDb(readDB, "postgres"),
		WriteDB: sqlx.NewDb(writeDB, "postgres"),
	}

	err = db.HealthCheck()

	assert.NoError(t, err)
}
