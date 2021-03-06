package session

import (
	dialect2 "awesomeProject4/geeorm/day2-reflect-schema/dialect"
	"awesomeProject4/geeorm/day2-reflect-schema/log"
	"awesomeProject4/geeorm/day2-reflect-schema/schema"
	clause2 "awesomeProject4/geeorm/day3-save-query/clause"
	"database/sql"
	"strings"
)

type Session struct {
	db *sql.DB
	dialect dialect2.Dialect
	refTable *schema.Schema
	clause clause2.Clause
	sql strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB,dialect dialect2.Dialect) *Session{
	return &Session{
		db:db,
		dialect:dialect,
	}
}

func(s *Session)Clear(){
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause2.Clause{}
}

func (s *Session) DB() *sql.DB{
	return s.db
}

func (s *Session) Raw(sql string,values ...interface{}) *Session{
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars,values...)
	return s
}

//Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result,err error){
	defer s.Clear()
	log.Info(s.sql.String(),s.sqlVars)
	if result,err = s.DB().Exec(s.sql.String(),s.sqlVars...);err != nil{
		log.Error(err)
	}
	return
}

//QueryRow gets a record from db
func(s *Session)QueryRow() *sql.Row{
	defer s.Clear()
	log.Info(s.sql.String(),s.sqlVars)
	return s.DB().QueryRow(s.sql.String(),s.sqlVars)
}

//QueryRows gets a list of records from db
func(s *Session) QueryRows()(rows *sql.Rows,err error){
	defer s.Clear()
	log.Info(s.sql.String(),s.sqlVars)
	if rows,err = s.DB().Query(s.sql.String(),s.sqlVars...);err != nil{
		log.Error(err)
	}
	return
}