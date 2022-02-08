// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/ia/alerts.proto

package iav1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/percona/pmm/api/managementpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Alert) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	if this.Rule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Rule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Rule", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *ListAlertsRequest) Validate() error {
	if this.PageParams != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PageParams); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PageParams", err)
		}
	}
	return nil
}
func (this *ListAlertsResponse) Validate() error {
	for _, item := range this.Alerts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Alerts", err)
			}
		}
	}
	if this.Totals != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Totals); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Totals", err)
		}
	}
	return nil
}
func (this *ToggleAlertsRequest) Validate() error {
	return nil
}
func (this *ToggleAlertsResponse) Validate() error {
	return nil
}
