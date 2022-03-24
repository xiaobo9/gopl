// ch3-05
// https://books.studygolang.com/gopl-zh/ch3/ch3-05.html
package comma

// comma inserts commas in a non-negative decimal integer string.
// “12345”处理后成为“12,345”
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

// 练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

// 练习 3.12： 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
