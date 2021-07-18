package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nqtinh/go-gin-project/pkg/util"
	"github.com/sirupsen/logrus"
)

func withBuilder(ctx context.Context, builder *sqlx.DB) context.Context {
	return context.WithValue(ctx, util.ContextKeyTx, builder)
}

// TransactionHandler is middleware that wraps current request within main DB transaction
func TransactionHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		logrus.Infof("DB stats: {%s}", dbStats(db.Stats()))

		// else we spawn a new tx
		logrus.Info("TransactionHandler: BeginTxx")
		tx, err := db.BeginTxx(ctx, nil)
		if err != nil {
			return
		}

		// wrap sqlx.Tx into TxLogger
		txLogger := &util.TxLogger{
			Tx:     tx,
			Logger: logrus.StandardLogger(),
		}
		ctx = context.WithValue(ctx, util.ContextKeyTx, txLogger)
		// set current tx to sharable context key
		c.Request = c.Request.WithContext(ctx)

		defer func() {
			if p := recover(); p != nil {
				txLogger.Tx.Rollback()
				logrus.Warn("TransactionHandler: Rollbacked")
				// propagate panic error to main Recovery
				panic(p)
			}
		}()

		c.Next()
		if c.Errors.Last() != nil {
			rerr := txLogger.Tx.Rollback()
			if rerr != nil {
				return
			}
			logrus.Warn("TransactionHandler: Rollbacked")
		}

		if err := txLogger.Tx.Commit(); err != nil {
			return
		}

		logrus.Info("TransactionHandler: Committed")
		return
	}
}

func dbStats(stats sql.DBStats) string {
	return fmt.Sprintf(`OpenConnections=%d, InUse=%d, Idle=%d, WaitDuration=%s, WaitCount=%d`, stats.OpenConnections, stats.InUse, stats.Idle, stats.WaitDuration, stats.WaitCount)
}
