package engine

import (
	"database/sql"
	"github.com/allegro/bigcache"
	"time"
)

type Components struct {
	Mysql    *sql.DB
	BigCache *bigcache.BigCache
}

func (e *Engine) InitComponent() error {
	e.components = &Components{}
	if err := e.components.initMysql(&e.Config().Mysql); err != nil {
		return err
	}

	if err := e.components.initBigCache(); err != nil {
		return err
	}
	ImmutableComponents = e.components
	return nil
}
func (cpt *Components) initMysql(config *ConfigMysql) error {
	return nil
	if db, err := sql.Open("mysql", config.Dsn); err == nil {
		db.SetConnMaxLifetime(time.Second * 20)
		db.SetMaxIdleConns(config.MaxIdle)
		db.SetMaxOpenConns(config.MaxOpen)
		if err := db.Ping(); err == nil {
			cpt.Mysql = db
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
func (cpt *Components) initBigCache() error {
	if cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute)); err == nil {
		cpt.BigCache = cache
		return nil
	} else {
		return err
	}
}
