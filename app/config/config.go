/**
 *  basic config
 **/
package config

type MysqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

var MysqlDsn string

func init() {
	mysqlConf := &MysqlConfig{
		"root",
		"lcp0578",
		"127.0.0.1",
		"3306",
		"gradmin",
		"utf8",
	}
	MysqlDsn = mysqlConf.User + ":" + mysqlConf.Password + "@tcp(" + mysqlConf.Host + ":" + mysqlConf.Port + ")/" + mysqlConf.Database + "?charset=" + mysqlConf.Charset
}
