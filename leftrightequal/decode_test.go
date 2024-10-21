package main

import (
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		encoded string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "LLRR=",
			args: args{
				encoded: "LLRR=",
			},
			want: "210122",
		},
		{
			name: "==RLL",
			args: args{
				encoded: "==RLL",
			},
			want: "000210",
		},
		{
			name: "=LLRR",
			args: args{
				encoded: "=LLRR",
			},
			want: "221012",
		},
		{
			name: "RRL=R",
			args: args{
				encoded: "RRL=R",
			},
			want: "012001",
		},
		{
			name: "LLLLL",
			args: args{
				encoded: "LLLLL",
			},
			want: "543210",
		},
		{
			name: "=LLLL",
			args: args{
				encoded: "=LLLL",
			},
			want: "443210",
		},
		{
			name: "=L=L=L",
			args: args{
				encoded: "=L=L=L",
			},
			want: "3322110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.encoded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
