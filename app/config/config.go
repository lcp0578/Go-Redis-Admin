/**
 *  basic config
 **/
package config

type MysqlConfig struct {
	User     string
	Password string
	Ip       string
	Port     int32
	Datebase string
	Charset  string
}

func init() {

}
