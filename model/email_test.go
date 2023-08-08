package model_test

import (
	"github.com/matryer/is"
	"github.com/shawkyelshalawy/Daily_Brief/model"
	"testing"
)

func TestEmail_IsValid(t *testing.T) {
	tests := []struct {
		address string
		valid   bool
	}{
		{"me@example.com", true},
		{"@example.com", false},
		{"me@", false},
		{"@", false},
		{"", false},
	}
	t.Run("reports valid email addresses", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.address, func(t *testing.T) {
				is := is.New(t)
				e := model.Email(test.address)
				is.Equal(test.valid, e.IsValid())
			})
		}
	})
}
