package tests

type LenderConfig struct {
	ID              int                       `json:"id"`
	Name            string                    `json:"name"`
	Active          bool                      `json:"active"`
	Tags            []Tag                     `json:"tags"`
	UIFlow          []string                  `json:"ui_flow"`
	UIFlowSettings  map[string]UIFlowSetting  `json:"ui_flow_settings,omitempty"`
	UIVersion       string                    `json:"ui_version"`
	Weight          int                       `json:"weight"`
	LenderID        int                       `json:"lender_id"`
	LeadExpiry      int                       `json:"lead_expiry"`
	LeadOfferExpiry int                       `json:"lead_offer_expiry"`
	Domain          string                    `json:"domain"`
	AuthSettings    AuthSettings              `json:"auth_settings"`
	DecisionEngines map[string]DecisionEngine `json:"decision_engines"`
	AddOnServices   map[string]AddOnService   `json:"add_on_services"`
	EKYCSettings    EKYCSettings              `json:"ekyc_settings"`
	OTPService      OTPService                `json:"otp_service"`
	DocumentSchema  DocumentSchema            `json:"document_schema"`
}

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

type StateDecision struct {
	TreeUUID         string   `json:"tree_uuid,omitempty"`
	UseAddOnServices []string `json:"use_add_on_services,omitempty"`
	EvaluationType   string   `json:"evaluation_type,omitempty"`
}

type AuthSettings struct {
	Callback        CallbackSettings `json:"callback"`
	DefaultAuthType string           `json:"default_auth_type"`
	FaceID          FaceIDSettings   `json:"face_id"`
}

type CallbackSettings struct {
	AuthRequired bool `json:"auth_required"`
}

type FaceIDSettings struct {
	Enabled bool `json:"enabled"`
}

type DecisionEngine struct {
	EvaluationType    string   `json:"evaluation_type,omitempty"`
	MaxWaitSeconds    int      `json:"max_wait_seconds,omitempty"`
	TreeUUID          string   `json:"tree_uuid,omitempty"`
	CreditTreeUUID    string   `json:"credit_tree_uuid,omitempty"`
	RiskGradeTreeUUID string   `json:"risk_grade_tree_uuid,omitempty"`
	UseAddOnServices  []string `json:"use_add_on_services,omitempty"`
}

type AddOnService struct {
	Dispatch []Dispatch `json:"dispatch"`
	Enabled  bool       `json:"enabled"`
}

type Dispatch struct {
	State           string     `json:"state"`
	UseTSFormFields []string   `json:"use_ts_form_fields,omitempty"`
	UseIDCardFields []string   `json:"use_id_card_fields,omitempty"`
	CICTypes        []string   `json:"cic_types,omitempty"`
	ExtraInfo       *ExtraInfo `json:"extra_info,omitempty"`
}

type ExtraInfo struct {
	WaitForValidServices bool `json:"wait_for_valid_services,omitempty"`
	DisableDeduplication bool `json:"disable_deduplication,omitempty"`
}

type EKYCSettings struct {
	Face           FaceSettings    `json:"face"`
	IDCard         IDCardSettings  `json:"id_card"`
	LogEvent       bool            `json:"log_event"`
	ClientSettings *ClientSettings `json:"client_settings,omitempty"`
}

type FaceFeatures struct {
	FaceCompare FeatureSettings `json:"face_compare"`
	Liveness    FeatureSettings `json:"liveness"`
	Sanity      FeatureSettings `json:"sanity"`
	SearchFace  FeatureSettings `json:"search_face"`
}

type FeatureSettings struct {
	Enabled              bool   `json:"enabled"`
	ExceededFailsCommand string `json:"exceeded_fails_command,omitempty"`
	MaxFails             int    `json:"max_fails,omitempty"`
	RetryCounterGroup    string `json:"retry_counter_group,omitempty"`
	CounterSeconds       int    `json:"counter_seconds,omitempty"`
	LockSeconds          int    `json:"lock_seconds,omitempty"`
	VideoEnabled         bool   `json:"video_enabled,omitempty"`
	SelfieType           string `json:"selfie_type,omitempty"`
	Level                string `json:"level,omitempty"`
	TreeUUID             string `json:"tree_uuid,omitempty"`
	TimeoutSeconds       int    `json:"timeout_seconds,omitempty"`
	SyncCheckEnabled     bool   `json:"sync_check_enabled,omitempty"`
}

type FaceSettings struct {
	Features FaceFeatures `json:"features"`
}

type IDCardSettings struct {
	BackFeatures    IDCardFeatures `json:"back_features"`
	Features        IDCardFeatures `json:"features"`
	FrontFeatures   IDCardFeatures `json:"front_features"`
	EnableIndexFace bool           `json:"enable_index_face"`
}

type ClientSettings struct {
	FlowID string `json:"flow_id,omitempty"`
}

type IDCardFeatures struct {
	FaceCompare            FeatureSettings `json:"face_compare"`
	IDTampering            FeatureSettings `json:"id_tampering"`
	OCR                    FeatureSettings `json:"ocr"`
	Sanity                 FeatureSettings `json:"sanity"`
	SearchIDPhoto          FeatureSettings `json:"search_id_photo"`
	RequiredIDCardWithChip FeatureSettings `json:"required_id_card_with_chip,omitempty"`
	QRCode                 FeatureSettings `json:"qr_code,omitempty"`
	Index                  IndexSettings   `json:"index,omitempty"`
}

type IndexSettings struct {
	FromState   string `json:"from_state,omitempty"`
	ToState     string `json:"to_state,omitempty"`
	UseWorkflow bool   `json:"use_workflow,omitempty"`
}

type OTPService struct {
	Auth          OTPAuth        `json:"auth"`
	CreditInsight *CreditInsight `json:"credit_insight,omitempty"`
}

type CreditInsight struct {
	SMSVersion string `json:"sms_version,omitempty"`
}

type OTPAuth struct {
	SMSTemplate string `json:"sms_template"`
	Type        string `json:"type"`
}

type DocumentSchema struct {
	Expressions []interface{} `json:"expressions"`
}
