package exporter

import (
    "fmt"
    "regexp"
    "strings"
    "encoding/json"
    "github.com/ytnobody/kikyo-agent/node"
)

func FromNodeInfo(n *node.NodeInfo) Host {
    ips := []string{}

    ipaddr_format, _ := regexp.Compile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
    loopbackip_format, _ := regexp.Compile(`^127\.`)

    for _, ni := range n.Net {
        for _, item := range ni.Addrs {
            ip := []byte(item.Addr)
            if ipaddr_format.Match(ip) && ! loopbackip_format.Match(ip) {
                ips = append(ips, item.Addr)
            }
        }
    }

    myhost := Host{
        Name:     n.Host.Hostname,
        CPU:      fmt.Sprintf("%s x %dCore", n.CPU[0].ModelName, n.CPU[0].Cores),
        Memory:   fmt.Sprintf("%dGB", n.Memory.Total / 1024 / 1024 / 1024),
        Disk:     fmt.Sprintf("%dGB", n.Disk.Total / 1024 / 1024 / 1024),
        OS:       fmt.Sprintf("%s %s", n.Host.Platform, n.Host.PlatformVersion),
        IP:       strings.Join(ips, ", "),
    }

    return myhost
}

type Host struct {
    Id       *int64   `json:"id"`
    Name     string   `json:"name"`
    CPU      string   `json:"cpu"`
    Memory   string   `json:"memory"`
    Disk     string   `json:"disk"`
    OS       string   `json:"os"`
    IP       string   `json:"ip"`
}

func (d Host) String() string {
    s, _ := json.Marshal(d)
    return string(s)
}
