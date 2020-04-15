package main

import (
	day1_database_sql "awesomeProject4/geeorm/day1-database-sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine,_ := day1_database_sql.NewEngine("sqlite3","gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_,_ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_,_ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_,_ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result,_ := s.Raw("INSERT INTO User(`Name`) values (?), (?)","Tom","Sam").Exec()
	count,_ := result.RowsAffected()
	fmt.Printf("Exec success,%d affected \n",count)


}
