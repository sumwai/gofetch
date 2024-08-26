package gofetch

import (
	"encoding/json"
	"net/url"
)

type Params map[interface{}]interface{}

func (p Params) String() string {
	var data string
	for k, value := range p {
		var v string
		switch value.(type) {
		case string:
			v = value.(string)
		default:
			ret, err := json.Marshal(value)
			if err != nil {
				continue
			}
			v = string(ret)
		}
		data = data + k.(string) + "=" + url.QueryEscape(v) + "&"
	}
	if len(data) > 0 {
		data = data[:len(data)-1]
	}
	return data
}
