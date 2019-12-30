package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

const DBMS = "mysql"
const  USER = "docker"
const  PASS = "docker"
const DBNAME = "test_database"

func DBInit()  {
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("データベース開けず！（dbInit）")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()

}

//DB追加
func DbInsert(text string, status string) {
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func DbUpdate(id int, text string, status string) {
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("データベース開けず！（dbUpdate)")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func DbDelete(id int) {
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("データベース開けず！（dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func DbGetAll() []Todo {
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func DbGetOne(id int) Todo {
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}