/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"testing"
)

func Test_validateInternalIP(t *testing.T) {
	host := &HetznerBareMetalHost{}
	if errs := validateInternalIP(host); len(errs) > 0 {
		t.Fatalf("no annotation: got errors %v", errs)
	}

	tests := []struct {
		internalIP  string
		expectError bool
	}{
		{internalIP: "", expectError: true},
		{internalIP: "10.0.1.2", expectError: false},
		{internalIP: "192.168.100.5", expectError: false},
		{internalIP: "10.0.1.2/24", expectError: true},
		{internalIP: "203.0.113.5", expectError: true},
		{internalIP: "fd00::2", expectError: true},
		{internalIP: "not-an-ip", expectError: true},
	}

	for _, tt := range tests {
		host := &HetznerBareMetalHost{}
		host.SetAnnotations(map[string]string{InternalIPAnnotation: tt.internalIP})
		if gotError := len(validateInternalIP(host)) > 0; gotError != tt.expectError {
			t.Fatalf("internalIP %q: got error %v, want %v", tt.internalIP, gotError, tt.expectError)
		}
	}
}
