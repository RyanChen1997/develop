package main

import (
	"errors"
	"strings"
	"unicode"
)

type Args struct {
	marshaler map[string]Marshaler
}

func NewArgs(schema string, args []string) (*Args, error) {
	a := new(Args)
	a.marshaler = make(map[string]Marshaler)

	if err := a.parseSchema(schema); err != nil {
		return nil, err
	}
	if err := a.parseArgs(args); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Args) parseSchema(schema string) error {
	for _, elem := range strings.Split(schema, ",") {
		if err := a.parseElement(elem); err != nil {
			return err
		}
	}
	return nil
}

func (a *Args) parseElement(elem string) error {
	if len(elem) == 0 {
		return nil
	}

	elemRune := []rune(elem)
	elemId := elemRune[0]
	if !unicode.IsLower(elemId) {
		return errors.New("INVALID ELEMENT ID: element id must be lower-case")
	}
	// if len(elemRune) == 1 {
	// 	a.marshaler[string(elemId)] = new(BoolMarshaler)
	// 	return nil
	// }

	elemTail := elemRune[1:]
	marshaler := parseMarshalerType(string(elemTail))
	if marshaler == nil {
		return errors.New("INVALID ELEMENT TAIL: element tail must be one of '', '#', '*'")
	}
	a.marshaler[string(elemId)] = marshaler
	return nil
}

func (a *Args) parseArgs(args []string) error {
	idx := 0
	for idx < len(args) {
		argKey := a.parseArgKey(strings.Trim(args[idx], " "))
		idx++
		if argKey == "" {
			continue
		}
		if idx == len(args) {
			break
		}
		marshaler, found := a.marshaler[argKey]
		if !found {
			return errors.New("INVALID ARG KEY")
		}
		argVal := strings.Trim(args[idx], " ")
		valid := marshaler.ValueValid(argVal)
		if !valid {
			return errors.New("INVALID ARG VALUE")
		}
		marshaler.Set(argVal)
		idx++
	}
	return nil
}

func (a *Args) parseArgKey(elem string) string {
	if len(elem) == 0 {
		return ""
	}
	elemRune := []rune(elem)
	if elemRune[0] != '-' {
		return ""
	}
	return string(elemRune[1:])
}

func parseMarshalerType(elem string) Marshaler {
	switch elem {
	case "":
		return new(BoolMarshaler)
	case "#":
		return new(IntMarshaler)
	case "*":
		return new(StringMarshaler)
	default:
		return nil
	}
}

func (a *Args) GetBool(k string) bool {
	marshaler, found := a.marshaler[k]
	if !found {
		return false
	}
	m, ok := marshaler.(*BoolMarshaler)
	if !ok {
		return false
	}
	return m.GetValue(k).(bool)
}

func (a *Args) GetString(k string) string {
	marshaler, found := a.marshaler[k]
	if !found {
		return ""
	}
	m, ok := marshaler.(*StringMarshaler)
	if !ok {
		return ""
	}
	return m.GetValue(k).(string)
}

func (a *Args) GetInt(k string) int {
	marshaler, found := a.marshaler[k]
	if !found {
		return 0
	}
	m, ok := marshaler.(*IntMarshaler)
	if !ok {
		return 0
	}
	return m.GetValue(k).(int)
}

type Marshaler interface {
	Set(v interface{})
	GetValue(k string) interface{}
	ValueValid(v interface{}) bool
}
