package ftwhttp

import (
	"net"
	"net/http"
	"time"
)

// ClientConfig provides configuration options for the HTTP client.
type ClientConfig struct {
	// ConnectTimeout is the timeout for connecting to a server.
	ConnectTimeout time.Duration
	// ReadTimeout is the timeout for reading a response.
	ReadTimeout time.Duration
}

// Client is the top level abstraction in http
type Client struct {
	Transport *Connection
	Jar       http.CookieJar
	config    ClientConfig
}

// Connection is the type used for sending/receiving data
type Connection struct {
	connection  net.Conn
	protocol    string
	readTimeout time.Duration
	duration    *RoundTripTime
}

// RoundTripTime abstracts the time a transaction takes
type RoundTripTime struct {
	begin time.Time
	end   time.Time
}

// FTWConnection is the interface method implement to send and receive data
type FTWConnection interface {
	Request(*Request)
	Response(*Response)
	GetTrackedTime() *RoundTripTime
	send([]byte) (int, error)
	receive() ([]byte, error)
}

// Destination is the host, port and protocol to be used when connecting to a remote host
type Destination struct {
	DestAddr string `default:"localhost"`
	Port     int    `default:"80"`
	Protocol string `default:"http"`
}

// RequestLine is the first line in the HTTP request dialog
type RequestLine struct {
	Method  string `default:"GET"`
	Version string `default:"HTTP/1.1"`
	URI     string `default:"/"`
}

// Request represents a request
// No Defaults represents the previous "stop_magic" behavior
type Request struct {
	requestLine         *RequestLine
	headers             Header
	cookies             http.CookieJar
	data                []byte
	raw                 []byte
	autoCompleteHeaders bool
}

// Response represents the http response received from the server/waf
type Response struct {
	RAW    []byte
	Parsed http.Response
}
