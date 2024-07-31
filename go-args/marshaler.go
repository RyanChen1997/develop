package main

import (
	"strconv"
	"strings"
)

type BoolMarshaler struct {
	V bool
}

func (b *BoolMarshaler) Set(v interface{}) {
	b.V = toBool(v)
}

func (b *BoolMarshaler) GetValue(k string) interface{} {
	return b.V
}

func (b *BoolMarshaler) ValueValid(v interface{}) bool {
	_, ok := v.(bool)
	if ok {
		return ok
	}
	s, ok := v.(string)
	if !ok {
		return ok
	}
	return strings.ToLower(s) == "true" || s == "1" || strings.ToLower(s) == "false" || s == "0"
}

type IntMarshaler struct {
	V int
}

func (i *IntMarshaler) Set(v interface{}) {
	i.V = toInt(v)
}

func (i *IntMarshaler) GetValue(k string) interface{} {
	return i.V
}

func (i *IntMarshaler) ValueValid(v interface{}) bool {
	_, ok := v.(int)
	if ok {
		return ok
	}

	s, ok := v.(string)
	if !ok {
		return ok
	}
	_, err := strconv.Atoi(s)
	return err == nil
}

type StringMarshaler struct {
	V string
}

func (s *StringMarshaler) Set(v interface{}) {
	s.V = toString(v)
}

func (s *StringMarshaler) GetValue(k string) interface{} {
	return s.V
}

func (s *StringMarshaler) ValueValid(v interface{}) bool {
	_, ok := v.(string)
	return ok
}
