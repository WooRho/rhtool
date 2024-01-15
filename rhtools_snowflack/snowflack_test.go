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
}
