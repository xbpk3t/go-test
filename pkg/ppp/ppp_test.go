package ppp

import (
	"reflect"
	"testing"
)

func TestGetDiscountInfo(t *testing.T) {
	type args struct {
		countryCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				countryCode: "CN",
			},
			want: &Result{
				CountryCode: "CN",
				CountryName: "中国",
				Discount:    0.9,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDiscountInfo(tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDiscountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiscountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
