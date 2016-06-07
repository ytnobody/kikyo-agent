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
    if len(args) < 3 {
        help()
    } else {
        rack := args[1]
        unit := args[2]
        unitid, err := strconv.ParseInt(unit, 10, 32)
        if err != nil {
            help()
        } else {
            err := ag.AddToRack(rack, unitid)
            if err != nil {
                os.Stderr.Write([]byte(err.Error() + "\n"))
                os.Exit(1)
            }
        }
    }
}

func help () {
    fmt.Println(`
Usage:
  kikyo [rackname (string)] [unit-number (integer)]
`)
}
