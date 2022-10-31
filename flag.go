package main

import "flag"

var (
	flagHashType string
	flagFilePath string
	flagMail     string
)

func readFlags() {
	flag.StringVar(&flagHashType, "HType", "", "Тип хеширования, реализованные типы MD5")
	flag.StringVar(&flagFilePath, "FilePath", "", "Абсолютный путь к файлу")
	flag.StringVar(&flagMail, "Mail", "", "Электронный почтовый адрес Шаблон Email@Domen.* \r\n"+
		"Разработчики: Пантелеев Е.С. Obmorozz@gmail.com, Коротков Д.А.  korimer@gmail.com")
	flag.Parse()
}

func isFlagHashType() bool {
	if flagHashType != "" {
		return true

	}
	return false
}

func isFlagFilePathSet() bool {
	if flagFilePath != "" {
		return true
	}
	return false
}

func isFlagMailSet() bool {
	if flagFilePath != "" {
		return true
	}
	return false
}

func isFlagsWasSet() bool {
	if isFlagFilePathSet() || isFlagMailSet() || isFlagHashType() {
		return true
	}
	return false
}
