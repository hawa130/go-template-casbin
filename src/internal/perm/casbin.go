package perm

import (
	"errors"

	"github.com/casbin/casbin/v2"
	entadapter "github.com/casbin/ent-adapter"
	zaplogger "github.com/casbin/zap-logger/v2"
	"github.com/hawa130/computility-cloud/internal/logger"
	"go.uber.org/zap"
)

var (
	adapter  *entadapter.Adapter
	enforcer *casbin.Enforcer
)

func Init(driverName, dataSourceName string) error {
	if err := initAdapter(driverName, dataSourceName); err != nil {
		return err
	}

	l := logger.Logger()
	if l == nil {
		return errors.New("logger is not initialized")
	}

	if err := initEnforcer(l); err != nil {
		return err
	}
	return nil
}

func initAdapter(driverName, dataSourceName string) error {
	a, err := entadapter.NewAdapter(driverName, dataSourceName)
	if err != nil {
		return err
	}
	adapter = a
	return nil
}

func initEnforcer(logger *zap.Logger) error {
	if adapter == nil {
		return errors.New("adapter is not initialized")
	}

	eLogger := zaplogger.NewLoggerByZap(logger, true)
	e, err := casbin.NewEnforcer("perm-model.conf", adapter, eLogger)
	if err != nil {
		return err
	}

	enforcer = e
	return nil
}

func Enforcer() *casbin.Enforcer {
	return enforcer
}
