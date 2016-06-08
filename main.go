package main

import (
    "fmt"
    "os"
    "strconv"
    "github.com/ytnobody/kikyo-agent/agent"
)

func main () {
    ag, err := agent.Create()
    if err != nil {
        os.Stderr.Write([]byte(err.Error() + "\n"))
        os.Exit(1)
    }

    args := os.Args
    if len(args) < 4 {
        help()
    } else {
        rack := args[1]
        unit := args[2]
        virtualID := args[3]

        unitid, err := strconv.ParseInt(unit, 10, 32)
        if err != nil {
            help()
        } 

        vid, err := strconv.ParseInt(virtualID, 10, 32)
        if err != nil {
            help()
        }

        err = ag.AddToRack(rack, unitid, vid)
        if err != nil {
            os.Stderr.Write([]byte(err.Error() + "\n"))
            os.Exit(1)
        }
    }
}

func help () {
    fmt.Println(`
Usage:
  kikyo [rackname (string)] [unit-number (integer)] [virtual-id (integer)]

    rackname:    rackname that contains this host
    unit-number: unit-number that racked this host
    virtual-id:  If this host is real node, specify 0. Otherwise, specify 1 or over.
`)
    os.Exit(0)
}
