package str

import (
	sf "github.com/bwmarrin/snowflake"
	"github.com/nbb2025/distri-domain/pkg/util/logger"
	"go.uber.org/zap"
)

var (
	node *sf.Node
	err  error
)

// 雪花算法工具初始化
// param: 机器节点(用于生成分布式id)
func SnowflakeInit(machineID int64) error {
	node, err = sf.NewNode(machineID)
	if err != nil {
		logger.Error("str.GenerateID err", zap.Error(err))
		return err
	}

	return nil
}

// 生成分布式id
func GenerateID() int64 {
	return node.Generate().Int64()
}

// 生成分布式id string类型
func GenerateIDStr() string {
	return node.Generate().String()
}
