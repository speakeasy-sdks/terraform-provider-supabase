// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type StorageFeatures struct {
	ImageTransformation StorageFeatureImageTransformation `json:"imageTransformation"`
}

func (o *StorageFeatures) GetImageTransformation() StorageFeatureImageTransformation {
	if o == nil {
		return StorageFeatureImageTransformation{}
	}
	return o.ImageTransformation
}
