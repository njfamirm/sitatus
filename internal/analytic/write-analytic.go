package analytic

import jsonLine "github.com/njfamirm/sitatus/pkg/json-line"

func WriteAnalytic(name string, ping int) error {
	err := jsonLine.WriteJSONLine("analytic/"+name+".jsonl", map[string]interface{}{
		"name": name,
		"ping": ping,
	})
	return err
}
