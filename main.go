package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/bitly/go-simplejson"
)

var replaceData string = "<!-- start -->\n"

func getData() (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/Roc-zhou/blog/issues?page=1&per_page=5", nil)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		data, errs := simplejson.NewJson(body)
		if errs != nil {
			return "", errs
		}
		result, _ := data.Array()
		for _, v := range result {
			d := v.(map[string]interface{})
			replaceData += fmt.Sprintf("- [%s](%s)\n", d["title"], d["html_url"])
		}
		replaceData += "<!-- end -->\n"
		return replaceData, nil
	}
	return "", nil
}

func repString(str string) (string, error) {
	re := regexp.MustCompile(`<!-- start -->([\w\W]*)<!-- end -->`)
	match := re.FindStringSubmatch(str)

	if val, err := getData(); err == nil {
		if len(match) != 0 {
			data := strings.Replace(str, match[0], val, -1)
			return data, nil
		}
	} else {
		data := strings.Replace(str, match[0], "获取失败，请手动更新服务！", -1)
		return data, err
	}
	return "", nil
}

func readFile() (string, error) {
	fi, err := os.Open("README.md")
	if err != nil {
		return "", err
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	// fmt.Println(string(chunks))
	return string(chunks), nil
}

func writeFile(str string) {
	filePath := "README.md"
	var err error
	_, err = os.OpenFile(filePath, os.O_TRUNC, 0666) //打开文件
	if err != nil {
		fmt.Println("file open fail", err)
		return
	}
	// 写文件
	err1 := ioutil.WriteFile(filePath, []byte(str), 0666)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("文件更新成功！")
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func main() {
	file, err := readFile()
	if err != nil {
		panic(err)
	}
	fileData, err1 := repString(file)
	if err1 != nil {
		panic(err)
	}
	writeFile(fileData)
}

// /<!-- start -->([\w\W]*)<!-- end -->/
// github-profile-trophy
