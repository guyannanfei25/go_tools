package tools

import (
    "net"
    "net/http"
    "time"
)

func NewTimeOutHttpClient(dailTimeout, readTimeout, writeTimeout time.Duration) *http.Client {
    // TODO: Check valid
    transport := &http.Transport{
        Dial: func(network, addr string) (net.Conn, error) {
            c, err := net.DialTimeout(network, addr, dailTimeout)
            if err != nil {
                return nil, err
            }

            return &timeoutConn{readTimeout, writeTimeout, c}, nil
        },
    }

    return &http.Client{Transport: transport}
}

type timeoutConn struct {
    readTimeout      time.Duration
    writeTimeout     time.Duration

    net.Conn
}

func (c *timeoutConn) Read(b []byte) (n int, err error) {
    c.Conn.SetReadDeadline(time.Now().Add(c.readTimeout))
    return c.Conn.Read(b)
}

func (c *timeoutConn) Write(b []byte) (n int, err error) {
    c.Conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
    return c.Conn.Write(b)
}
