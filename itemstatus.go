package mowos

type itemStatusCode string

const (
    OK itemStatusCode = "ok"
    WARN itemStatusCode = "warning"
    CRIT itemStatusCode = "critical"
)

// ItemStatus is returned from each item and sent to the monitor
type ItemStatus struct {
    Status itemStatusCode `json:"status"`
    Value interface{} `json:"value"`
}
