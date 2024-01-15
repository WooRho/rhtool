package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowflack(t *testing.T) {
	dm := NewCustomNode()
	fmt.Println(dm.GenerateID().UInt64())
	fmt.Println(dm.GenerateID().UInt64())
	fmt.Println(dm.GenerateID().Int64())
	fmt.Println(dm.GenerateID().Int64())
	//1705969424593152
	//1705969424593153
	//1705969424593154
	//1705969424593155
}
