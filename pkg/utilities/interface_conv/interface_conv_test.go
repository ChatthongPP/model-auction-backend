package interface_conv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToUint(t *testing.T) {
	type args struct {
		iParem interface{}
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "string to int",
			args: args{
				iParem: "1",
			},
			want: uint64(1),
		},
		{
			name: "float32 to int",
			args: args{
				iParem: float32(2.34),
			},
			want: uint64(2),
		},
		{
			name: "float64 to int",
			args: args{
				iParem: float64(7.34),
			},
			want: uint64(7),
		},
		{
			name: "int to int",
			args: args{
				iParem: int(3),
			},
			want: uint64(3),
		},
		{
			name: "int32 to int",
			args: args{
				iParem: int32(4),
			},
			want: uint64(4),
		},
		{
			name: "int64 to int",
			args: args{
				iParem: int64(5),
			},
			want: uint64(5),
		},
		{
			name: "can't convert",
			args: args{
				iParem: "One",
			},
			want: uint64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToUint(tt.args.iParem)
			xType := reflect.TypeOf(tt.args.iParem)
			fmt.Println("Type:", xType, " Parem:", tt.args.iParem, " Want:", tt.want, " Got:", got)
			if got != tt.want {
				t.Errorf("ToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	type args struct {
		fParem interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "int to float64",
			args: args{
				fParem: 1,
			},
			want: float64(1.00),
		},
		{
			name: "int32 to float64",
			args: args{
				fParem: int32(1),
			},
			want: float64(1.00),
		},
		{
			name: "int64 to float64",
			args: args{
				fParem: int64(1),
			},
			want: float64(1.00),
		},
		{
			name: "string to float64",
			args: args{
				fParem: "2.34",
			},
			want: float64(2.34),
		},
		{
			name: "float32 to float64",
			args: args{
				fParem: float32(2.34),
			},
			want: float64(2.34),
		},
		{
			name: "float64 to float64",
			args: args{
				fParem: float64(2.34),
			},
			want: float64(2.34),
		},
		{
			name: "can't convert",
			args: args{
				fParem: "One",
			},
			want: float64(0.00),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToFloat(tt.args.fParem)
			xType := reflect.TypeOf(tt.args.fParem)
			fmt.Println("Type:", xType, " Parem:", tt.args.fParem, " Want:", tt.want, " Got:", got)
			if got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
