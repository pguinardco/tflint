// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSagemakerModelInvalidNameRule checks the pattern is valid
type AwsSagemakerModelInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerModelInvalidNameRule returns new rule with default attributes
func NewAwsSagemakerModelInvalidNameRule() *AwsSagemakerModelInvalidNameRule {
	return &AwsSagemakerModelInvalidNameRule{
		resourceType:  "aws_sagemaker_model",
		attributeName: "name",
		max:           63,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9](-*[a-zA-Z0-9])*`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerModelInvalidNameRule) Name() string {
	return "aws_sagemaker_model_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerModelInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerModelInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerModelInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerModelInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
