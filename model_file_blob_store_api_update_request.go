/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.29.0-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus

type FileBlobStoreApiUpdateRequest struct {
	SoftQuota *BlobStoreApiSoftQuota `json:"softQuota,omitempty"`
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system Nexus Repository Manager has access to or it can be a path relative to the sonatype-work directory.
	Path string `json:"path,omitempty"`
}
