package gohttp_mock

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/AryanGodara/http-client-golang/core"
)

var (
	mockupServer = mockServer{
		mocks:       make(map[string]*Mock),
		enabled:     false,
		serverMutex: sync.Mutex{},
		httpClient:  &httpClientMock{},
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	httpClient core.HttpClient

	mocks map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = false
}

func DeleteMocks() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.mocks = make(map[string]*Mock) // Remove prev map, create a fresh one
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()

	key := method + url + m.cleanBody(body)
	hasher.Write([]byte(key))

	key = hex.EncodeToString(hasher.Sum(nil))

	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)

	if body == "" {
		return ""
	}

	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")

	return body
}

func IsMockServerEnabled() bool {
	return mockupServer.enabled
}

func GetMockedClient() core.HttpClient {
	return mockupServer.httpClient
}
