package mysqlclusterinit

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	// support mysql
	_ "github.com/go-sql-driver/mysql"
)

// InitMySQLCluster 初始化MySQL Master-Master集群
func (s Settings) InitMySQLCluster() (r Result, err error) {
	if s.ValidateAndSetDefault(Validate, SetDefault) != nil {
		os.Exit(1)
	}

	if r.Sqls, err = s.createMySQCluster(); err != nil {
		return r, err
	}

	r.HAProxy = s.createHAProxyConfig()

	if s.Debug {
		logrus.Infof("SQL:%s", strings.Join(r.Sqls, ";\n"))
		logrus.Infof("HAProxy:%s", r.HAProxy)

		return r, err
	}

	if err := s.overwriteHAProxyCnf(&r); err != nil {
		return r, err
	}

	if err := s.restartHAProxy(); err != nil {
		return r, err
	}

	return r, err
}
