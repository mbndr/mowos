package agent

import (
	"github.com/mbndr/mowos"
)

// an item is responsible for retrieving data from the device
type item interface {
	// set all parameters
	setParams(params map[string]interface{})
	// set a parameter
	setParam(k string, v interface{})
	// get a parameter
	getParam(k string) interface{}
	// returns data and status
	// here you check for service status, partition size etc
	getResponse() mowos.ItemStatus
}

// getItem returns an item depending on the 'type' key
func getItem(itemData map[string]interface{}) item {
	var item item

	// get item by type
	switch itemData["type"] {
	case "service":
		item = &serviceItem{}
	}

	// add params if valid item was found
	if item != nil {
		item.setParams(itemData)
	}

	return item
}
