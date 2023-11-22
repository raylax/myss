package binlog

import (
	"bytes"
	"io"
	"testing"
)

func Test_readerImpl_check(t *testing.T) {
	type fields struct {
		r io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "check positive",
			fields: fields{
				r: bytes.NewReader([]byte("\xfebin")),
			},
			want: true,
		},
		{
			name: "check negative",
			fields: fields{
				r: bytes.NewReader([]byte("\xfebix")),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &readerImpl{
				r: tt.fields.r,
			}
			if got := r.check(); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
