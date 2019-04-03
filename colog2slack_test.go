package colog2slack

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/comail/colog"
)

func GetTestToken(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("error")
		return ""
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error")
		return ""
	}
	return string(b)
}

func Test_colog2slack(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		Enable(GetTestToken(".golog2slack_incomming_url"),
			colog.LError, colog.LAlert)
		log.Printf("trace: <NotPost> this is a trace log.")
		log.Printf("debug: <NotPost> this is a debug log.")
		log.Printf("info: <NotPost> this is an info log.")
		log.Printf("warn: <NotPost> this is a warning log.")
		log.Printf("error: this is an error log.")
		log.Printf("alert: this is an alert log.")
	})
}
