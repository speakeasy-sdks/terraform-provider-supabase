// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ProjectUpgradeInitiateResponse struct {
	TrackingID string `json:"tracking_id"`
}

func (o *ProjectUpgradeInitiateResponse) GetTrackingID() string {
	if o == nil {
		return ""
	}
	return o.TrackingID
}
