package errors

import (
	"encoding/json"
	"fmt"
)

type ErrorType = string

const (
	INTERNAL    ErrorType = "internal"
	APPLICATION           = "application"
)

type Field struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type Error struct {
	Type    ErrorType `json:"type"`
	Code    string    `json:"code"`
	Path    string    `json:"path"`
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Cause   error     `json:"cause"`
	Context map[string]interface{}
	Fields  []Field
}

func New(t ErrorType, c string) *Error {
	return &Error{
		Type:    t,
		Code:    c,
		Context: make(map[string]interface{}),
		Fields:  make([]Field, 0),
	}
}

func NewApplication(c string) *Error {
	return New(APPLICATION, c)
}

func (e *Error) SetPath(path string) *Error {
	e.Path = path
	return e
}

func (e *Error) SetStatus(status int) *Error {
	e.Status = status
	return e
}

func (e *Error) SetMessage(msg string, args ...interface{}) *Error {
	e.Message = fmt.Sprintf(msg, args...)
	return e
}

func (e *Error) SetCause(cause error) *Error {
	e.Cause = cause
	return e
}

func (e *Error) SetContext(context map[string]interface{}) *Error {
	for k, v := range context {
		e.Context[k] = v
	}
	return e
}

func (e *Error) AddContext(k string, v interface{}) *Error {
	e.Context[k] = v
	return e
}

func (e *Error) AddFields(fields []Field) *Error {
	for _, v := range fields {
		e.Fields = append(e.Fields, v)
	}
	return e
}

func (e *Error) AddField(f string, v interface{}) *Error {
	e.Fields = append(e.Fields, Field{f, v})
	return e
}

func (e *Error) Error() string {
	str, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(str)
}
