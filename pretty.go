package tabslib

import "encoding/json"

func PrettyString(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
