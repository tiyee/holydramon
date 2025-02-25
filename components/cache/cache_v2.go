package cache

import "time"

type ICacheM interface {
	ICacheE
  MSet(values map[string][]byte)error
  MSetE(values map[string][]byte, exp time.Duration) error
  MGet(keys ...string)([][]byte,error)
}

var _cacheM ICacheM

func InitCacheM(c ICacheM) {
	_cacheM = c
}
func MSet(values map[string][]byte) error {
  return _cacheM.MSet(values)
}
func MSetE(values map[string][]byte, exp time.Duration) error {
  return _cacheM.MSetE(values, exp)
}
func MGet(keys ...string)([][]byte, error) {
  return _cacheM.MGet(keys...)
}
