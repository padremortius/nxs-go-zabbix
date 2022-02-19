package zabbix

// For `ItemObject` field: `type`
const (
	ItemTypeZabbixAgent       = 0
	ItemTypeZabbixTrapper     = 2
	ItemTypeSimpleCheck       = 3
	ItemTypeZabbixInternal    = 5
	ItemTypeZabbixAgentActive = 7
	ItemTypeWebItem           = 9
	ItemTypeExternalCheck     = 10
	ItemTypeDatabaseMonitor   = 11
	ItemTypeIPMIAgent         = 12
	ItemTypeSSHAgent          = 13
	ItemTypeTelnetAgent       = 14
	ItemTypeCalculated        = 15
	ItemTypeJMXAgent          = 16
	ItemTypeSNMPTrap          = 17
	ItemTypeDependentItem     = 18
	ItemTypeHTTPAgent         = 19
	ItemTypeSNMPAgent         = 20
	ItemTypeScript            = 21
)

// For `ItemObject` field: `value_type`
const (
	ItemValueTypeNumFloat    = 0
	ItemValueTypeCharacter   = 1
	ItemValueTypeLog         = 2
	ItemValueTypeNumUnsigned = 3
	ItemValueTypeText        = 4
)

// For `ItemObject` field: `flags`
const (
	ItemFlagsPlainItem      = 0
	ItemFlagsDiscoveredItem = 4
)

// For `ItemObject` field: `state`
const (
	ItemStateNormal       = 0
	ItemStateNotSupported = 1
)

// For `ItemObject` field: `status`
const (
	ItemStatusEnabled  = 0
	ItemStatusDisabled = 1
)

// ItemObject struct is used to store host operations results
//
// see: https://www.zabbix.com/documentation/5.4/en/manual/api/reference/item/object#item
type ItemObject struct {
	ItemId      string `json:"itemid,omitempty"`
	Delay       string `json:"delay,omitempty"`
	HostId      string `json:"hostid,omitempty"`
	InterfaceId string `json:"interfaceid,omitempty"`
	Key         string `json:"key_,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	ValueType   int    `json:"value_type,omitempty"` // has defined consts, see above
	Descr       string `json:"description,omitempty"`
	Error       string `json:"error,omitempty"`
	Flags       int    `json:"flags,omitempty"` // has defined consts, see above
	LastClock   int    `json:"lastclock,omitempty"`
	LastValue   string `json:"lastvalue,omitempty"`
	State       int    `json:"state,omitempty"`  // has defined consts, see above
	Status      int    `json:"status,omitempty"` // has defined consts, see above
}

// ItemTagObject struct is used to store item tag
//
// see: https://www.zabbix.com/documentation/5.4/en/manual/api/reference/item/object#item-tag
type ItemTagObject struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// ItemGetParams struct is used for item get requests
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/item/get#parameters
type ItemGetParams struct {
	GetParameters

	ItemIDs      []string        `json:"itemids,omitempty"`
	GroupIDs     []string        `json:"groupids,omitempty"`
	TemplateIDs  []string        `json:"templateids,omitempty"`
	HostIDs      []string        `json:"hostids,omitempty"`
	ProxyIDs     []string        `json:"proxyids,omitempty"`
	InterfaceIDs []string        `json:"interfaceids,omitempty"`
	GraphIDs     []string        `json:"graphids,omitempty"`
	TriggerIDs   []string        `json:"triggerids,omitempty"`
	WebItems     int             `json:"webitems,omitempty"`
	Inherited    bool            `json:"inherited,omitempty"`
	Templated    bool            `json:"templated,omitempty"`
	Monitored    bool            `json:"monitored,omitempty"`
	Group        string          `json:"group,omitempty"`
	Host         string          `json:"host,omitempty"`
	EvalType     int             `json:"evaltype,omitempty"`
	WithTriggers bool            `json:"with_triggers,omitempty"`
	Tags         []ItemTagObject `json:"tags,omitempty"`
	LimitSelect  int             `json:"limitSelect,omitempty"`

	SelectHosts SelectQuery `json:"selectHosts,omitempty"`
	// SelectInterfaces    SelectQuery `json:"selectInterfaces,omitempty"`    // not implemented yet
	// SelectTriggers      SelectQuery `json:"selectTriggers,omitempty"`      // not implemented yet
	// SelectGraphs        SelectQuery `json:"selectGraphs,omitempty"`        // not implemented yet
	// SelectDiscoveryRule SelectQuery `json:"selectDiscoveryRule,omitempty"` // not implemented yet
	SelectItemDiscovery SelectQuery `json:"selectItemDiscovery,omitempty"`
	// SelectPreprocessing SelectQuery `json:"selectPreprocessing,omitempty"` // not implemented yet
	SelectTags SelectQuery `json:"selectTags,omitempty"`
	// SelectValueMap      SelectQuery `json:"selectValueMap,omitempty"`      // not implemented yet
}

// ItemGet gets hosts
func (z *Context) ItemGet(params ItemGetParams) ([]ItemObject, int, error) {

	var result []ItemObject

	status, err := z.request("item.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
