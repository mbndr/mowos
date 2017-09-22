package agent

import (
	"encoding/json"
	"sync"

	"github.com/mbndr/mowos"
)

// dispatcher holds all items
type dispatcher struct {
	items []item
}

// get the response of each item the dispatcher holds
func (d *dispatcher) getItemResponses() mowos.AgentResponse {
	values := make(map[string]mowos.ItemStatus)

	var wg sync.WaitGroup

	for _, i := range d.items {
		wg.Add(1)
		// start goroutine to retrieve the response
		go func(i item) {
			defer wg.Done()
			values[i.getParam("key").(string)] = i.getResponse()
		}(i)
	}

	// wait for all goroutines to finish
	wg.Wait()

	return values
}

// get all item responses and convert it to json byte array
func (d *dispatcher) getItemResponsesBytes() ([]byte, error) {
	resp := d.getItemResponses()

	bytes, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// logItems logs all registered items
func (d *dispatcher) logItems() {
	for _, i := range d.items {
		mowos.Log.Debugf("reg item '%s' (%s)", i.getParam("name"), i.getParam("type"))
	}
}
