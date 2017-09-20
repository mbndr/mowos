package mowos

type itemStatusCode uint8

const (
    OK itemStatusCode = iota
    WARN
    CRIT
)

// ItemStatus is returned from each item and sent to the monitor
type ItemStatus struct {
    Status itemStatusCode `json:"status"`
    Value interface{} `json:"value"`
}
