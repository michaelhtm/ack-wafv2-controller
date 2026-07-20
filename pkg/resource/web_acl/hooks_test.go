// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package web_acl

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"

	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
)

func TestValidateLoggingResourceARN(t *testing.T) {
	const webACLARN = "arn:aws:wafv2:us-west-2:111122223333:regional/webacl/my-acl/abc-123"

	cases := []struct {
		name         string
		resourceARN  *string
		wantErr      bool
		wantTerminal bool
	}{
		{
			name:        "nil is allowed",
			resourceARN: nil,
			wantErr:     false,
		},
		{
			name:        "matching own ARN is allowed",
			resourceARN: aws.String(webACLARN),
			wantErr:     false,
		},
		{
			name:         "different ARN is a terminal error",
			resourceARN:  aws.String("arn:aws:wafv2:us-west-2:111122223333:regional/webacl/other-acl/xyz-789"),
			wantErr:      true,
			wantTerminal: true,
		},
		{
			name:         "empty string is a terminal error",
			resourceARN:  aws.String(""),
			wantErr:      true,
			wantTerminal: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateLoggingResourceARN(tc.resourceARN, webACLARN)
			if tc.wantErr && err == nil {
				t.Fatalf("expected an error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if tc.wantTerminal {
				var termErr *ackerr.TerminalError
				if !errors.As(err, &termErr) {
					t.Fatalf("expected a terminal error, got %T: %v", err, err)
				}
			}
		})
	}
}
