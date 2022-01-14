package db

import "fmt"

type Proxy struct {
	/* 0 is all (http, https)
	 * 1 is http
	 * 2 is https
	 */
	protocol  byte
	ip        string
	port      uint16
	addr      string
	country   string
	location  string
	latency   uint16
	anonymous bool
	validated bool
	createdAt int64
	updatedAt int64
}

func NewProxy(ip string, port uint16) *Proxy {
	return &Proxy{
		ip:   ip,
		port: port,
		addr: fmt.Sprintf("%v:%d", ip, port),
	}
}

func (proxy *Proxy) Protocol(protocol byte) *Proxy {
	proxy.protocol = protocol
	return proxy
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
