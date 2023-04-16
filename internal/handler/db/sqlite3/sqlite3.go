package sqlite3

import (
	"database/sql"
	"time"
)

const (
	taqqueFolderName = ".taqque"
	timeFormat       = "2006-01-02 15:04:05.999999 -0700 MST"
)

type sqlite3 struct {
	db *sql.DB
}

func newSqlite3(userHomeDir string, db *sql.DB) *sqlite3 {
	return &sqlite3{
		db: db,
	}
}

func (s *sqlite3) initDB(createTableSQL string) error {
	_, err := s.db.Exec(createTableSQL)
	return err
}

func getTaqqueFolderPath(userHomeDir string) string {
	return userHomeDir + "/" + taqqueFolderName
}

func getTaqqueDBFilePath(userHomeDir, tableName string) string {
	return getTaqqueFolderPath(userHomeDir) + "/" + tableName + ".db"
}

func (s *sqlite3) getNow() string {
	return s.getSqlite3Time(time.Now()).Format(timeFormat)
}

func (s *sqlite3) getSqlite3Time(baseTime time.Time) time.Time {
	return baseTime.Truncate(time.Microsecond)
}

func (s *sqlite3) parseTime(timeStr string) (time.Time, error) {
	return time.Parse(timeFormat, timeStr)
}
