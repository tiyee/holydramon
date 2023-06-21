package jwt

import (
	"testing"
)

func TestUserToken(t *testing.T) {
	user := User{Uid: 1, Name: "test"}
	j := New[*User]([]byte("hello world"))
	bs := j.Encode(&user)
	u := User{Uid: 0, Name: ""}
	if err := j.Decode(bs, &u); err == nil {
		if u.Uid != user.Uid {
			t.Errorf("test error\n")
		}
	}
}
