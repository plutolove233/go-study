package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("one dog appeared. ")
}
