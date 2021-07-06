package dnacos

import (
	"dragon/core/dragon/conf"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

// https://github.com/nacos-group/nacos-sdk-go
var NamingClient naming_client.INamingClient

func Init() {
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./log/nacos/log",
		CacheDir:            "./log/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: viper.GetString("nacos.ip"),
			Port:   viper.GetUint64("nacos.port"),
		},
	}

	// 创建服务发现客户端
	var err error
	NamingClient, err = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Println("创建服务发现客户端失败", err)
	}

	// 创建动态配置客户端
	//configClient, err := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})

	// Register instance：RegisterInstance
	success, err := NamingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          conf.IntranetIp,
		Port:        viper.GetUint64("server.port"),
		ServiceName: viper.GetString("server.servicename"),
		ClusterName: viper.GetString("nacos.clustername"),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		GroupName:   viper.GetString("nacos.groupname"),
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": viper.GetString("nacos.idc")},
	})
	if !success {
		log.Fatalln("nacos服务注册失败", err)
	}
	log.Println("nacos服务注册成功：", conf.IntranetIp+":", viper.GetUint64("server.port"))
}

func DeregisterInstance() {
	NamingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          conf.IntranetIp,
		Port:        viper.GetUint64("server.port"),
		Cluster:     viper.GetString("nacos.clustername"),
		ServiceName: viper.GetString("server.servicename"),
		GroupName:   viper.GetString("nacos.groupname"),
		Ephemeral:   true,
	})
}

// SelectOneHealthyInstance
func SelectOneHealthyInstance(serviceName string, groupName string, clusterNames []string) (instanceAddr string, instance *model.Instance, err error) {
	instance, err = NamingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   groupName,    // default value is DEFAULT_GROUP
		Clusters:    clusterNames, // default value is DEFAULT
	})
	if instance == nil || err != nil {
		return
	}

	port := strconv.FormatInt(int64(instance.Port), 10)
	instanceAddr = instance.Ip + ":" + port
	return
}
