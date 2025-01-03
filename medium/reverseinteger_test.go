package medium

import "testing"

func Test_reverse(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 123",
			args: args{123},
			want: 321,
		},
		{
			name: "case -123",
			args: args{-123},
			want: -321,
		},
		{
			name: "case 120",
			args: args{120},
			want: 21,
		},
		{
			name: "case -2147483412",
			args: args{-2147483412},
			want: -2143847412,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
