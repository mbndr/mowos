package agent

import (
    "github.com/mbndr/mowos"
)

// dispatcher holds all items
type dispatcher struct {
    items []item
}

// getValues returns a map of all return values of the items
// TODO pack in goroutines
func (d *dispatcher) getValues() map[string]mowos.ItemStatus {
    values := make(map[string]mowos.ItemStatus)

    for _, i := range d.items {
        values[i.getParam("key").(string)] = i.getResponse()
    }

    return values
}


// logItems logs all registered items
func (d *dispatcher) logItems() {
    for _, i := range d.items {
        mowos.Log.Debugf("reg item '%s' (%s)", i.getParam("name"), i.getParam("type"))
    }
}
