package print

import (
	"fmt"
	"os"

	"autoTranslateSrt/parse"
)

// SaveSRT 将翻译后的字幕块切片保存为新的字幕文件
func SaveSRT(units []parse.SrtUnit, outputPath string) error {
	if len(units) == 0 {
		return fmt.Errorf("没有可保存的字幕数据")
	}

	var content string
	for _, unit := range units {
		content += fmt.Sprintf("%s\n%s\n%s\n\n", unit.No, unit.Time, unit.Content)
	}

	// 添加文件末尾换行符
	content += "\n"

	// 写入文件
	err := os.WriteFile(outputPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	fmt.Printf("成功保存字幕文件: %s\n", outputPath)
	fmt.Printf("共保存 %d 条字幕\n", len(units))
	return nil
}
