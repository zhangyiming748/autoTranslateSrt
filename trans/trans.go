package trans

// 这里使用translate-shell这个工具实现翻译 指定Google翻译引擎
import (
	"fmt"
	"os/exec"
	"strings"

	"autoTranslateSrt/parse"
)

func Translateshell(srtUnit []parse.SrtUnit) []parse.SrtUnit {
	total := len(srtUnit)
	for i := range srtUnit {
		// 显示进度提示
		if total > 10 && i%10 == 0 {
			fmt.Printf("[进度 %d/%d] (%.1f%%)\n", i, total, float64(i)/float64(total)*100)
		}

		src := strings.TrimSpace(srtUnit[i].Content)
		if src == "" {
			continue
		}

		// 直接修改原切片位置的内容
		srtUnit[i].Content = Translate(src)
	}
	fmt.Printf("[进度 %d/%d] (100.0%%) 翻译完成\n", total, total)
	return srtUnit
}

func Translate(src string) string {
	// trans -b -e google -x 127.0.0.1:8889 -target zh-CN hello
	var (
		cmd  *exec.Cmd
		args []string
	)
	args = append(args, "-b", "-target", "zh-CN", src)
	args = append(args, "-e", "google")
	//args = append(args, "-x", "127.0.0.1:8889")
	args = append(args, "-target", "zh-CN")
	args = append(args, src)
	cmd = exec.Command("trans", args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("翻译失败: %v\n", err)
		return src
	}
	return strings.TrimSpace(string(output))
}
