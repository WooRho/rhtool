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
func NewCustomNode(nodeNum int64) NodeIface {
	node := newCustomNode(nodeNum)
	node.Init()
	return node
}
