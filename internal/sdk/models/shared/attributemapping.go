// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AttributeMapping struct {
	Keys map[string]AttributeValue `json:"keys"`
}

func (o *AttributeMapping) GetKeys() map[string]AttributeValue {
	if o == nil {
		return map[string]AttributeValue{}
	}
	return o.Keys
}