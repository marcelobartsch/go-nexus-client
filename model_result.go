/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.29.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus

type Result struct {
	Healthy bool `json:"healthy,omitempty"`
	Message string `json:"message,omitempty"`
	Error_ *Throwable `json:"error,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Time int64 `json:"time,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
