package parse

import (
	"regexp"
	"strings"
)

type SrtUnit struct {
	No      string // 序号
	Time    string // 时间轴
	Content string // 内容
}

// ParseSRT 解析SRT字幕文件内容，返回字幕块切片
// 支持单行和多行字幕，兼容 \n 和 \r\n 换行符
func ParseSRT(content string) []SrtUnit {
	// 将 \r\n 替换为 \n，统一换行符格式
	content = strings.ReplaceAll(content, "\r\n", "\n")

	// 按空行分割字幕块（兼容 \n 和 \r\n）
	blockRegex := regexp.MustCompile(`\d+\n\d{2}:\d{2}:\d{2},\d{3} --> \d{2}:\d{2}:\d{2},\d{3}`)

	// 找到所有字幕块的起始位置
	locations := blockRegex.FindAllStringIndex(content, -1)
	if len(locations) == 0 {
		return nil
	}

	var units []SrtUnit

	for i, loc := range locations {
		// 提取当前字幕块的文本
		start := loc[0]
		var end int
		if i+1 < len(locations) {
			end = locations[i+1][0]
		} else {
			end = len(content)
		}

		blockText := content[start:end]
		lines := strings.Split(blockText, "\n")

		if len(lines) < 2 {
			continue
		}

		// 解析序号
		no := strings.TrimSpace(lines[0])

		// 解析时间轴
		time := strings.TrimSpace(lines[1])

		// 解析内容（从第3行开始，可能是多行）
		var contentLines []string
		for j := 2; j < len(lines); j++ {
			line := strings.TrimSpace(lines[j])
			if line != "" {
				contentLines = append(contentLines, line)
			}
		}

		content := strings.Join(contentLines, "\n")

		units = append(units, SrtUnit{
			No:      no,
			Time:    time,
			Content: content,
		})
	}

	return units
}
