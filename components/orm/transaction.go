package orm

import (
	"github.com/tiyee/holydramon/components"
	"github.com/tiyee/holydramon/components/log"
	"github.com/tiyee/holydramon/components/orm"
)

func AutoTransaction(trans func(db orm.IDBHandle) error) error {
	conn, err := components.WDb.Begin()
	if err != nil {
		log.Error("transaction begin err", log.String("error", err.Error()))
		return err
	}
	if err = trans(conn); err != nil {
		log.Error("transaction exec failed",
			log.String("error", err.Error()),
		)
		if rbErr := conn.Rollback(); rbErr != nil {
			log.Error("transaction Rollback failed", log.String("error", rbErr.Error()))
		}
		return err
	}
	if cmtErr := conn.Commit(); cmtErr != nil {
		log.Error("transaction Commit failed", log.String("error", cmtErr.Error()))
		return cmtErr
	}
	return nil

}
