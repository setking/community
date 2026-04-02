package utils

import (
	"errors"
	"fmt"
	"net"
)

// 获取可用端口
func GetFreePort() (int, error) {
	// ":0" 让系统自动选择可用端口，兼容 IPv4 和 IPv6
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, fmt.Errorf("get free port: %w", err)
	}
	defer l.Close()

	addr, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		return 0, errors.New("get free port: unexpected address type")
	}
	return addr.Port, nil
}
