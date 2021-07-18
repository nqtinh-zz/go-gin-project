package postgresql

import (
	"math"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/nqtinh/go-gin-project/pkg/setting"
	"github.com/sirupsen/logrus"

	// Postgres driver for the database/sql package
	_ "github.com/lib/pq"
)

var (
	dbErr error
	db    *sqlx.DB
	once  sync.Once
)

// GetInstance return a connection to current database instance
func GetInstance() (*sqlx.DB, error) {
	once.Do(func() {
		db, dbErr = sqlx.Connect("postgres", setting.Config.DSN)
		if dbErr != nil {
			return
		}

		logrus.Infof("init: postgresql connected on %s.", setting.Config.DSN)

		// we applied 3 rules:
		// #1: SetMaxOpenConns
		// #2: SetMaxIdleConns with ratio of 50% by default, set by DbMaxIdleConnsRate
		// #3: SetConnMaxLifetime should be tuned to periodically re-established idle conns if SetMaxIdleConns
		// Ref: https://making.pusher.com/production-ready-connection-pooling-in-go/
		maxIdleConns := int(math.Floor(float64(setting.Config.DbMaxOpenConns) * setting.Config.DbMaxIdleConnsRate))
		db.SetMaxOpenConns(setting.Config.DbMaxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
		db.SetConnMaxLifetime(setting.Config.DbConnMaxLifetime)

		logrus.Infof("PostgreSQL configuration: SetMaxOpenConns=%d, SetMaxIdleConns=%d, SetConnMaxLifetime=%s", setting.Config.DbMaxOpenConns, maxIdleConns, setting.Config.DbConnMaxLifetime.String())
	})

	return db, dbErr
}
