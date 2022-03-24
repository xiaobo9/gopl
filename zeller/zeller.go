// 蔡勒公式
// 罗马教皇决定在1582年10月4日后使用格利戈里历法
// 1582年10月4日后：w = (d + 1+ 2*m+3*(m+1)/5+y+y/4-y/100+y/400)%7;
// 1582年10月4日前：w = (d+1+2*m+3*(m+1)/5+y+y/4+5) % 7;

package zeller

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var wMap = map[int]string{
	0: "周日",
	1: "周一",
	2: "周二",
	3: "周三",
	4: "周四",
	5: "周五",
	6: "周六",
}

func Zeller(day string) (int, error) {
	want := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !want.MatchString(day) {
		fmt.Println("error ", day)
		return 0, errors.New("参数错误, " + day)
	}
	strs := strings.Split(day, "-")
	if len(strs) != 3 {
		fmt.Println("error ", day)
		return 0, errors.New("参数错误, " + day)
	}
	y, _ := strconv.Atoi(strs[0])
	m, _ := strconv.Atoi(strs[1])
	d, _ := strconv.Atoi(strs[2])
	var w int
	if y < 1582 || (y == 1582 && m < 10) || (y == 1582 && m == 10 && d <= 4) {
		w = (d + 1 + 2*m + 3*(m+1)/5 + y + y/4 + 5) % 7
	} else if y == 1582 && m == 10 && (d >= 5 && d <= 14) {
		fmt.Println(day)
		return 0, errors.New("参数错误, " + day)
	} else {
		w = (d + 1 + 2*m + 3*(m+1)/5 + y + y/4 - y/100 + y/400) % 7
	}

	fmt.Println(day, " ", w, ": ", wMap[w])
	return w, nil
}
