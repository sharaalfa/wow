package handler

import (
	"bytes"
	"errors"
	"net"
	"time"
)

type MockConn struct {
	ReadBuffer  *bytes.Buffer
	WriteBuffer *bytes.Buffer
	closed      bool
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	if m.closed {
		return 0, errors.New("connection closed")
	}
	return m.ReadBuffer.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	if m.closed {
		return 0, errors.New("connection closed")
	}
	return m.WriteBuffer.Write(b)
}

func (m *MockConn) Close() error {
	m.closed = true
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return nil
}

func (m *MockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
