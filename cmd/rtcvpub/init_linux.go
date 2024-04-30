package main

import (
	"github.com/ahaooahaz/rtcvpub/cmd/rtcvpub/child/version"
	"github.com/ahaooahaz/rtcvpub/cmd/rtcvpub/child/webrtc"

	"github.com/sirupsen/logrus"
)

func init() {
	err := initEnv()
	if err != nil {
		panic(err.Error())
	}

	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(webrtc.Cmd)
}

func initEnv() (err error) {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return
}
