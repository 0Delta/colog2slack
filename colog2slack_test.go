package colog2slack

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_colog2slack(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(".golog2slack_incomming_url")
			if err != nil {
				fmt.Println("error")
			}
			defer f.Close()
			b, err := ioutil.ReadAll(f)
			colog2slack(string(b))
			log.Printf("trace: this is a trace log.")
			log.Printf("debug: this is a debug log.")
			log.Printf("info: this is an info log.")
			log.Printf("warn: this is a warning log.")
			log.Printf("error: this is an error log.")
			log.Printf("alert: this is an alert log.")
		})
	}
}
