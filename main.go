package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// MYdata is json struct
type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

// Str get string value
func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

var qry string = "select * from mydata where id = ?"

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	s := "2" // input id
	if s == "" {
		panic("no input")
	}
	n, er := strconv.Atoi(s)
	if er != nil {
		panic(er)
	}
	rs, er := con.Query(qry, n)
	if er != nil {
		panic(er)
	}
	for rs.Next() {
		var md Mydata
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}
	fmt.Println("*** end ***")
}
