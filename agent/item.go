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
    // get item
    switch itemData["type"] {
    case "service":
        item = &serviceItem{}
    //case "partition":
    //    item = &serviceItem{}
    }
    // return nil if item not valid
    if item == nil {
        return nil
    }
    // fill item
    item.setParams(itemData)
    return item
}
