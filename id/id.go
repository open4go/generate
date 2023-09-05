package id

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

// GenID 生成一个新的id
// 可以添加前缀
func GenID(prefix string) string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "-"
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return fmt.Sprintf("%s%d", prefix, id)
}
