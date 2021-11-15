package main

import (
	"fmt"
)

// Platform names that Sentry.io is expecting
const COCOA = "cocoa"
const JAVA = "java"

// https://develop.sentry.dev/sdk/event-payloads/#optional-attributes see 'error' attribute
// https://github.com/getsentry/sentry/blob/master/src/sentry/models/eventerror.py#L8
type ProcessingWarning struct {
	Type  string
	Value string
}

// Several types of Processing Errors for when user_identities, environment, application_info, device_info or crash_reports are not provided
const MISSING_ATTRIBUTE = "missing_attribute"

type Value struct {
	Type       string                 `json:"type"`
	Value      string                 `json:"value"`
	Stacktrace map[string]interface{} `json:"stacktrace"`
	Thread_id  int32                  `json:"thread_id"`
	Id         int32                  `json:"id"`
	Crashed    bool                   `json:"crashed"`
}
type Frame struct {
	Function         string `json:"function"`
	Filename         string `json:"filename"`
	Lineno           int32  `json:"lineno"`
	Package          string `json:"package"`
	Instruction_addr string `json:"instruction_addr"`
	In_app           bool   `json:"in_app"`
}

type Error struct {
	Breadcrumbs     []map[string]interface{} `json:"breadcrumbs"`
	Contexts        map[string]interface{}   `json:"contexts"`
	Culprit         string                   `json:"culprit"`
	Environment     string                   `json:"environment"`
	EventId         string                   `json:"event_id"`
	Exception       map[string]interface{}   `json:"exception"`
	Extra           map[string]interface{}   `json:"extra"`
	Fingerprint     []string                 `json:"fingerprint"`
	Grouping_config map[string]interface{}   `json:"grouping_config"`
	Hashes          []string                 `json:"hashes"`
	Key_id          string                   `json:"key_id"`
	Level           string                   `json:"level"`
	Logger          string                   `json:"logger"`
	Metadata        map[string]interface{}   `json:"metadata"`
	Message         string                   `json:"message"`
	Modules         map[string]interface{}   `json:"modules"`
	Platform        string                   `json:"platform"`
	Received        float64                  `json:"received"`
	Project         int                      `json:"project"`
	Release         string                   `json:"release"`
	Request         map[string]interface{}   `json:"request"`
	Sdk             map[string]interface{}   `json:"sdk"`
	Tags            [][]string               `json:"tags"`
	Timestamp       float64                  `json:"timestamp"`
	Threads         map[string]interface{}   `json:"threads"`
	Type            string                   `json:"type"`
	User            map[string]interface{}   `json:"user"`
	Version         string                   `json:"version"`
}

// Can set any parts of the Webhook or Custom_event as a Tag, which is a K,V pair and reportable in Discover
func (e *Error) setTags(webhook Webhook) {
	fmt.Println("> setTags")

	e.Tags = append(e.Tags, []string{"translation_layer_version", "6"})

	tagItem := []string{"type", "test"}
	e.Tags = append(e.Tags, tagItem)

	// deviceInfo := webhook.Device_info
	// applicationInfo := webhook.Application_info

	// deviceModel, ok := deviceInfo["device_model"].(string)
	// if ok == true {
	// 	tagItem := []string{"device_model", deviceModel}
	// 	e.Tags = append(e.Tags, tagItem)
	// }

	// buildIdentifier, ok := deviceInfo["build_identifier"].(string)
	// if ok == true {
	// 	tagItem := []string{"build_identifier", buildIdentifier}
	// 	e.Tags = append(e.Tags, tagItem)
	// }

	// applicationVersion, ok := applicationInfo["application_version"].(string)
	// if ok == true {
	// 	tagItem := []string{"application_version", applicationVersion}
	// 	e.Tags = append(e.Tags, tagItem)
	// }
}
