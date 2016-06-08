package state

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "errors"
)

var StateFile = "/tmp/kikyo.state"

// Save agent state
func Save (jsonStr string) error {
    content := []byte(fmt.Sprintf("%s", jsonStr))
    err := ioutil.WriteFile(StateFile, content, os.ModePerm)

    return err
}

// Load agent state
func Load () (*State, error) {
    if NotExists() {
        return nil, errors.New("StateFile is not exists")
    }

    content, err := ioutil.ReadFile(StateFile)
    if err != nil {
        return nil, err
    }

    var st State
    err = json.Unmarshal(content, &st)
    if err != nil {
        return nil, err
    }

    return &st, nil
}


// Check StateFile is not exists
func NotExists () bool {
    _, err := os.Stat(StateFile)
    return os.IsNotExist(err)
}

type Host struct {
    Id        int64  `json:"id"`
    Name      string `json:"name"`
    Rack      string `json:"rack"`
    Unit      int64  `json:"unit"`
    VirtualID int64  `json:"virtual_id"`
    Size      int64  `json:"size"`
    CPU       string `json:"cpu"`
    Memory    string `json:"memory"`
    Disk      string `json:"disk"`
    OS        string `json:"os"`
    HWID      string `json:"hwid"`
    ModelName string `json:"modelname"`
    Ip        string `json:"ip"`
    Note      string `json:"node"`
    Status    string `json:"status"`
    CreateAt  string `json:"create_at"`
    UpdateAt  string `json:"update_at"`
}

type State struct {
    Host *Host `json:"host"`
}
