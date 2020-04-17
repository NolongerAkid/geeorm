package day1_database_sql

import (
	"awesomeProject4/geeorm/day1-database-sql/log"
	"awesomeProject4/geeorm/day1-database-sql/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver,source string)(e *Engine,err error){
	db,err := sql.Open(driver,source)
	if err != nil {
		log.Error(err)
		return
	}
	//Send a ping to make sure the database connection is alive
	if err = db.Ping();err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db:db}
	log.Info("Connect database success")
	return
}

func(engine *Engine) Close(){
	if err := engine.db.Close();err != nil{
		log.Error("Fail to close database")
	}
	log.Info("Close database success")
}
func (engine *Engine) NewSession() *session.Session{
	return session.New(engine.db)
}
