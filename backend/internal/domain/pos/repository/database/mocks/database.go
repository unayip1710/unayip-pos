package mock

import "github.com/unayip1710/unayip-pos/utils/mocker"

type mockContainer struct {
	mocker mocker.MockDataStore
}

func NewMockDatabase() *mockContainer {
	return &mockContainer{
		mocker: mocker.New(),
	}
}
