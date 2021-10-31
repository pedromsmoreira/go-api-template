package logs_test

import (
	"errors"
	"go-api-template/internal/logs"
	"testing"
)

func TestLogIf(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"no error",
			args{
				err: nil,
			},
		},
		{
			"error",
			args{
				err: errors.New("this is error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logs.LogErrorIf(tt.args.err != nil, "test %v failed: %w", tt.name, tt.args.err)
		})
	}
}
