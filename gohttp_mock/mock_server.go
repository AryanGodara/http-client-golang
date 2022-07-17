package gohttp_mock

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/AryanGodara/http-client-golang/core"
)

var (
	MockupServer = mockServer{
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

func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = true
}

func (m *mockServer) Stop() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = false
}

func (m *mockServer) DeleteMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.mocks = make(map[string]*Mock) // Remove prev map, create a fresh one
}

func (m *mockServer) AddMock(mock Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	key := m.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	m.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()

	key := method + url + MockupServer.cleanBody(body)
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

func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

func (m *mockServer) GetMockedClient() core.HttpClient {
	return m.httpClient
}
