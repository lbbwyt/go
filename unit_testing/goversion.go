// goversion
package main

import (
	"unit_testing/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
