package errors

import (
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		name     string
		create   func() error
		expected error
	}{{
		"default",
		func() error {
			return New(INTERNAL, "CODE")
		},
		&Error{
			Type:    INTERNAL,
			Code:    "CODE",
			Context: make(map[string]interface{}),
		},
	}, {
		"set status",
		func() error {
			return New(APPLICATION, "CODE").SetStatus(401)
		},
		&Error{
			Type:    APPLICATION,
			Code:    "CODE",
			Status:  401,
			Context: make(map[string]interface{}),
		},
	}, {
		"set message",
		func() error {
			return New(APPLICATION, "CODE").SetMessage("%d;%.2f;%s", 1, 1.2, "Hi")
		},
		&Error{
			Type:    APPLICATION,
			Code:    "CODE",
			Message: "1;1.20;Hi",
			Context: make(map[string]interface{}),
		},
	}, {
		"context and fields",
		func() error {
			return New(APPLICATION, "CODE").
				SetContext(map[string]interface{}{
					"Prop1": "Value1",
					"Prop2": 2,
				}).
				AddFields([]Field{{"field1", "value1"}}).
				SetContext(map[string]interface{}{
					"Prop3": []int{5, 6},
				}).
				AddFields([]Field{{"field2", "value2"}})
		},
		&Error{
			Type: APPLICATION,
			Code: "CODE",
			Context: map[string]interface{}{
				"Prop1": "Value1",
				"Prop2": 2,
				"Prop3": []int{5, 6},
			},
			Fields: []Field{
				{"field1", "value1"},
				{"field2", "value2"},
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.create()
			if !reflect.DeepEqual(err, test.expected) {
				t.Errorf("%s\n-e: %v\n-a: %v", test.name, test.expected, err)
			}
		})
	}
}
