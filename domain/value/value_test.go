package value_test

import (
	"testing"

	"github.com/tabakazu/go-webapp/domain/value"
)

func TestUsernameOrEmail_IsEmail(t *testing.T) {
	tests := []struct {
		name string
		v    value.UsernameOrEmail
		want bool
	}{
		{"It is expected to receive true with the email argument", value.UsernameOrEmail("foo@bar.baz"), true},
		{"It is expected to receive true with the username argument", value.UsernameOrEmail("foo_bar_baz"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.v.IsEmail(); got != tt.want {
				t.Errorf("UsernameOrEmail.IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
