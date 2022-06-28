# Pinyincode
`pinyincode` 用于将包含中文的字符串转为拼音码

比如将 `中文` 转为 `zw`，将 `拼音码` 转为 `pym`

# Installation
```
go get -u github.com/koulerz/pinyincode
```

# Warning ⚠️
- 不能够保证多音字的拼音码符合预期。比如 `重庆` 的拼音码可能为 `cq` 或 `zq`

# Usage
```go
package main

import (
	"fmt"

	"github.com/koulerz/pinyincode"
)

func main() {
	// 保留英文字母和数字
	// output: sjnhHelloWorld
	fmt.Println(pinyincode.Code("世界-你好！Hello_World!", pinyincode.FlagWithEnglish|pinyincode.FlagWithNumber))

	// 保留英文字母，数字，连字符和下划线
	// output: sj-nhHello_World
	fmt.Println(pinyincode.Code("世界-你好！Hello_World!", pinyincode.FlagWithEnglish|pinyincode.FlagWithNumber|pinyincode.FlagWithHyphen|pinyincode.FlagWithUnderscore))

	// 仅保留中文拼音
	// output: sjnh
	fmt.Println(pinyincode.Code("世界-你好！Hello_World!", pinyincode.FlagOnlyPinyin))

	// 仅保留中文拼音，FlagOnlyPinyin 会使其他 Flag 失效
	// output: sjnh
	fmt.Println(pinyincode.Code("世界-你好！Hello_World!", pinyincode.FlagOnlyPinyin|pinyincode.FlagWithEnglish|pinyincode.FlagWithNumber))
}
```

# Flags
| Flag               | Description                 |
|:-------------------|:----------------------------|
| FlagWithEnglish    | 保留大小写英文字母                   |
| FlagWithNumber     | 保留数字                        |
| FlagWithHyphen     | 保留连字符                       |
| FlagWithUnderscore | 保留下划线                       |
| FlagOnlyPinyin     | 仅保留中文拼音。该 Flag 会使其他 Flag 失效 |

# Dependencies
- [mozillazg/go-pinyin](https://github.com/mozillazg/go-pinyin)

# License
Under the MIT License.
