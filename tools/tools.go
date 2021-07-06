package tools

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// get client intranet
func GetClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Can not find the client ip address!")

}

// UnicodeToZh unicode转中文
// 举例：http://talent.qh-1.cn/pc/httpclient/orgs?cur_code=440300000000\u0026page=0\u0026pageSize=10 转为 http://talent.qh-1.cn/pc/httpclient/orgs?cur_code=440300000000&page=0&pageSize=10
func UnicodeToZh(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
