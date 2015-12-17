package main

import (
	"github.com/op/go-logging"
	"io/ioutil"
	"os"
	"github.com/bluemonk3y/angrytweet"
)

var logger = logging.MustGetLogger("json-parse-test")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var dir, error0 = os.Getwd()

	check(error0)
	logger.Info("Opening JSON file from WorkingDir:" + dir)

	var byte, error1  = ioutil.ReadFile(".\\users.json")
	check(error1)

	var users, error2 = angrytweet.Decode(byte)

	check(error2)

//	println("USers:" + users.)
	for index, each := range users.Users {
		logger.Info("Iterate %i / %s", index, each)
	}
	logger.Info("Got Users: %s", users)

}