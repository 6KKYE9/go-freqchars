package freqchars

import (
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	Char  string
	Count int
}

// Count 统计每个字符出现次数，中文按一个字算。返回按字符排序的列表。
func Count(s string) []Item {
	m := map[string]int{}
	for _, r := range s {
		m[string(r)]++
	}
	out := make([]Item, 0, len(m))
	for k, v := range m {
		out = append(out, Item{Char: k, Count: v})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Char < out[j].Char })
	return out
}

// Top 取出现最多的前 n 个（n<=0 表示全部），按次数降序；次数相同按字符升序。
func Top(s string, n int) []Item {
	all := Count(s)
	sort.Slice(all, func(i, j int) bool {
		if all[i].Count != all[j].Count {
			return all[i].Count > all[j].Count
		}
		return all[i].Char < all[j].Char
	})
	if n > 0 && n < len(all) {
		all = all[:n]
	}
	return all
}

// Bar 把统计结果画成文本条形图，bar 是单个条块字符，scale 是"每格代表几个计数"。
func Bar(items []Item, bar string, scale int) string {
	if scale <= 0 {
		scale = 1
	}
	var b strings.Builder
	for _, it := range items {
		cells := it.Count / scale
		if it.Count%scale != 0 {
			cells++
		}
		b.WriteString(it.Char)
		b.WriteString(" ")
		b.WriteString(strings.Repeat(bar, cells))
		b.WriteString(" ")
		b.WriteString(strconv.Itoa(it.Count))
		b.WriteString("\n")
	}
	return b.String()
}
