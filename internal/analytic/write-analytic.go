package analytic

import (
	"fmt"
	"strings"
	"time"

	jsonLine "github.com/njfamirm/sitatus/pkg/json-line"
)

type Analytic struct {
	Name string `json:"name"`
	Ping int    `json:"ping"`
	Time int64  `json:"time"`
}

var AnalyticFormatVersion = 0

func WriteAnalytic(name string, ping int) error {
	currentMonthName := strings.ToLower(time.Now().Month().String())
	currentYear := time.Now().Year()
	folderName := fmt.Sprintf("analytic/%d/%s-%d", AnalyticFormatVersion, currentMonthName, currentYear)

	if err := jsonLine.CreateFolder(folderName); err != nil {
		return err
	}

	err := jsonLine.WriteJSONLine(fmt.Sprintf("%s/%s.jsonl", folderName, name), Analytic{
		Name: name,
		Ping: ping,
		Time: time.Now().Unix(),
	})
	return err
}
