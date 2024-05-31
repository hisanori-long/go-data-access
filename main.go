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

// var qry string = "select * from mydata where name like ? or mail like ?"

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	ids := "3" // input
	id, er := strconv.Atoi(ids)
	if er != nil {
		panic(er)
	}
	qry := "select * from mydata where id = ?"
	rw := con.QueryRow(qry, id)
	tgt := mydatafmRw(rw)
	nm := "JiroJiro"  // input
	ml := "jiro@jiro" // input
	ge := "25"        // input
	ag, _ := strconv.Atoi(ge)

	if nm == "" {
		nm = tgt.Name
	}
	if ml == "" {
		ml = tgt.Mail
	}
	if ge == "" {
		ag = tgt.Age
	}

	qry = "update mydata set name = ?, mail = ?, age = ? where id = ?"
	con.Exec(qry, nm, ml, ag, id)
	showRecord(con)
}

// print all records.
func showRecord(con *sql.DB) {
	qry := "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str())
	}
}

// get mydata from Rows.
func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

// get Mydata from Row.
func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}
