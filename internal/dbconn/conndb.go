package dbconn

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type ConnDB struct {
	*xorm.Engine
}

func Get() (conn *ConnDB) {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "txvier", "txvier", "localhost", 3306, "txvier")
	engine, err := xorm.NewEngine("mysql", connArgs)
	if err != nil {
		engine.Close()
		panic("failed to connect database")
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(5)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	engine.SetTableMapper(tbMapper)
	return &ConnDB{
		Engine: engine,
	}
}

func (conn *ConnDB) WrapTX(f func() error) {
	session := conn.Engine.NewSession()
	defer session.Close()
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := f(); err != nil {
		panic(err)
	}
	session.Commit()
}
