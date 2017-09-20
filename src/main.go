package main

import (
	//"net/http"
	//"github.com/go-chi/chi"
	//"./API_v1"
	//"github.com/go-xorm/xorm"
	//_ "github.com/lib/pq"
	//"fmt"
	"time"
)

type User struct {
	Id int64
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	//r := chi.NewRouter()
	//
	//// Mount the api sub-router
	//r.Mount("/api/v1", API_v1.ApiRouter())
	//
	//http.ListenAndServe(":8080", r)
	//engine, err := xorm.NewEngine("postgres", "dbname=xorm_test host=127.0.0.1 port=5432 sslmode=disable");
	//if (err != nil)	{
	//	fmt.Print(err)
	//}
	//defer engine.Close()
	//
	//if _, err = engine.Exec("SET SEARCH_PATH TO \"abc\""); err != nil {
	//	fmt.Print(err)
	//}
	//
	//if err = engine.DB().Ping(); err != nil {
	//	fmt.Print(err)
	//}
	//
	//err = engine.Sync2(new(User))
	//if (err != nil)	{
	//	fmt.Print(err)
	//}
	//
	//user := User{Name:"Hov"}
	//affected, err := engine.Insert(&user)
	//if (err != nil)	{
	//	fmt.Print(err)
	//}
	//fmt.Print(affected)

}
