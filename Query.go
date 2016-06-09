package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Query struct {
	dbId       int
	connString string
	db         *sql.DB
}

func (p *Query) Open(dbId int, connstr string) error {
	var err error
	p.connString = connstr
	p.dbId = dbId
	p.db, err = sql.Open("mssql", connstr)
	if err != nil {
		fmt.Println("cannot open database:", err.Error())
		return err
	}

	return nil
}

func (p *Query) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *Query) Query(query string) error {
	if p.db == nil {
		return nil
	}

	rows, err := p.db.Query(query)
	if err != nil {
		fmt.Println("cannot query:", err.Error())
		return err
	}

	cols, _ := rows.Columns()
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		for i := 0; i < len(vals); i++ {
			switch v := vals[i].(type) {
			default:
                fmt.Print(v)
			}
            fmt.Printf("\t")
		}
        fmt.Println()

	}

	return nil
}
