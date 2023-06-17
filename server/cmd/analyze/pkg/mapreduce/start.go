package mapreduce

import (
	"errors"
)

func StartWordCount(files []string, filed string) ([]*KV, error) {
	if len(files) == 0 {
		return nil, errors.New("no data")
	}
	wc := NewMaster(files)
	middleData, err := wc.Map(filed)
	if err != nil {
		return nil, err
	}

	groups := wc.Generalize(middleData)

	result := wc.Reduce(groups)

	return result, nil
}
