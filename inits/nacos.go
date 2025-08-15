package inits

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yuanzhichao-star/public_pkg/config"
)

func InitNaCos() {
	nacosCong := config.AppCong.NaCos
	//初始化客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosCong.NamespaceId, //命名空间id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// 初始化服务端配置
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacosCong.Host,
			ContextPath: "/nacos",
			Port:        uint64(nacosCong.Port),
			Scheme:      "http",
		},
	}
	//创建服务发现客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	//获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosCong.DataId,
		Group:  nacosCong.Group,
	})
	if err != nil {
		panic(err)
	}
	//获取到的配置是json，反序列化
	err = json.Unmarshal([]byte(content), &config.AppCong)
	if err != nil {
		panic(err)
	}
	//热更新
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: nacosCong.DataId,
		Group:  nacosCong.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		panic(err)
	}
}
