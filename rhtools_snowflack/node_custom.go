package snowflake

import (
	"github.com/WooRho/rhtool/rhtool_common"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var _ NodeIface = &customNode{}

const (
	defaultNodeBits uint8 = 4 // 节点数，最大 10 位(1-10)，可以有 1024 个节点
	dufaultStepBits uint8 = 8 // 计数序列码，最大 12 位(1-12)，每毫秒产生 4096 个 ID
)

type customNode struct {
	node     *snowflake.Node
	nodeOnce sync.Once
	nodeErr  error
	nodeNum  int64
}

func newCustomNode(nodeNum int64) *customNode {
	return &customNode{
		nodeNum: nodeNum,
	}
}

func (n *customNode) Init() {
	n.nodeOnce.Do(func() {
		snowflake.NodeBits = defaultNodeBits // 4位 最多16个节点
		snowflake.StepBits = dufaultStepBits // 8位 每毫秒产生 255 个 ID（为了兼容js）

		//节点
		if n.nodeNum == 0 {
			n.nodeNum = 1
		}
		n.node, n.nodeErr = snowflake.NewNode(n.nodeNum)
		if n.nodeErr != nil {
			panic(n.nodeErr)
		}
	})
}

func (n *customNode) GenerateID() UniqueID {
	if n.node == nil {
		n.Init()
	}

	id := n.node.Generate()
	n.nodeOnce.Do(func() {
		rhtool_common.Persistence("snow", id)
	})

	return UniqueID(id)
}
