package cmd

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestRemoveCRLF(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Empty slice",
			args: args{data: []byte{}},
			want: []byte{},
		},
		{
			name: "One byte slice - no crlf - no lf",
			args: args{data: []byte{'A'}},
			want: []byte{'A'},
		},
		{
			name: "One byte slice - with crlf",
			args: args{data: []byte{'A', '\r', '\n'}},
			want: []byte{'A'},
		},
		{
			name: "One byte slice - with lf",
			args: args{data: []byte{'A', '\n'}},
			want: []byte{'A'},
		},
		{
			name: "Multi bytes slice - no crlf - no lf",
			args: args{data: []byte{'A', 'B', 'C'}},
			want: []byte{'A', 'B', 'C'},
		},
		{
			name: "Multi bytes slice - with crlf",
			args: args{data: []byte{'A', 'B', 'C', '\r', '\n'}},
			want: []byte{'A', 'B', 'C'},
		},
		{
			name: "Multi bytes slice - with lf",
			args: args{data: []byte{'A', 'B', 'C', '\n'}},
			want: []byte{'A', 'B', 'C'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCRLF(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeCRLF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		in        io.Reader
		direction Direction
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name:    "Empty input - dos2unix",
			args:    args{in: new(bytes.Buffer), direction: dos2unix},
			wantOut: "",
			wantErr: false,
		},
		{
			name:    "Empty input - unix2dos",
			args:    args{in: new(bytes.Buffer), direction: unix2dos},
			wantOut: "",
			wantErr: false,
		},
		{
			name:    "Non empty input - unsupported direction",
			args:    args{in: bytes.NewBuffer([]byte("ABC\n")), direction: "unsupported"},
			wantOut: "",
			wantErr: true,
		},
		{
			name:    "Non empty input - dos2unix - dos format",
			args:    args{in: bytes.NewBuffer([]byte("1\r\nABC\r\n\r\nDEF\r\n")), direction: dos2unix},
			wantOut: "1\nABC\n\nDEF\n",
			wantErr: false,
		},
		{
			name:    "Non empty input - dos2unix - unix format",
			args:    args{in: bytes.NewBuffer([]byte("1\nABC\n\nDEF\n")), direction: dos2unix},
			wantOut: "1\nABC\n\nDEF\n",
			wantErr: false,
		},
		{
			name:    "Non empty input - dos2unix - No EOL",
			args:    args{in: bytes.NewBuffer([]byte("ABC")), direction: dos2unix},
			wantOut: "",
			wantErr: false,
		},
		{
			name:    "Non empty input - unix2dos - unix format",
			args:    args{in: bytes.NewBuffer([]byte("1\nABC\n\nDEF\n")), direction: unix2dos},
			wantOut: "1\r\nABC\r\n\r\nDEF\r\n",
			wantErr: false,
		},
		{
			name:    "Non empty input - unix2dos - dos format",
			args:    args{in: bytes.NewBuffer([]byte("1\r\nABC\r\n\r\nDEF\r\n")), direction: unix2dos},
			wantOut: "1\r\nABC\r\n\r\nDEF\r\n",
			wantErr: false,
		},
		{
			name:    "Non empty input - unix2dos - No EOL",
			args:    args{in: bytes.NewBuffer([]byte("ABC")), direction: unix2dos},
			wantOut: "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := convert(tt.args.in, out, tt.args.direction); (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("convert() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
