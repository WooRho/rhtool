package snowflake

import "github.com/bwmarrin/snowflake"

// UniqueID 唯一标识
type UniqueID snowflake.ID

func (v UniqueID) UInt64() uint64 {
	return uint64(v)
}
