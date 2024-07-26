package p2p

import "encoding/json"

func ParserIpAddr(payload []byte) (struct {
	Ip string "json:\"addr\""
}, error) {
	f := struct {
		Ip string `json:"addr"`
	}{}

	if err := json.Unmarshal(payload, &f); err != nil {
		return f, err
	}

	return f, nil
}

func ParserInfo(payload []byte) (struct {
	Info string "json:\"info\""
}, error) {
	f := struct {
		Info string `json:"info"`
	}{}

	if err := json.Unmarshal(payload, &f); err != nil {
		return f, err
	}

	return f, nil
}
