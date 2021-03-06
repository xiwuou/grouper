package tool

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ---------------------------------------------
// 文件遍历
// ---------------------------------------------

// GetFileList 获取文件列表
func GetFileList(path string) (newPathList []string) {
	// 增加后缀
	if path[len(path)-1:] != "/" {
		path += "/"
	}
	// 开始遍历
	iterateOverFiles(path, func(newPath string) {
		newPathList = append(newPathList, newPath)
	})
	// 等待遍历结束
RE:
	length := len(newPathList)
	time.Sleep(2 * time.Second)
	for len(newPathList) != length {
		goto RE
	}
	return newPathList
}

// iterateOverFiles 遍历指定路径的文件
func iterateOverFiles(path string, up func(newPath string)) {
	// 获取路径
	fs, _ := ioutil.ReadDir(path)
	// 执行遍历
	for _, file := range fs {
		if file.IsDir() {
			// 遇到文件夹时就开启一个并发递归
			// go iterateOverFiles(path+file.Name()+"/", up)
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				iterateOverFiles(path+file.Name()+"/", up)
				wg.Done()
			}()
			wg.Wait()
		} else {
			newPath := path + file.Name()
			// fmt.Println("扫描: ", path[:len(path)-1])
			up(newPath) // 调用函数参数
		}
	}
}

// ---------------------------------------------
// 退出程序
// ---------------------------------------------

func GoodBye() {
	fmt.Println("⚠ 按回车或回复任意，退出程序。")
	reader := bufio.NewReader(os.Stdin) // 读取命令行
	osWin := IsOsWindows()              // 当前系统是否为windows
	_ = readInput(reader, osWin)
	// 按 CTRL+C 或输入 exit 以退出程序
	// t := strings.Split(text, " ")
	// if len(t) == 1 && strings.Compare("exit", text) == 0 {
	// 	fmt.Sprintln("Bye~ Bye~")
	// 	os.Exit(1)
	// }
	if true {
		fmt.Sprintln("Bye~ Bye~")
		os.Exit(1)
	}
}

// IsOsWindows 获取当前计算机系统类型，是否为Windows
func IsOsWindows() bool {
	// runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	sysType := runtime.GOOS
	// fmt.Println(fmt.Sprintf("您的系统是%v，采用%v架构", runtime.GOOS, runtime.GOARCH))
	switch sysType {
	case "linux":
		return false // LINUX系统
	case "windows":
		return true // windows系统
	case "darwin":
		return false // LINUX系统
	default:
		return false // 其他系统
	}
}

// 读取用户输入
func readInput(reader *bufio.Reader, osWin bool) (text string) {
	if osWin {
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
	} else {
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
	}
	return text
}
