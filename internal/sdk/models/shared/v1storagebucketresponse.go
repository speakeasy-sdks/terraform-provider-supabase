// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V1StorageBucketResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Public    bool   `json:"public"`
}

func (o *V1StorageBucketResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *V1StorageBucketResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *V1StorageBucketResponse) GetOwner() string {
	if o == nil {
		return ""
	}
	return o.Owner
}

func (o *V1StorageBucketResponse) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *V1StorageBucketResponse) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *V1StorageBucketResponse) GetPublic() bool {
	if o == nil {
		return false
	}
	return o.Public
}
