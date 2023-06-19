package jwt

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/tiyee/holydramon/components/crypto/base64"
	"github.com/tiyee/holydramon/components/crypto/hmac"
)

type IPayload interface {
	Encode() []byte
	Decode([]byte) error
}
type Header struct {
	Algorithm string `json:"alg"`
	Typ       string `json:"typ"`
	Expired   int64  `json:"exp"`
	Audience  string `json:"aud"`
}
type IOpt[T IPayload] func(j *JWT[T])
type JWT[T IPayload] struct {
	header Header
	secret []byte
}

func New[T IPayload](secret []byte, opts ...IOpt[T]) *JWT[T] {
	jwt := &JWT[T]{
		header: Header{
			Algorithm: "HS256",
			Typ:       "JWT",
		},
		secret: secret,
	}
	for _, opt := range opts {
		opt(jwt)
	}
	return jwt
}
func (j *JWT[T]) SetHeader(h Header) {
	j.header = h
}
func (j *JWT[T]) Header() Header {
	return j.header
}
func (j *JWT[T]) Encode(pl T) []byte {
	var header, payload []byte
	if bs, err := json.Marshal(j.header); err == nil {
		header = base64.UrlEncode(bs)
	}

	payload = base64.UrlEncode(pl.Encode())
	body := make([]byte, 0, len(header)+1+len(payload)+1+len(j.secret))
	body = append(body, header...)
	body = append(body, '.')
	body = append(body, payload...)

	signature := hmac.Sha256Encrypt(body, j.secret)

	body = append(body, '.')
	body = append(body, signature...)
	buff := bytes.NewBufferString("Bearer ")
	buff.Write(body)
	return buff.Bytes()

}
func (j *JWT[T]) Decode(bs []byte, pl T) error {
	if bytes.Index(bs, []byte("Bearer ")) != 0 {
		return errors.New("invalid token")
	}
	bs = bs[7:]
	jwt := bytes.Split(bs, []byte{'.'})
	if len(jwt) != 3 {
		return errors.New("invalid token")
	}
	if bs, err := base64.UrlDecode(jwt[0]); err == nil {
		if err := json.Unmarshal(bs, &j.header); err != nil {
			return err
		}
	} else {
		return err
	}
	body := [][]byte{jwt[0], jwt[1]}
	if signature := hmac.Sha256Encrypt(bytes.Join(body, []byte{'.'}), j.secret); string(signature) == string(jwt[2]) {
		if bs, err := base64.UrlDecode(jwt[1]); err == nil {
			if err := pl.Decode(bs); err == nil {
				return nil
			} else {
				return err
			}
		} else {
			return err
		}

	} else {
		return errors.New("dismatch token")
	}
}
func Encode[T IPayload](payload T, secret []byte) []byte {
	j := New[T](secret)
	return j.Encode(payload)
}
func Decode[T IPayload](bs []byte, secret []byte, payload T) error {
	j := New[T](secret)
	return j.Decode(bs, payload)
}
