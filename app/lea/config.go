// 用于更新app.conf配置文件
package lea

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// GenerateHash generates bcrypt hash from plaintext password
func UpdateConfig(name, value string) error {
	res := false
	ConfigPath := "src/github.com/wiselike/leanote2/conf/app.conf"
	file, err := os.OpenFile(ConfigPath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("open config file fail, err: %s", err)
	}

	buf := bufio.NewReader(file)
	output := make([]byte, 0, buf.Size())
	n1 := name + "="
	n2 := name + " "
	for {
		line, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		l := strings.TrimSpace(string(line))
		if strings.HasPrefix(l, n1) || strings.HasPrefix(l, n2) {
			newline := fmt.Sprintf("%s=%s", name, value)
			line = []byte(newline)
			res = true
		}
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
	file.Close()

	if res {
		if err := writeToFile(ConfigPath, output); err != nil {
			return fmt.Errorf("write config file err: %v", err)
		} else {
			return nil
		}
	}
	return fmt.Errorf("no config-name found")
}

func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
