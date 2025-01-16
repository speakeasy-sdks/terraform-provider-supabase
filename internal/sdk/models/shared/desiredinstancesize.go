// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DesiredInstanceSize string

const (
	DesiredInstanceSizeMicro         DesiredInstanceSize = "micro"
	DesiredInstanceSizeSmall         DesiredInstanceSize = "small"
	DesiredInstanceSizeMedium        DesiredInstanceSize = "medium"
	DesiredInstanceSizeLarge         DesiredInstanceSize = "large"
	DesiredInstanceSizeXlarge        DesiredInstanceSize = "xlarge"
	DesiredInstanceSizeTwoxlarge     DesiredInstanceSize = "2xlarge"
	DesiredInstanceSizeFourxlarge    DesiredInstanceSize = "4xlarge"
	DesiredInstanceSizeEightxlarge   DesiredInstanceSize = "8xlarge"
	DesiredInstanceSizeTwelvexlarge  DesiredInstanceSize = "12xlarge"
	DesiredInstanceSizeSixteenxlarge DesiredInstanceSize = "16xlarge"
)

func (e DesiredInstanceSize) ToPointer() *DesiredInstanceSize {
	return &e
}
func (e *DesiredInstanceSize) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "micro":
		fallthrough
	case "small":
		fallthrough
	case "medium":
		fallthrough
	case "large":
		fallthrough
	case "xlarge":
		fallthrough
	case "2xlarge":
		fallthrough
	case "4xlarge":
		fallthrough
	case "8xlarge":
		fallthrough
	case "12xlarge":
		fallthrough
	case "16xlarge":
		*e = DesiredInstanceSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DesiredInstanceSize: %v", v)
	}
}
