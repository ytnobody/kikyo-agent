# github.com/ytnobody/kikyo-agent

Agent for [Kikyo](https://github.com/ytnobody/Kikyo)

## SYNOPSIS

    ## Add your host information to kikyo, then rack to 20th unit in "myrack-0102"
    KIKYO_BASEURL=http://your.kikyo.hostname:5000 kikyo-agent myrack-0102 20 0

## USAGE

    kikyo [rackname (string)] [unit-number (integer)] [virtual-id (integer)]
    
    rackname:    rackname that contains this host
    unit-number: unit-number that racked this host
    virtual-id:  If this host is real node, specify 0. Otherwise, specify 1 or over.

## AUTHOR

[ytnobody](https://github.com/ytnobody)
