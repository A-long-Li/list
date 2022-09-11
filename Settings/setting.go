/**
 *@Description
 *@author          liyajun
 *@create          2022-09-11 15:41
 */

package Settings

import "gopkg.in/ini.v1"

var (
	Conf = new(AppConfig)
)

type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
}

type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DataBase string `ini:"database"`

	Host string `ini:"host"`
	Port int    `ini:"port"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
