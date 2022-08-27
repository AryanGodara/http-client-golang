package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/AryanGodara/http-client-golang/gohttp_mock"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package 'examples'")

	gohttp_mock.MockupServer.Start()
	defer gohttp_mock.MockupServer.Stop()

	os.Exit(m.Run())
}
