package convert

import (
	"fmt"
	"regexp"
)

type AssertResult struct {
	caller string
	actual string
	checker string
	expect string
}

func assert(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>\s*[A-Za-z0-9._()]+),(?P<checker>\s*\w),(?P<expect>\s*[A-Za-z0-9._()]+)\)`
	r := regexp.MustCompile(p)
	match := r.FindStringSubmatch(line)
	groupNames := r.SubexpNames()
	fmt.Printf("%v, %v \n", match, groupNames)

	return &AssertResult{}, nil
}