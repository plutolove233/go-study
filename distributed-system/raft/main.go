package main

import (
	"distributed-system/raft/models"
	"distributed-system/raft/raftrpc"
	"log"

	"github.com/spf13/viper"
)

type ServerCfg struct {
	Index   int    `mapstructure:"index"`
	Address string `mapstructure:"address"`
}

type Config struct {
	Me       ServerCfg `mapstructure:"server"`
	Clusters []string  `mapstructure:"clusters"`
}

var node = Config{}

func init() {
	vconfig := viper.New()
	vconfig.AddConfigPath("./config")
	vconfig.SetConfigName("cfg")
	vconfig.SetConfigType("yaml")
	if err := vconfig.ReadInConfig(); err != nil {
		log.Fatalln("read config file failed, error=", err.Error())
	}
	vconfig.Unmarshal(&node)
}

func main() {
	clusters := make([]*raftrpc.ClientEnd, 0)
	// record other peer's information
	for _, address := range node.Clusters {
		clusters = append(clusters, raftrpc.NewClientEnd(address))
	}

	raft := models.NewRaft(node.Me.Index, clusters)
	raft.RunRPCService(node.Me.Address)
	raft.Start()
	select {}
}
