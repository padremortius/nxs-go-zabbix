package zabbix

// For `TriggerObject` field: `flags`
const (
	TriggerFlagsPlain      = 0
	TriggerFlagsDiscovered = 4
)

// For `TriggerObject` field: `priority`
const (
	TriggerPriorityNotClassified = 0
	TriggerPriorityInformation   = 1
	TriggerPriorityWarinig       = 2
	TriggerPriorityAverage       = 3
	TriggerPriorityHigh          = 4
	TriggerPriorityDisaster      = 5
)

// For `TriggerObject` field: `state`
const (
	TriggerStateIsUpToDate = 0
	TriggerStateUnknown    = 1
)

// For `TriggerObject` field: `type`
const (
	TriggerTypeGenSingleEvent    = 0
	TriggerTypeGenMultipleEvents = 1
)

// For `TriggerObject` field: `value`
const (
	TriggerValueOk      = 0
	TriggerValueProblem = 1
)

// For `TriggerObject` field: `recovery_mode`
const (
	TriggerRMExpression         = 0 //default value
	TriggerRMRecoveryExpression = 1
	TriggerRMNone               = 2
)

// For `TriggerObject` field: `correlation_mode`
const (
	TriggerCMAllProblems = 0 //default value
	TriggerCMIfTagMatch  = 1
)

// For `TriggerObject` field: `manual_close`
const (
	TriggerManualCloseNo  = 0 //default value
	TriggerManualCloseYes = 1
)

// TriggerObject struct is used to store host operations results
//
// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/trigger/object#trigger
type TriggerObject struct {
	TriggerId          string `json:"triggerid,omitempty"`
	Description        string `json:"description,omitempty"`
	Expression         string `json:"expression,omitempty"`
	Opdata             string `json:"opdata,omitempty"`
	Comments           string `json:"comment,omitempty"`
	Error              string `json:"error,omitempty"`
	Flags              int    `json:"flags,omitempty"`
	LastChange         int    `json:"lastchange,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	State              int    `json:"state,omitempty"`
	Status             int    `json:"status,omitempty"`
	TemplateID         string `json:"templateid,omitempty"`
	Type               int    `json:"type,omitempty"`
	Url                string `json:"url,omitempty"`
	Value              int    `json:"value,omitempty"`
	RecoveryMode       int    `json:"recovery_mode,omitempty"`
	RecoveryExpression string `json:"recovery_expression,omitempty"`
	CorrelationMode    int    `json:"correlation_mode,omitempty"`
	CorrelationTag     string `json:"correlation_tag,omitempty"`
	ManualClose        int    `json:"manual_close,omitempty"`
	Discover           int    `json:"discover,omitempty"`
	Uuid               string `json:"uuid,omitempty"`
}

// TriggerTagObject struct is used to store item tag
//
// see: https://www.zabbix.com/documentation/5.4/en/manual/api/reference/trigger/object#trigger-tag
type TriggerTagObject struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// TriggerGetParams struct is used for item get requests
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/trigger/get#parameters
type TriggerGetParams struct {
	GetParameters

	TriggerIDs     []string `json:"triggerids,omitempty"`
	GroupIDs       []string `json:"groupids,omitempty"`
	TemplateIDs    []string `json:"templateids,omitempty"`
	HostIDs        []string `json:"hostids,omitempty"`
	ItemIDs        []string `json:"itemids,omitempty"`
	ApplicationIDs []string `json:"applicationids,omitempty"`
	Functions      []string `json:"functions,omitempty"`

	Group                       string             `json:"group,omitempty"`
	Host                        string             `json:"host,omitempty"`
	Inherited                   bool               `json:"inherited,omitempty"`
	Templated                   bool               `json:"templated,omitempty"`
	Dependent                   bool               `json:"dependent,omitempty"`
	Monitored                   int                `json:"monitored,omitempty"`
	Active                      int                `json:"active,omitempty"`
	Maintenance                 bool               `json:"maintenance,omitempty"`
	WithUnacknowledgedEvents    int                `json:"withUnacknowledgedEvents,omitempty"`
	WithAcknowledgedEvents      int                `json:"withAcknowledgedEvents,omitempty"`
	WithLastEventUnacknowledged int                `json:"withLastEventUnacknowledged,omitempty"`
	SkipDependent               int                `json:"skipDependent,omitempty"`
	LastChangeSince             int                `json:"lastChangeSince,omitempty"`
	LastChangeTill              int                `json:"lastChangeTill,omitempty"`
	OnlyTrue                    int                `json:"only_true,omitempty"`
	MinSeverity                 int                `json:"min_severity,omitempty"`
	EvalType                    int                `json:"evaltype,omitempty"`
	Tags                        []TriggerTagObject `json:"tags,omitempty"`
	ExpandComment               int                `json:"expandComment,omitempty"`
	ExpandDescription           int                `json:"expandDescription,omitempty"`
	ExpandExpression            int                `json:"expandExpression,omitempty"`
	LimitSelect                 int                `json:"limitSelect,omitempty"`

	SelectGroups SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts  SelectQuery `json:"selectHosts,omitempty"`
	SelectItems  SelectQuery `json:"selectItems,omitempty"`
	// SelectFunctions SelectQuery `json:"selectFunctions,omitempty"` // not implemented yet
}

// TriggerGet gets hosts
func (z *Context) TriggerGet(params TriggerGetParams) ([]TriggerObject, int, error) {

	var result []TriggerObject

	status, err := z.request("trigger.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
