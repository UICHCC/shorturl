package util

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func InitSnowflake() {
	node, _ = snowflake.NewNode(1)
}

func NextUrl() string {
	id := node.Generate()
	return id.Base58()
}
