/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.29.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus

type UploadDefinitionXo struct {
	Format string `json:"format,omitempty"`
	MultipleUpload bool `json:"multipleUpload,omitempty"`
	ComponentFields []UploadFieldDefinitionXo `json:"componentFields,omitempty"`
	AssetFields []UploadFieldDefinitionXo `json:"assetFields,omitempty"`
}
