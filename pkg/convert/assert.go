package convert

import (
	"fmt"
	"regexp"
	"strings"
)

type AssertResult struct {
	caller string
	actual string
	checker string
	expect string
}

func assert(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>\s*[A-Za-z0-9._()]+),(?P<checker>\s*\w+),(?P<expect>\s*[A-Za-z0-9._()]+)\)`
	r := regexp.MustCompile(p)
	match := r.FindStringSubmatch(line)
	groupNames := r.SubexpNames()
	fmt.Printf("%v, %v, %d, %d\n", match, groupNames, len(match), len(groupNames))
	result := &AssertResult{}
	if len(match) == len(groupNames) {
		// 转换为map
		for i, name := range groupNames {
			if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
				if name == "caller" {
					result.caller = strings.TrimSpace(match[i])
				}
				if name == "actual" {
					result.actual = strings.TrimSpace(match[i])
				}
				if name == "checker" {
					result.checker = strings.TrimSpace(match[i])
				}
				if name == "expect" {
					result.expect = strings.TrimSpace(match[i])
				}
			}
		}
	}

	return result, nil
}