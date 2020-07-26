package main

import (
	. "fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	Println(u.Scheme)
	Println(u.User)
	Println(u.User.Username())
	p, _ := u.User.Password()
	Println(p)

	Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	Println(host)
	Println(port)

	Println(u.Path)
	Println(u.Fragment)

	Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	Println(m)
	Println(m["k"][0])
}
