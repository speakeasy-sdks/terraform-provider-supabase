// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type StorageFeatureImageTransformation struct {
	Enabled bool `json:"enabled"`
}

func (o *StorageFeatureImageTransformation) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}