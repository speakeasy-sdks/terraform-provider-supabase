// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ReadReplicaRegion - Region you want your read replica to reside in
type ReadReplicaRegion string

const (
	ReadReplicaRegionUsEast1      ReadReplicaRegion = "us-east-1"
	ReadReplicaRegionUsEast2      ReadReplicaRegion = "us-east-2"
	ReadReplicaRegionUsWest1      ReadReplicaRegion = "us-west-1"
	ReadReplicaRegionUsWest2      ReadReplicaRegion = "us-west-2"
	ReadReplicaRegionApEast1      ReadReplicaRegion = "ap-east-1"
	ReadReplicaRegionApSoutheast1 ReadReplicaRegion = "ap-southeast-1"
	ReadReplicaRegionApNortheast1 ReadReplicaRegion = "ap-northeast-1"
	ReadReplicaRegionApNortheast2 ReadReplicaRegion = "ap-northeast-2"
	ReadReplicaRegionApSoutheast2 ReadReplicaRegion = "ap-southeast-2"
	ReadReplicaRegionEuWest1      ReadReplicaRegion = "eu-west-1"
	ReadReplicaRegionEuWest2      ReadReplicaRegion = "eu-west-2"
	ReadReplicaRegionEuWest3      ReadReplicaRegion = "eu-west-3"
	ReadReplicaRegionEuNorth1     ReadReplicaRegion = "eu-north-1"
	ReadReplicaRegionEuCentral1   ReadReplicaRegion = "eu-central-1"
	ReadReplicaRegionEuCentral2   ReadReplicaRegion = "eu-central-2"
	ReadReplicaRegionCaCentral1   ReadReplicaRegion = "ca-central-1"
	ReadReplicaRegionApSouth1     ReadReplicaRegion = "ap-south-1"
	ReadReplicaRegionSaEast1      ReadReplicaRegion = "sa-east-1"
)

func (e ReadReplicaRegion) ToPointer() *ReadReplicaRegion {
	return &e
}
func (e *ReadReplicaRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-central-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "sa-east-1":
		*e = ReadReplicaRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReadReplicaRegion: %v", v)
	}
}

type SetUpReadReplicaBody struct {
	// Region you want your read replica to reside in
	ReadReplicaRegion ReadReplicaRegion `json:"read_replica_region"`
}

func (o *SetUpReadReplicaBody) GetReadReplicaRegion() ReadReplicaRegion {
	if o == nil {
		return ReadReplicaRegion("")
	}
	return o.ReadReplicaRegion
}
