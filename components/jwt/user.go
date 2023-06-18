package jwt

import (
	"bytes"
	"encoding/json"
	"github.com/tiyee/holydramon/components/crypto/base64"
)

type User struct {
	Uid  int64  `json:"u"`
	Name string `json:"n"`
}

func (u *User) Encode() []byte {
	buff := bytes.Buffer{}
	if bs, err := json.Marshal(u); err == nil {
		buff.Write(base64.UrlEncode(bs))
	}
	return buff.Bytes()

}
func (u *User) Decode(bs []byte) error {
	return json.Unmarshal(bs, u)
}
