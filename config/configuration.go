package config

import (
	"os"
)

/*
**	Возвращает структуру настроек - без заполнения ее полей пакет считается неинициализированным
 */

func GetConfig() *Configuration {
	if Conf == nil {
		Conf = &Configuration{}
	}
	return Conf
}

var (
	Conf *Configuration
	//Db  [] DbConnect
)

type Configuration struct {
	Mailer Mailer `yaml:"Mailer"`
}

type Mailer struct {
	Smtp Smtp `yaml:"Smtp"`
}
type Smtp struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	From string `yaml:"From"`
	Pass string `yaml:"Pass"`
}

/*
Функция опредления корневой папки проекта
runtime.Caller выдает путь к папки модуля в которой лежит файл исходного кода
os.Executable выдает путь к папке в которой был запущен бинарник
*/

func RootDirIdentification() (string, error) {
	var rootDir string
	rootDir, err := os.Getwd()
	if err != nil {
		return rootDir, err
	}
	return rootDir, err

}
