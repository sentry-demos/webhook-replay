package main

// https://docs.mparticle.com/integrations/webhook/event/
type Webhook struct {
	Action string                 `json:"action"`
	Actor  map[string]interface{} `json:"actor"`
	Data   Data                   `json:"data"`

	// WORKS..
	// Data map[string]interface{} `json:"data"`

	// Crash_reports    []CrashReport            `json:"events"`
	// Environment      string                   `json:"environment"`
	// User_identities  []map[string]interface{} `json:"user_identities"`
	// Device_info      map[string]interface{}   `json:"device_info"`
	// Application_info map[string]interface{}   `json:"application_info"`
}

type Data struct {
	Error map[string]interface{} `json:"error"`
}
