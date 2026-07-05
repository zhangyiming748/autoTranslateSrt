package parse

import (
	"fmt"
	"os"
	"testing"
)

func TestParseSRTEnglish(t *testing.T) {
	// 解析英文字幕文件
	content, err := os.ReadFile("../Murdercise.2023.1080p.WEBRip.x264.AAC-[YTS.MX]-en.srt")
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	units := ParseSRT(string(content))

	fmt.Printf("========== 英文字幕文件解析结果 ==========\n")
	fmt.Printf("总共解析到 %d 条字幕\n", len(units))
	fmt.Println("-------------------------------------------")

	// 打印前5条字幕作为示例
	maxCount := 5
	if len(units) < maxCount {
		maxCount = len(units)
	}
	for i := 0; i < maxCount; i++ {
		fmt.Printf("[%d] 序号: %s | 时间轴: %s\n", i+1, units[i].No, units[i].Time)
		fmt.Printf("    内容: %s\n", units[i].Content)
		fmt.Println()
	}
}

func TestParseSRTChinese(t *testing.T) {
	// 解析中文字幕文件
	content, err := os.ReadFile("/Users/zen/github/autoTranslateSrt/Murdercise.2023.1080p.WEBRip.x264.AAC-[YTS.MX]-en.srt")
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	units := ParseSRT(string(content))

	fmt.Printf("========== 中文字幕文件解析结果 ==========\n")
	fmt.Printf("总共解析到 %d 条字幕\n", len(units))
	fmt.Println("-------------------------------------------")

	// 打印前5条字幕作为示例
	maxCount := 5
	if len(units) < maxCount {
		maxCount = len(units)
	}
	for i := 0; i < maxCount; i++ {
		fmt.Printf("[%d] 序号: %s | 时间轴: %s\n", i+1, units[i].No, units[i].Time)
		fmt.Printf("    内容: %s\n", units[i].Content)
		fmt.Println()
	}
}
