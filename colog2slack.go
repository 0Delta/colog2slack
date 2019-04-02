package colog2slack

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/comail/colog"
)

type myHook struct {
	levels []colog.Level
}

func (h *myHook) Levels() []colog.Level {
	return h.levels
}

type slackMessageFormat struct {
	Attachments []slackMessageAttachmentFormat `json:"attachments"`
}
type slackMessageAttachmentFormat struct {
	Text      string   `json:"text"`
	Mrkdwn_in []string `json:"mrkdwm_in"`
	Footer    string   `json:"footer"`
	Ts        int64    `json:"ts"`
	Color     string   `json:"color"`
}

func (this *slackMessageFormat) AddAttachment(e *colog.Entry) {
	colorMap := map[colog.Level]string{
		colog.LAlert:   "#f42000",
		colog.LError:   "#ff6851",
		colog.LWarning: "#ffcc66",
		colog.LInfo:    "#36a64f",
		colog.LDebug:   "#81e1e1",
		colog.LTrace:   "#cccccc",
	}

	msg := string(e.Message)
	if e.Level == colog.LAlert {
		msg = "<!here>\n" + msg
	}

	this.Attachments = append(this.Attachments,
		slackMessageAttachmentFormat{
			Text:      msg,
			Mrkdwn_in: []string{"text"},
			Footer:    filepath.Base(e.File),
			Ts:        e.Time.Unix(),
			Color:     colorMap[e.Level],
		})

}

func GetSlackMsgFmt() slackMessageFormat {
	return slackMessageFormat{}
}

func (h *myHook) Fire(e *colog.Entry) error {
	// fmt.Println(e.Level, e.Time)
	// fmt.Println(filepath.Base(e.File))
	// fmt.Printf("We got an entry: \n%#v", e)

	msg := GetSlackMsgFmt()
	msg.AddAttachment(e)
	err := post2slack(slack_incommig_url, msg)
	if err != nil {
		return err
	}
	return nil
}

func post2slack(url string, message slackMessageFormat) error {
	jsonbyte, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonbyte),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

var slack_incommig_url = ""

func Enable(ImcommigUrl string) {
	slack_incommig_url = ImcommigUrl
	hook := &myHook{
		levels: []colog.Level{
			colog.LAlert,
			colog.LError,
			colog.LWarning,
			colog.LInfo,
			colog.LDebug,
			colog.LTrace,
		},
	}
	colog.AddHook(hook)
	colog.Register()
}
