package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func DifferentLevels(level logrus.Level) {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(level)

	logrus.WithFields(logrus.Fields{
		"genre": "metal",
		"name":  "Rammstein",
	}).Info("Ich Will")

	logrus.WithFields(logrus.Fields{
		"genre": "post-grunge",
		"name":  "Garbage",
	}).Warn("I Think I’m Paranoid")

	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "Any music is awesome",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I will be logged with common and other fields")
	contextLogger.Error("Me too, maybe")

	logrus.WithFields(logrus.Fields{
		"genre": "rock",
		"name":  "The Rasmus",
	}).Fatal("Livin' in a World Without You")

}

func main() {

	DifferentLevels(logrus.DebugLevel)

	// создаём файл info.log и обрабатываем ошибку
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	// откладываем закрытие файла
	defer file.Close()

	// устанавливаем вывод логов в файл
	logrus.SetOutput(file)
	// устанавливаем вывод логов в формате JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// устанавливаем уровень предупреждений
	logrus.SetLevel(logrus.WarnLevel)

	// определяем стандартные поля JSON
	logrus.WithFields(logrus.Fields{
		"genre": "metal",
		"name":  "Rammstein",
	}).Info("Немецкая метал-группа, образованная в январе 1994 года в Берлине.")

	logrus.WithFields(logrus.Fields{
		"omg":  true,
		"name": "Garbage",
	}).Warn("В 2021 году вышел новый альбом No Gods No Masters.")

	logrus.WithFields(logrus.Fields{
		"omg":  true,
		"name": "Linkin Park",
	}).Fatal("Группа Linkin Park взяла паузу после смерти вокалиста Честера Беннингтона 20 июля 2017 года.")
}
