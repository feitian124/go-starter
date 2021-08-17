package mathx

import "testing"

func TestRoundInt(t *testing.T) {
	type args struct {
		i   int64
		dec int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{1, 0}, 1},
		{"2", args{2146213728964879326, -15}, 2146000000000000000},
		{"3", args{123456789, -5}, 123500000},
		{"4", args{50, -2}, 100},
		{"5", args{150, 2}, 150},
		{"-1", args{-1, 0}, -1},
		{"-2", args{-2146213728964879326, -15}, -2146000000000000000},
		{"-3", args{-123456789, -5}, -123500000},
		{"-4", args{-50, -2}, -100},
		{"-5", args{-150, 2}, -150},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundInt(tt.args.i, tt.args.dec); got != tt.want {
				t.Errorf("RoundInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
