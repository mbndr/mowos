package agent

import (
	"os/exec"

	"github.com/mbndr/mowos"
)

// serviceItem gets data about a service status
type serviceItem struct {
	params map[string]interface{}
}

func (s *serviceItem) setParams(params map[string]interface{}) {
	s.params = params
}

func (s *serviceItem) setParam(k string, v interface{}) {
	s.params[k] = v
}

func (s *serviceItem) getParam(k string) interface{} {
	return s.params[k]
}

// get Response checks weather the service is running or not
func (s *serviceItem) getResponse() mowos.ItemStatus {
	var res mowos.ItemStatus

	cmd := exec.Command("service", s.getParam("service").(string), "status")
	err := cmd.Run()

	// service not found or other ExitError
	if err != nil {
		res.Status = mowos.CRIT
		res.Value = "Service is not running"
		return res
	}

	res.Status = mowos.OK
	res.Value = "Service is running"

	return res
}
