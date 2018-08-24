package dao

import (
	"database/sql"
	"fmt"

	"github.com/liuhengloveyou/GSLB/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type RR struct {
	ID         int
	Domain     string
	Ttl        uint32
	Type       uint16
	Class      uint16
	Data       sql.NullString
	Group      sql.NullString
	UpdateTime sql.RawBytes `db:"update_time"`
}

func init() {
	var e error

	if db, e = sqlx.Connect("mysql", common.ServConfig.Mysql); e != nil {
		panic(e)
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	if e = db.Ping(); e != nil {
		panic(e)
	}
}

func SelectRRsFromMysql(d []string) (rr []*common.RR, e error) {
	r := []RR{}

	sql := "SELECT * FROM ns.rr where domain in ('" + d[0] + "'"
	for i := 1; i < len(d); i++ {
		sql = sql + ", '" + d[i] + "'"
	}
	sql = sql + ");"
	common.Logger.Debug("SelectRRsFromMysql: " + sql)

	e = db.Select(&r, sql)
	common.Logger.Info(fmt.Sprintf("SelectRRsFromMysql end: %v %v", r, e))
	if e != nil {
		return
	}

	for i := 0; i < len(r); i++ {
		t := &common.RR{
			ID:     r[i].ID,
			Domain: r[i].Domain,
			Ttl:    r[i].Ttl,
			Type:   r[i].Type,
			Class:  r[i].Class,
		}

		if r[i].Data.Valid {
			t.Data = r[i].Data.String
		}

		if r[i].Group.Valid {
			t.Group = r[i].Group.String
		}

		rr = append(rr, t)
	}

	common.Logger.Info(fmt.Sprintf("SelectRRsFromMysql ended: %#v %d\n", rr, len(rr)))
	return rr, nil
}

func LoadRRFromMysql(t string) (rr []*common.RR, e error) {
	r := []RR{}

	sql := "select *  from rr where update_time >= " + t
	common.Logger.Debug("LoadRRFromMysql: " + sql)

	e = db.Select(&r, sql)
	common.Logger.Info(fmt.Sprintf("LoadRRFromMysql end: %v %v", r, e))
	if e != nil {
		return
	}

	for i := 0; i < len(r); i++ {
		t := &common.RR{
			ID:     r[i].ID,
			Domain: r[i].Domain,
			Ttl:    r[i].Ttl,
			Type:   r[i].Type,
			Class:  r[i].Class,
		}

		if r[i].Data.Valid {
			t.Data = r[i].Data.String
		}

		if r[i].Group.Valid {
			t.Group = r[i].Group.String
		}

		rr = append(rr, t)
	}

	common.Logger.Info(fmt.Sprintf("LoadRRFromMysql ended: %#v %d\n", rr, len(rr)))
	return rr, nil
}

func SelectRulesFromMysql(domains []string) (rules []*common.Rule, e error) {
	r := []common.Rule{}

	sql := "SELECT * FROM ns.rule where domain in ('" + domains[0] + "'"
	for i := 1; i < len(domains); i++ {
		sql = sql + ", '" + domains[i] + "'"
	}
	sql = sql + ");"
	common.Logger.Debug("SelectRulesFromMysql: " + sql)

	e = db.Select(&r, sql)
	common.Logger.Info(fmt.Sprintf("SelectRulesFromMysql end: %v %v", r, e))
	if e != nil {
		return
	}

	return nil, nil
}
