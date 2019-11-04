package yaml

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

func TestYamlEncode(t *testing.T) {
	data := map[string]interface{}{
		"test": `etcdEnv: "dev"
# 开启debug会有更多打印
debug: true
# 本地调用算法地址，该模式下不会使用etcd
swaggerStatic: html
# etcd配置
etcdAddr:
#  - "http://localhost:2379"
  - "http://etcd-0.etcd.dev.svc.cluster.local:2379"
# 项目api_serv的配置
apiServerConfig:
  apiVersion: "/alpha"
  listenAddr: ":8080"
  debug: true
  title: "algo_proxy"
  description: "算法代理"
  # swaggerMgrAddr: "192.168.0.91:10120"
  # swaggerProjectId: "placmol9osalxlkr"
  jaegerAgent: "192.168.0.15:6831"
  # swaggerFile: "doc/openapi.yaml"
  swaggerServerMap:
      prod: "http://192.168.0.42:8080"`,
	}
	var writer bytes.Buffer
	err := yaml.NewEncoder(&writer).Encode(data)
	require.NoError(t, err)
	t.Log(writer.String())
}

func TestYamlEncodeMultiline(t *testing.T) {
	data := []string{
		"val1",
		`sonar-scanner \
-Dsonar.projectKey=code-review-demo \
-Dsonar.sources=. \
-Dsonar.host.url=http://192.168.0.3:9000 \
-Dsonar.login=3494646fc9b746274ea52f3df53c6d5d3791709a`,
	}
	fmt.Println(data)
	var writer bytes.Buffer
	err := yaml.NewEncoder(&writer).Encode(data)
	require.NoError(t, err)
	t.Log(writer.String())
}

func TestDefaultType(t *testing.T) {
	data := `
replicaCount: 2
`

	receiver := map[string]interface{}{}
	err := yaml.NewDecoder(strings.NewReader(data)).Decode(&receiver)
	require.NoError(t, err)
	t.Logf("%T %v", receiver["replicaCount"], receiver["replicaCount"])
}
