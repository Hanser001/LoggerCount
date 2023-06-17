package mapreduce

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Master struct {
	mu sync.Mutex

	middleData chan []*KV
	files      []string
}

type KV struct {
	Key   string
	value string
	Count int64
}

func NewMaster(files []string) *Master {
	return &Master{mu: sync.Mutex{}, middleData: make(chan []*KV, len(files)), files: files}
}

func (m *Master) Map(filed string) ([]*KV, error) {
	if len(m.files) == 0 {
		return nil, errors.New("no data")
	}

	for _, v := range m.files {
		go doMap(v, filed, m.middleData)
	}

	middleData := make([]*KV, 0, len(m.files))

	for i := 0; i < len(m.files); i++ {
		select {
		case kvs := <-m.middleData:
			middleData = append(middleData, kvs...)
		}
	}

	return middleData, nil
}

func doMap(filePath, filed string, middleData chan<- []*KV) {
	file, err := os.Open(filePath)
	if err != nil {
		middleData <- nil
		return
	}
	defer file.Close()

	Info, err := file.Stat()
	if err != nil {
		middleData <- nil
		return
	}

	outPut := make([]*KV, 0, Info.Size())

	scanner := bufio.NewScanner(file)
	data := make(map[string]interface{})
	for scanner.Scan() {
		err = json.Unmarshal([]byte(scanner.Text()), &data)
		if err != nil {
			continue
		}
		outPut = append(outPut, &KV{Key: data[filed].(string), value: ""})
	}

	middleData <- outPut
}

func (m *Master) Reduce(groups map[string][]string) []*KV {
	var wg sync.WaitGroup
	result := make([]*KV, 0, len(groups))

	for key, value := range groups {
		wg.Add(1)
		go func(key string, value []string) {
			defer wg.Done()
			kv := doReduce(key, value)
			m.mu.Lock()
			result = append(result, kv)
			m.mu.Unlock()
		}(key, value)
	}
	wg.Wait()

	return result
}

func doReduce(key string, v []string) *KV {
	return &KV{Key: key, Count: int64(len(v))}
}

func (m *Master) Generalize(middleData []*KV) map[string][]string {
	groups := make(map[string][]string)

	for _, v := range middleData {
		groups[v.Key] = append(groups[v.Key], v.value)
	}

	return groups
}
