package main

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type StateDecision struct {
	TreeUUID         string   `json:"tree_uuid,omitempty"`
	UseAddOnServices []string `json:"use_add_on_services,omitempty"`
	EvaluationType   string   `json:"evaluation_type,omitempty"`
}

type UIFlowSetting struct {
	UIVersion           string         `json:"ui_version,omitempty"`
	CanSkip             bool           `json:"can_skip,omitempty"`
	IgnoreConflictField bool           `json:"ignore_conflict_field,omitempty"`
	UINavigation        []string       `json:"ui_navigation,omitempty"`
	CanMoveNextState    bool           `json:"can_move_next_state,omitempty"`
	AlwaysShow          bool           `json:"always_show,omitempty"`
	StateDecision       *StateDecision `json:"state_decision,omitempty"`
}

func main() {

}
