package agent
import (
    "fmt"
    "bytes"
    "time"
    "os"
    "net/http"
    "errors"
    "io/ioutil"
    "github.com/ytnobody/kikyo-agent/node"
    "github.com/ytnobody/kikyo-agent/exporter"
    "github.com/ytnobody/kikyo-agent/state"
)

var VERSION = 0.01
var TimeOut = 15

func Create () (*Agent, error) {
    baseurl := os.Getenv("KIKYO_BASEURL")
    if baseurl == "" {
        err := errors.New("You have to set KIKYO_BASEURL environment value");
        return nil, err
    }
    client := &http.Client{
        Timeout: time.Duration(TimeOut) * time.Second,
    }
    ag := &Agent{
        BaseURL: baseurl,
        Client:  client,
    }
    return ag, nil
}

type Agent struct {
    BaseURL string
    Client  *http.Client
}

func (ag *Agent) AddToRack (rack string, unit int64) error {
    ni, err  := node.Info()
    if err != nil {
        return err
    }
    host := exporter.FromNodeInfo(ni)

    st, err := state.Load()
    if err == nil {
        host.Id = &st.Host.Id
    }

    payload := host.String()

    endpoint := fmt.Sprintf("%s/v1/rack/%s/%d", ag.BaseURL, rack, unit)
    buffer := bytes.NewBuffer([]byte(payload))

    req, err := http.NewRequest("POST", endpoint, buffer)
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", fmt.Sprintf("kikyo.agent/%0.02f", VERSION))

    res, err   := ag.Client.Do(req)
    content, _ := ioutil.ReadAll(res.Body)
    body       := string(content)
    res.Body.Close()

    if err != nil {
        return err
    }

    if res.StatusCode != 200 {
        err = errors.New(fmt.Sprintf("Request failure: %s", res.Status))
        return err
    }

    err = state.Save(body)

    return err
}
