package convert

import (
	"fmt"
	"regexp"
	"strings"
)

type AssertResult struct {
	match bool
	caller string
	actual string
	checker string
	expect string
}

func Equals(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*Equals),(?P<expect>.+)\)`
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
				result.match = true
			}
		}
	}

	return result, nil
}

func DeepEquals(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*DeepEquals),(?P<expect>.+)\)`
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
				result.match = true
			}
		}
	}

	return result, nil
}


func IsNil(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*IsNil)\)`
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
				result.match = true
			}
		}
	}

	return result, nil
}

func NotNil(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*NotNil)\)`
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
				result.match = true
			}
		}
	}

	return result, nil
}

func IsTrue(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*IsTrue)\)`
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
				result.match = true
			}
		}
	}

	return result, nil
}

func IsFalse(line string) (*AssertResult, error) {
	return nil, nil
}

func LessEqual(line string) (*AssertResult, error) {
	return nil, nil
}

func Greater(line string) (*AssertResult, error) {
	return nil, nil
}

func GreaterEqual(line string) (*AssertResult, error) {
	return nil, nil
}
