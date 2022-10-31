package main

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
	"runtime"
)

var Mailer *smtpclient.SMTP

func main() {
	readFlags()
	if !isFlagsWasSet() {
		fmt.Println("Для работы приложения необходимо подать все аргументы\r\n" +
			"для получения справки введите флаг -help")
		os.Exit(1)
	}

	RootDir, err := config.RootDirIdentification()
	if err != nil {
		fmt.Printf("ошибка при определении корневой дирректории : " + err.Error())
	}

	err = initer(RootDir)
	if err != nil {
		fmt.Printf("ошибка при иницмализации программы : " + err.Error())
	}

	var hashString string
	switch flagHashType {
	case "MD5":
		f, err := os.OpenFile(flagFilePath, os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println("Lib open file err  " + err.Error())
		}
		defer f.Close()
		h := md5.New()
		//TODO Переделать на потоковое чтение
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println("Lib read to hasher " + err.Error())
		}
		hashString = fmt.Sprintf("%x", h.Sum(nil))
	default:
		log.Printf("неизвестный тип хеширования : " + flagHashType)
	}

	message := "Контрольная сумма файла " + flagFilePath + " \r\n" +
		"Алгоритм хеширования  " + flagHashType + "\r\n" +
		"хеш " + hashString

	log.Printf(message)

	log.Printf("отправка сообщения в почту")
	err = Mailer.Send(flagMail, "Проверка КС из "+runtime.GOOS, message)
	if err != nil {
		fmt.Println("ошибка при отправке почты" + err.Error())
	}
	log.Printf("почта отправлена")

}

func initer(rootDir string) error {
	conf := config.GetConfig()
	f, err := os.ReadFile(path.Join(rootDir, "conf.yaml"))
	if err != nil {
		fmt.Println("Lib open file err  " + err.Error())
	}
	err = yaml.Unmarshal(f, conf)
	if err != nil {
		fmt.Println("ошибка при разборе файла конфигурации проекта  " + err.Error())
	}
	smtp, err := smtpclient.NewSMTP(
		conf.Mailer.Smtp.Host,
		conf.Mailer.Smtp.Port,
		conf.Mailer.Smtp.From,
		conf.Mailer.Smtp.Pass)
	if err != nil {
		fmt.Println("ошибка при инициализации почты  " + err.Error())

	}
	Mailer = smtp
	log.Printf("инициализировано")
	return nil
}
