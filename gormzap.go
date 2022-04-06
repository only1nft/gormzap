package gormzap

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type GormZap struct {
	L *zap.SugaredLogger
	logger.Interface
}

func New(l *zap.Logger) GormZap {
	return GormZap{L: l.Sugar()}
}

func (gz GormZap) Info(_ context.Context, msg string, args ...interface{}) {
	gz.L.Info(append([]interface{}{msg}, args...)...)
}

func (gz GormZap) Warn(_ context.Context, msg string, args ...interface{}) {
	gz.L.Warn(append([]interface{}{msg}, args...)...)
}

func (gz GormZap) Error(_ context.Context, msg string, args ...interface{}) {
	gz.L.Error(append([]interface{}{msg}, args...)...)
}

func (gz GormZap) Trace(_ context.Context, _ time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	gz.L.With(
		zap.String("sql", sql),
		zap.Int64("rowsAffected", rowsAffected),
		zap.Error(err),
	).Debug()
}
