package mocker

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sync"
)

type (
	mockStore struct {
		mu        sync.Mutex
		dataStore map[hash][]mockData
	}

	MockDataStore interface {
		Patch(name string, data ...interface{})
		Get(name string) mockData
	}

	inputToHash struct {
		Name string `json:"name"`
	}

	mockData []interface{}

	hash [16]byte
)

func New() MockDataStore {
	return &mockStore{
		dataStore: map[hash][]mockData{},
	}
}

func (mock *mockStore) Patch(name string, data ...interface{}) {
	mock.mu.Lock()
	defer mock.mu.Unlock()

	input := inputToHash{
		Name: name,
	}
	inputHash := toHash(input)

	output := mockData{}
	output = append(output, data...)

	mock.dataStore[inputHash] = append(mock.dataStore[inputHash], output)
}

func (mock *mockStore) Get(name string) mockData {
	mock.mu.Lock()
	defer mock.mu.Unlock()

	input := inputToHash{
		Name: name,
	}

	inputHash := toHash(input)

	arrOutput, exists := mock.dataStore[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic(fmt.Sprintf("mock not available for \"%s\"", name))
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.dataStore[inputHash] = arrOutput

	return output
}

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}
