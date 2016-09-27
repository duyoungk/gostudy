package main

import (
	"database/sql"
	"fmt"

	"container/list"

	_ "github.com/denisenkom/go-mssqldb"
	"strconv"
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

func rowsToList(rows *sql.Rows) *list.List {
	results := list.New()
	
	cols, _ := rows.Columns()
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
	}

	for rows.Next() {
		err := rows.Scan(vals...)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		row := make(map[string]interface{})

		for i := 0; i < len(vals); i++ {

			r := vals[i].(*interface{})

			switch v := (*r).(type) {
			case nil:
				row[cols[i]] = nil
			case bool:
				if v {
					row[cols[i]] = true
				} else {
					row[cols[i]] = false
				}
			case []byte:
				row[cols[i]] = string(v)
			default:
				row[cols[i]] = v
			}
		}
		results.PushBack(row)
	}

	return results
}

func (p *Query) Query(query string) *list.List {
	if p.db == nil {
		return nil
	}

	rows, err := p.db.Query(query)
	if err != nil {
		fmt.Println("cannot query:", err.Error())
		return nil
	}


	return rowsToList(rows)

}

func (p *Query) Proc(query string, args ...interface{}) *list.List {
	
	if len(args) > 0 {
		for i := 1; i <= len(args); i++ {
			if i > 1 {
				query += ","
			}
			query += " ?" + strconv.Itoa(i)
		}
	}
	
	stmt, err := p.db.Prepare(query)
	if err != nil {
		fmt.Println("Prepare:", err.Error())
		return nil
	}
	defer stmt.Close()
	
	rows, err := stmt.Query(args...)
	if err != nil {
		fmt.Println("Query:", err.Error())
		return nil
	}
	
	return rowsToList(rows)
}

func NewQuery() *Query {
	return &Query{}
}