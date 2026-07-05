package main

import (
	"fmt"
	"os"
	"path/filepath"

	"autoTranslateSrt/parse"
	"autoTranslateSrt/print"
	"autoTranslateSrt/trans"
)

func main() {
	// 获取当前目录下的 srt 文件夹中所有 .srt 文件的绝对路径
	srtDir := "./srt"
	files, err := filepath.Glob(filepath.Join(srtDir, "*.srt"))
	if err != nil {
		fmt.Printf("读取 srt 文件夹失败: %v\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("srt 文件夹中没有找到 .srt 文件")
		os.Exit(1)
	}

	fmt.Printf("找到 %d 个字幕文件\n\n", len(files))

	// 循环处理每个字幕文件
	for _, file := range files {
		fmt.Printf("========== 处理文件: %s ==========\n", file)
		processSRT(file)
		fmt.Println()
	}

	fmt.Println("所有字幕处理完成！")
}

func processSRT(inputPath string) {
	// 1. 解析原字幕
	content, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	units := parse.ParseSRT(string(content))
	fmt.Printf("解析到 %d 条字幕\n", len(units))

	if len(units) == 0 {
		fmt.Println("警告: 字幕文件为空，跳过")
		return
	}

	// 2. 翻译为中文
	translatedUnits := trans.Translateshell(units)
	fmt.Println("翻译完成")

	// 3. 保存为新字幕
	// 生成输出文件名（添加 -zh-CN 后缀）
	outputPath := generateOutputPath(inputPath)
	err = print.SaveSRT(translatedUnits, outputPath)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		return
	}
}

// generateOutputPath 根据输入文件路径生成输出文件路径（添加 -zh-CN 后缀）
func generateOutputPath(inputPath string) string {
	dir := filepath.Dir(inputPath)
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]

	// 在扩展名前插入 -zh-CN
	newName := name + "-zh-CN" + ext
	return filepath.Join(dir, newName)
}
