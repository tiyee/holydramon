package base64

import "encoding/base64"

func UrlEncode(src []byte) []byte {
	env := base64.URLEncoding
	buf := make([]byte, env.EncodedLen(len(src)))
	env.Encode(buf, src)
	return buf
}
func UrlDecode(src []byte) ([]byte, error) {
	enc := base64.URLEncoding
	dbuf := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dbuf, src)
	return dbuf[:n], err
}
