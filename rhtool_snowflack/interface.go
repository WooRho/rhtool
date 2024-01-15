package snowflake

// NodeIface SnowFlakeNode
type NodeIface interface {
	// Init 初始化
	Init()
	// GenerateID 生成UniqueID
	GenerateID() UniqueID
}

// NewCustomNode 自定义
// nodeNum use const best
// dm := NewCustomNode()
// 实际开发中需要init它 作为全局服务实例
// In actual development, it needs to be init as a global service instance
func NewCustomNode(nodeNum ...int64) NodeIface {
	var num int64
	if len(nodeNum) > 0 {
		num = nodeNum[0]
	}
	node := newCustomNode(num)
	node.Init()
	return node
}
