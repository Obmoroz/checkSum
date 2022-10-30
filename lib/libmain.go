package main

import "C"
import (
	"checkSum/config"
	"checkSum/email/smtpclient"
	"crypto/md5"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path"
)

var Mailer *smtpclient.SMTP

//export MakeHash
func MakeHash(HType, FilePath, Mail string) {
	var hashString string
	switch HType {
	case "MD5":
		f, err := os.OpenFile(FilePath, os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println("Lib open file err  " + err.Error())
		}
		defer f.Close()
		h := md5.New()
		//TODO Переделать на потоковое чтение
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println("Lib read to hasher " + err.Error())
		}
		log.Printf("Call RunLib: M from lib %x ", h.Sum(nil))
		hashString = fmt.Sprintf("%x", h.Sum(nil))
	default:
		log.Printf("неизвестный тип хеширования : " + HType)
	}

	err := Mailer.Send(Mail, "Проверка КС", "Контрольная сумма файла "+FilePath+" \r\n"+
		"Алгоритм хеширования  "+HType+"\r\n"+
		"хеш "+hashString)
	if err != nil {
		fmt.Println("ошибка при отправке почты" + err.Error())
	}
	fmt.Println("почта отправлена")

}

func init() {
	conf := config.GetConfig()

	ROOTDIR, err := config.RootDirIdentification()
	if err != nil {
		fmt.Println("ошибка при определении корневой дирректории  " + err.Error())
	}
	f, err := os.ReadFile(path.Join(ROOTDIR, "conf.yaml"))
	if err != nil {
		fmt.Println("Lib open file err  " + err.Error())
	}
	err = yaml.Unmarshal(f, conf)
	if err != nil {
		fmt.Println("ошибка при разборе файла конфигурации проекта  " + err.Error())
	}
	fmt.Printf("%+v", conf)

	smtp, err := smtpclient.NewSMTP(
		conf.Mailer.Smtp.Host,
		conf.Mailer.Smtp.Port,
		conf.Mailer.Smtp.From,
		conf.Mailer.Smtp.Pass)
	if err != nil {
		fmt.Println("ошибка при инициализации почты  " + err.Error())

	}
	Mailer = smtp
	fmt.Println("инициалезированно")

}

func main() {}
