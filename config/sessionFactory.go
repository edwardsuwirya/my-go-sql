package config

import "database/sql"

type SessionFactory struct {
	Db *sql.DB
}

type Session struct {
	Db *sql.DB
	Tx *sql.Tx
}

func NewSessionFactory(driverName, dataSourceName string) (*SessionFactory, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	factory := new(SessionFactory)
	factory.Db = db
	return factory, nil
}

func (sf *SessionFactory) GetSession() *Session {
	session := new(Session)
	session.Db = sf.Db
	return session
}

func (s *Session) Begin() error {
	if s.Tx == nil {
		tx, err := s.Db.Begin()
		if err != nil {
			return err
		}
		s.Tx = tx
		return nil
	}
	return nil
}

func (s *Session) Exec(query string, args ...interface{}) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.Exec(query, args...)
	}
	return s.Db.Exec(query, args...)
}
func (s *Session) ExecStatement(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.Stmt(stmt).Exec(args...)
	}
	return stmt.Exec(args)
}

func (s *Session) Prepare(query string) (*sql.Stmt, error) {
	if s.Tx != nil {
		return s.Tx.Prepare(query)
	}
	return s.Db.Prepare(query)
}

func (s *Session) Rollback() error {
	if s.Tx != nil {
		err := s.Tx.Rollback()
		if err != nil {
			return err
		}
		s.Tx = nil
		return nil
	}
	return nil
}

func (s *Session) Commit() error {
	if s.Tx != nil {
		err := s.Tx.Commit()
		if err != nil {
			return err
		}
		s.Tx = nil
		return nil
	}
	return nil
}

func (s *Session) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if s.Tx != nil {
		return s.Tx.Query(query, args...)
	}
	return s.Db.Query(query, args...)
}
func (s *Session) QueryStatement(stmt *sql.Stmt, args ...interface{}) (*sql.Rows, error) {
	if s.Tx != nil {
		return s.Tx.Stmt(stmt).Query(args...)
	}
	return stmt.Query(args...)
}

func (s *Session) QueryRow(query string, args ...interface{}) *sql.Row {
	if s.Tx != nil {
		return s.Tx.QueryRow(query, args...)
	}
	return s.Db.QueryRow(query, args...)
}
func (s *Session) QueryRowStatement(stmt *sql.Stmt, args ...interface{}) *sql.Row {
	if s.Tx != nil {
		return s.Tx.Stmt(stmt).QueryRow(args...)
	}
	return stmt.QueryRow(args...)
}
