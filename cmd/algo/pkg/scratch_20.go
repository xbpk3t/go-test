package main

import (
	"fmt"
	"net/url"
)

func main() {
	type args struct {
		isHTTPS bool
		host    string
		path    string
		values  [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{true, "www.tab.com", "", [][]string{{"x", "111"}, {"y", "222"}}}, "https://www.tab.com?x=111&y=222"},
		{"", args{false, "www.tab.com", "", [][]string{{"x", "111"}, {"y", "222"}}}, "http://www.tab.com?x=111&y=222"},
	}
	for _, tt := range tests {
		if got := GenURL(tt.args.isHTTPS, tt.args.host, tt.args.path, tt.args.values); got != tt.want {
			// t.Errorf("GenURL() = %v, want %v", got, tt.want)
			fmt.Errorf("GenURL() = %v, want %v", got, tt.want)
		}
	}
}

func GenURL(isHTTPS bool, host, path string, values [][]string) string {
	ssl := ""
	if isHTTPS {
		ssl = "https"
	} else {
		ssl = "http"
	}
	res := &url.URL{
		Scheme: ssl,
		Host:   host,
		Path:   path,
	}
	query := url.Values{}
	for _, value := range values {
		query.Add(value[0], value[1])
	}
	res.RawQuery = query.Encode()
	return res.String()
}
