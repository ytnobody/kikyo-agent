package yabitz

import (
    "fmt"
    "regexp"
    "encoding/json"
    "github.com/ytnobody/niji/node"
)

func FromNodeInfo(n *node.NodeInfo) Host {
    node_type := "VMware(Guest)"
    if n.Host.VirtualizationSystem == "" {
        node_type = "real"
    }

    localip  := []string{}
    globalip := []string{}

    ipaddr_format, _ := regexp.Compile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
    localip_format, _ := regexp.Compile(`^(192|172|10)\.`)
    loopbackip_format, _ := regexp.Compile(`^127\.`)

    for _, ni := range n.Net {
        for _, item := range ni.Addrs {
            ip := []byte(item.Addr)
            if ipaddr_format.Match(ip) && ! loopbackip_format.Match(ip) {
                if localip_format.Match([]byte(item.Addr)) {
                    localip = append(localip, item.Addr)
                } else {
                    globalip = append(globalip, item.Addr)
                }
            }
        }
    }

    myhost := Host{
        Type:     node_type,
        Service:  51,
        Status:   "IN_SERVICE",
        CPU:      fmt.Sprintf("%s x %dCore", n.CPU[0].ModelName, n.CPU[0].Cores),
        Memory:   fmt.Sprintf("%dGB", n.Memory.Total / 1024 / 1024 / 1024),
        Disk:     fmt.Sprintf("%dGB", n.Disk.Total / 1024 / 1024 / 1024),
        OS:       fmt.Sprintf("%s %s", n.Host.OS, n.Host.PlatformVersion),
        DNSName:  []string{ n.Host.Hostname },
        LocalIP:  localip,
        GlobalIP: globalip,
    }

    return myhost
}

type Host struct {
    Type     string   `json:"type"`
    Service  int      `json:"service"`
    Status   string   `json:"status"`
    RackUnit string   `json:"rackunit"`
    HWID     string   `json:"hwid"`
    HWInfo   string   `json:"hwinfo"`
    CPU      string   `json:"cpu"`
    Memory   string   `json:"memory"`
    Disk     string   `json:"disk"`
    OS       string   `json:"os"`
    DNSName  []string `json:"dnsnames"`
    LocalIP  []string `json:"localips"`
    GlobalIP []string `json:"globalips"`
}

func (d Host) String() string {
    s, _ := json.Marshal(d)
    return string(s)
}
