package angrytweet

import (
	"github.com/op/go-logging"
	"encoding/json"
)

//
// easy generation: http://json2struct.mervine.net/
// easy validation http://jsonlint.com/
//

var logger = logging.MustGetLogger("json-parse")

type Record struct {
	Users []struct {
		Password string `json:"password"`
		URL      string `json:"url"`
		User     string `json:"user"`
	} `json:"users"`
}

func Decode(jsonBlob []byte) (users *Record, err error) {
	users = new(Record)
//	err = json.NewDecoder(r).Decode(record)
	json.Unmarshal(jsonBlob, &users)
	return users, err
}


