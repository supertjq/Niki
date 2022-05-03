package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//Vitamin information
type Vitamin struct {
	Va float64
	Vb float64
	Vc float64
}

//Food information
type Food struct {
	Name    string
	Protein float64
	Fat     float64
	Vita    Vitamin
	Carbon  float64
	Next    *Food
}

//
const (
	USERNAME = "root"
	PASSWORD = "Tjq216828"
	IP       = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "test"
)

var DB *sql.DB
var Head *Food
var Tail *Food

func Getfooddata() *Food {
	InitDB()
	return Getfood()
}

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "") //this is just a path to get your mysql
	DB, _ = sql.Open("mysql", path)                                                                                    //this means to connect to the database  !!!!!!!!!!!
	DB.SetConnMaxLifetime(100)                                                                                         //this means to set the max number of the connection
	DB.SetConnMaxIdleTime(10)                                                                                          //this aims to set the max number of free connection
	if err := DB.Ping(); err != nil {                                                                                  //this is to check whether the database is connected
		fmt.Println("something is going wrong ! :")
		fmt.Println(err)
		return
	}
}

//Code above is to get connected to the database
//
//Code below is to get information from the database
func Getfood() *Food {
	rows, err := DB.Query("select name,Protein,Va,Vb,Vc,Fat,Carbon from food where name = ?", 0)
	if err != nil {
		fmt.Println("DB,query() error : ", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		newfood := new(Food)
		err2 := rows.Scan(&newfood.Name, &newfood.Protein, &newfood.Vita.Va, &newfood.Vita.Vb, &newfood.Vita.Vc, &newfood.Fat, &newfood.Carbon)
		if err2 != nil {
			fmt.Println("rows.Scan() error : ", err2)
			return nil
		}
		if Head == nil && Tail == nil {
			Head = newfood
			Tail = Head
			newfood.Next = nil
		} else {
			Tail.Next = newfood
			Tail = Tail.Next
		}
	}
	return Head
}
