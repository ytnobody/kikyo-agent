package node

import (
    "encoding/json"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/net"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/disk"
)

// Get host information and return as NodeInfo struct
func Info() (*NodeInfo, error) {
    nil_info := &NodeInfo{}

    host_info, err := host.Info()
    if err != nil {
        return nil_info, err
    }

    nic_info, err := net.Interfaces()
    if err != nil {
        return nil_info, err
    }

    cpu_info, err := cpu.Info()
    if err != nil {
        return nil_info, err
    }

    mem_info, err := mem.VirtualMemory()
    if err != nil {
        return nil_info, err
    }

    disk_info, err := disk.Usage("/")
    if err != nil {
        return nil_info, err
    }

    node_info := &NodeInfo{
        Host: *host_info,
        Net: nic_info,
        CPU: cpu_info,
        Memory: *mem_info,
        Disk: *disk_info,
    }

    return node_info, nil
}

type NodeInfo struct {
    Host   host.InfoStat         `json:"host"`
    Net    []net.InterfaceStat   `json:"net"`
    CPU    []cpu.InfoStat        `json:"cpu"`
    Memory mem.VirtualMemoryStat `json:"memory"`
    Disk   disk.UsageStat        `json:"disk"`
}

func (d NodeInfo) String() string {
    s, _ := json.Marshal(d)
    return string(s)
}

