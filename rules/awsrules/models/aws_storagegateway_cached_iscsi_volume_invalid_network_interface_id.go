// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule checks the pattern is valid
type AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule returns new rule with default attributes
func NewAwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule() *AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule {
	return &AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule{
		resourceType:  "aws_storagegateway_cached_iscsi_volume",
		attributeName: "network_interface_id",
		pattern:       regexp.MustCompile(`^\A(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}\z$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule) Name() string {
	return "aws_storagegateway_cached_iscsi_volume_invalid_network_interface_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\A(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}\z$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
