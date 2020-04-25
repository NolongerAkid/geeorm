package day1_database_sql

import (
	dialect2 "awesomeProject4/geeorm/day2-reflect-schema/dialect"
	"awesomeProject4/geeorm/day2-reflect-schema/log"
	"awesomeProject4/geeorm/day2-reflect-schema/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
	dialect dialect2.Dialect
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
	dial,ok := dialect2.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found",driver)
		return
	}

	e = &Engine{db:db,dialect:dial}
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
	return session.New(engine.db,engine.dialect)
}
