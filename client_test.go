package swagger

import (
	"context"
	"os"
	"testing"
)

var (
	userName string
	password string
	basePath string
	ctx      context.Context
	client   *APIClient
)

func TestMain(m *testing.M) {
	userName = os.Getenv("HARBOR_USERNAME")
	password = os.Getenv("HARBOR_PASSWORD")
	basePath = os.Getenv("HARBOR_BASEPATH")

	//fmt.Printf("HARBOR_USERNAME=%s HARBOR_PASSWORD=%s HARBOR_BASEPATH=%s\n", userName, password, basePath)

	ctx = context.Background()
	ctx = context.WithValue(ctx, ContextBasicAuth, BasicAuth{
		UserName: userName,
		Password: password,
	})
	cfg := NewConfiguration()
	cfg.BasePath = basePath
	client = NewAPIClient(cfg)
	m.Run()
}

func TestPingApiService_GetPing(t *testing.T) {
	data, rsp, err := client.PingApi.GetPing(ctx, nil)
	if err != nil {
		t.Fatalf("PingApi.GetPing error:%+v", err)
	}

	t.Logf("PingApi.GetPing data:%+v", data)
	t.Logf("PingApi.GetPing rsp:%+v", rsp)
}

func TestStatisticApi(t *testing.T) {
	req := new(StatisticApiGetStatisticOpts)
	data, rsp, err := client.StatisticApi.GetStatistic(ctx, req)
	if err != nil {
		t.Fatalf("get statistic error:%+v", err)
	}

	t.Logf("get statistic data:%+v", data)
	t.Logf("get statistic rsp:%+v", rsp)
}

func TestRegistryApi(t *testing.T) {
	var req RegistryPing
	rsp, err := client.RegistryApi.PingRegistry(ctx, req, nil)
	if err != nil {
		t.Fatalf("RegistryAPI.PingRegistry error:%+v", err)
	}

	t.Logf("RegistryAPI.PingRegistry rsp:%+v", rsp)
}
