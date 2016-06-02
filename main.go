package main

import (
    "fmt"
    "github.com/ytnobody/niji/node"
    "github.com/ytnobody/niji/exporter/yabitz"
)

func main() {
    node_info, _ := node.Info()
    yabitz_info := yabitz.FromNodeInfo(node_info)
    fmt.Println(yabitz_info)
}

