package engine

import (
	"fmt"
	"os"
	"regexp"
)

const PosBefore = 1
const PosAfter = 2
const MatchCmp = 1
const MatchReg = 2

type HookMatcher struct {
	adapter int
	value   string
	reg     *regexp.Regexp
}

func NewHookMatcher(value string, rule int) *HookMatcher {
	hm := &HookMatcher{
		value: value,
	}

	switch rule {
	case MatchReg:
		hm.adapter = MatchReg
		if reg_, err := regexp.Compile(value); err == nil {
			hm.reg = reg_
		} else {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		break
	case MatchCmp:

		hm.adapter = MatchCmp
		break
	default:
		fmt.Println("undefined HookMatcher rule")
		os.Exit(1)
	}

	return hm

}
func (rm *HookMatcher) Match(url string) bool {
	if rm.adapter == MatchCmp {
		return rm.value == url
	}
	if rm.adapter == MatchReg {
		return rm.reg.Match([]byte(url))
	}
	return false

}

type Hook struct {
	include     []*HookMatcher
	exclude     []*HookMatcher
	Pos         int
	HandlerFunc HandlerFunc
}

func NewHook(pos int, handle HandlerFunc, includes []*HookMatcher, excludes []*HookMatcher) *Hook {
	if includes == nil{
		includes = make([]*HookMatcher,0)
	}
	if excludes ==nil{
		excludes = make([]*HookMatcher,0)
	}
	return &Hook{
		include:     includes,
		exclude:     excludes,
		Pos:         pos,
		HandlerFunc: handle,
	}
}
func (hk *Hook) Include(value string, rule int) *Hook {
	hk.include = append(hk.include, NewHookMatcher(value, rule))
	return hk
}
func (hk *Hook) Exclude(value string, rule int) *Hook {
	hk.exclude = append(hk.exclude, NewHookMatcher(value, rule))
	return hk
}
func (hk *Hook) Match(rt *Route) bool {
	for _, matcher := range hk.exclude {
		if matcher.Match(rt.Path) {
			return false
		}
	}
	for _, matcher := range hk.include {
		if matcher.Match(rt.Path) {
			return true
		}
	}
	return false
}
