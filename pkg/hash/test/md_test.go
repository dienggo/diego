package test

import (
	"github.com/dienggo/diego/pkg/hash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5Hash(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hash test case",
			args: args{
				str: "Blackpink in your area",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hash.MD5Hash(tt.args.str)
			assert.NotEmpty(t, got, "not empty")
		})
	}
}
