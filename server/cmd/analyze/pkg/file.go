package pkg

import (
	"bufio"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
	"net/http"
	"os"
	"summer/server/shared/consts"
	"sync"
	"time"
)

func DownloadFile(url string) ([]string, error) {
	// url comes from minio object
	resp, err := http.Get(url)
	if err != nil {
		klog.Fatal("url error")
	}
	defer resp.Body.Close()

	flag := time.Now().Unix()
	makePath := fmt.Sprintf(consts.NewLocalFilePath, flag)
	tmp, err := os.Create(makePath)
	defer os.Remove(makePath)
	defer tmp.Close()

	_, err = io.Copy(tmp, resp.Body)
	if err != nil {
		return nil, err
	}

	// split file
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(tmp)
	linesCh := make(chan string)
	fileNum := 1
	paths := make([]string, 0)

	makePath = fmt.Sprintf(consts.NewLocalFilePath, fileNum)
	out, err := os.Create(makePath)
	if err != nil {
		return []string{makePath}, nil
	}
	paths = append(paths, makePath)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for line := range linesCh {
			go possessFile(line, out, &wg)
		}
	}()

	for i := 1; scanner.Scan(); {
		linesCh <- scanner.Text()
		i++
		if i >= consts.Lines {
			out.Close()
			fileNum++
			makePath = fmt.Sprintf("./task%d.log", fileNum)
			paths = append(paths, makePath)
			out, err = os.Create(makePath)
			if err != nil {
				klog.Errorf("create task file failed err:" + err.Error())
				return paths, nil
			}
			i = 1
		}
	}

	close(linesCh)
	out.Close()
	wg.Wait()
	return paths, nil
}

func possessFile(line string, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(file, line)
}
