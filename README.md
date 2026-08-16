# go-freqchars

统计一段文本里每个字符出现了多少次，画成纯文本条形图。看密码强度、猜编码、或者纯粹手痒想数数都行。中文按一个字算，不拆偏旁。

可以只取出现最多的前 N 个，也能调"每格代表几个计数"来压缩长尾。

## 装

```bash
go build -o freqchars ./cmd/freqchars
```

## 用

```bash
echo "aabbbc" | ./freqchars
# a ## 2
# b ### 3
# c # 1

echo "aabbbc" | ./freqchars -top 1 -bar =
# b === 3

echo "aaaaab" | ./freqchars -scale 2
# a ## 5
# b # 1
```

参数：
- `-top N`：只显示前 N 个高频字符
- `-bar`：条块字符，默认 `#`
- `-scale`：每格代表的计数，默认 1

## 当库用

```go
import "freqchars"

items := freqchars.Top("aabbbc", 2)     // 按次数降序的前 2 个
freqchars.Bar(items, "#", 1)            // 画成条形文本
```

## License

MIT
