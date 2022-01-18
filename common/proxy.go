package common

import "fmt"

type Proxy struct {
	// http, https, none
	schema    string
	ip        string
	port      uint16
	Addr      string
	country   string
	location  string
	latency   uint16
	anonymous bool
	Validated bool
	CreatedAt int64
	UpdatedAt int64
}

func NewProxy(schema string, ip string, port uint16) *Proxy {
	return &Proxy{
		schema: schema,
		ip:     ip,
		port:   port,
		Addr:   fmt.Sprintf("%v://%v:%d", schema, ip, port),
	}
}

func (proxy *Proxy) Country(country string) *Proxy {
	proxy.country = country
	return proxy
}

func (proxy *Proxy) Location(location string) *Proxy {
	proxy.location = location
	return proxy
}

func (proxy *Proxy) Latency(latency uint16) *Proxy {
	proxy.latency = latency
	return proxy
}

func (proxy *Proxy) Anonymous(anonymous bool) *Proxy {
	proxy.anonymous = anonymous
	return proxy
}
