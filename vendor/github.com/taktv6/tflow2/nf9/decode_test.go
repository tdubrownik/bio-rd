// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nf9

import (
	"net"
	"testing"

	"github.com/taktv6/tflow2/convert"
)

/*func TestDecode(t *testing.T) {
	s := []byte{0, 0, 0, 8, 201, 173, 78, 201, 2, 0, 229, 27, 75, 201, 2, 0, 0, 0, 0, 249, 0, 6, 187, 71, 213, 103, 123, 68, 213, 103, 10, 5, 0, 0, 11, 0, 0, 0, 15, 0, 65, 0, 15, 0, 65, 0, 26, 187, 1, 239, 181, 153, 192, 66, 185, 34, 93, 13, 31, 65, 195, 66, 185, 0, 8, 201, 173, 78, 201, 2, 0, 229, 27, 75, 201, 2, 0, 0, 0, 0, 249, 0, 6, 183, 71, 213, 103, 7, 39, 213, 103, 224, 156, 0, 0, 153, 2, 0, 0, 15, 0, 65, 0, 15, 0, 65, 0, 30, 80, 0, 105, 187, 153, 192, 66, 185, 136, 100, 80, 151, 65, 195, 66, 185, 128, 0, 221, 1, 0, 0, 0, 0, 8, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 6, 180, 71, 213, 103, 164, 62, 213, 103, 160, 0, 0, 0, 4, 0, 0, 0, 21, 0, 28, 0, 21, 0, 28, 0, 16, 80, 0, 87, 204, 185, 192, 66, 185, 147, 23, 217, 172, 93, 193, 66, 185, 64, 0, 223, 1, 0, 0, 0, 8, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 0, 0, 137, 0, 6, 191, 71, 213, 103, 248, 44, 213, 103, 125, 17, 0, 0, 57, 0, 0, 0, 21, 0, 72, 0, 21, 0, 72, 0, 24, 187, 1, 145, 226, 185, 192, 66, 185, 88, 160, 125, 74, 84, 193, 66, 185, 0, 8, 237, 240, 149, 1, 185, 92, 229, 27, 75, 201, 2, 0, 0, 0, 0, 241, 0, 6, 194, 71, 213, 103, 124, 61, 213, 103, 164, 0, 0, 0, 3, 0, 0, 0, 39, 0, 22, 0, 39, 0, 22, 0, 19, 89, 216, 80, 0, 235, 5, 64, 100, 41, 193, 66, 185, 243, 121, 19, 50, 128, 0, 221, 1, 0, 0, 0, 221, 134, 201, 173, 78, 201, 2, 0, 229, 27, 75, 201, 2, 0, 0, 0, 0, 129, 0, 6, 185, 71, 213, 103, 234, 62, 213, 103, 201, 53, 0, 0, 177, 0, 0, 0, 21, 0, 73, 0, 21, 0, 73, 0, 24, 187, 1, 181, 211, 201, 173, 78, 254, 255, 201, 2, 2, 0, 0, 0, 0, 0, 0, 128, 254, 11, 0, 0, 0, 0, 0, 0, 0, 35, 0, 14, 64, 80, 20, 0, 42, 179, 79, 172, 109, 9, 172, 109, 133, 55, 19, 15, 48, 96, 34, 3, 42, 104, 0, 222, 1, 0, 0, 8, 237, 240, 149, 1, 185, 92, 229, 27, 75, 201, 2, 0, 0, 0, 0, 241, 0, 6, 194, 71, 213, 103, 226, 68, 213, 103, 201, 21, 0, 0, 18, 0, 0, 0, 116, 0, 22, 0, 116, 0, 22, 0, 26, 172, 230, 187, 1, 101, 0, 64, 100, 49, 193, 66, 185, 36, 107, 175, 54, 0, 8, 201, 173, 78, 201, 2, 0, 229, 27, 75, 201, 2, 0, 0, 0, 0, 241, 0, 6, 194, 71, 213, 103, 222, 67, 213, 103, 211, 5, 0, 0, 6, 0, 0, 0, 15, 0, 65, 0, 15, 0, 65, 0, 27, 80, 0, 243, 165, 153, 192, 66, 185, 138, 98, 227, 172, 65, 195, 66, 185, 0, 8, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 0, 0, 129, 0, 6, 188, 71, 213, 103, 188, 71, 213, 103, 122, 0, 0, 0, 1, 0, 0, 0, 184, 0, 15, 0, 184, 0, 15, 0, 24, 145, 193, 230, 15, 213, 1, 64, 100, 16, 193, 66, 185, 210, 7, 182, 193, 188, 0, 221, 1, 0, 0, 0, 221, 134, 212, 186, 30, 36, 78, 204, 229, 27, 75, 201, 2, 0, 0, 0, 0, 241, 0, 6, 179, 71, 213, 103, 215, 49, 213, 103, 248, 17, 0, 0, 13, 0, 0, 0, 119, 0, 16, 0, 119, 0, 16, 0, 26, 2, 201, 187, 1, 220, 90, 4, 46, 254, 94, 0, 2, 0, 0, 0, 0, 0, 0, 128, 254, 34, 44, 143, 56, 96, 67, 7, 176, 0, 70, 21, 1, 96, 34, 3, 42, 142, 0, 0, 0, 12, 176, 206, 250, 14, 3, 19, 240, 128, 40, 3, 42, 104, 0, 222, 1, 0, 0, 0, 0, 8, 237, 240, 149, 1, 185, 92, 229, 27, 75, 201, 2, 0, 0, 0, 0, 241, 0, 6, 183, 71, 213, 103, 47, 68, 213, 103, 54, 23, 0, 0, 10, 0, 0, 0, 73, 0, 22, 0, 73, 0, 22, 0, 26, 79, 154, 187, 1, 59, 4, 64, 100, 85, 193, 66, 185, 43, 156, 16, 199, 68, 0, 221, 1, 0, 0, 0, 221, 134, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 6, 179, 71, 213, 103, 179, 71, 213, 103, 61, 0, 0, 0, 1, 0, 0, 0, 21, 0, 34, 0, 21, 0, 34, 0, 16, 80, 0, 251, 209, 201, 173, 78, 254, 255, 201, 2, 2, 0, 0, 0, 0, 0, 0, 128, 254, 16, 32, 0, 0, 0, 0, 0, 0, 33, 8, 1, 64, 80, 20, 0, 42, 159, 9, 125, 55, 155, 45, 217, 165, 2, 0, 20, 1, 96, 34, 3, 42, 100, 0, 220, 1, 0, 8, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 1, 179, 71, 213, 103, 19, 59, 213, 103, 152, 0, 0, 0, 2, 0, 0, 0, 15, 0, 93, 0, 15, 0, 93, 0, 3, 3, 0, 0, 153, 192, 66, 185, 119, 160, 222, 68, 31, 194, 66, 185, 60, 0, 228, 1, 2, 0, 0, 1, 6, 0, 56, 0, 6, 0, 80, 0, 1, 0, 5, 0, 1, 0, 4, 0, 4, 0, 21, 0, 4, 0, 22, 0, 4, 0, 1, 0, 4, 0, 2, 0, 2, 0, 253, 0, 2, 0, 252, 0, 2, 0, 14, 0, 2, 0, 10, 0, 2, 0, 11, 0, 2, 0, 7, 0, 4, 0, 15, 0, 4, 0, 12, 0, 4, 0, 8, 0, 18, 0, 228, 1, 80, 0, 0, 0, 0, 0, 0, 0, 8, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 0, 0, 129, 0, 6, 192, 71, 213, 103, 192, 71, 213, 103, 52, 0, 0, 0, 1, 0, 0, 0, 21, 0, 178, 0, 21, 0, 178, 0, 16, 187, 1, 62, 139, 185, 192, 66, 185, 168, 8, 125, 74, 54, 194, 66, 185, 68, 0, 221, 1, 0, 0, 0, 0, 8, 201, 173, 78, 201, 2, 0, 229, 27, 75, 201, 2, 0, 0, 17, 189, 71, 213, 103, 189, 71, 213, 103, 76, 0, 0, 0, 1, 0, 0, 0, 15, 0, 65, 0, 15, 0, 65, 0, 0, 123, 0, 234, 170, 153, 192, 66, 185, 221, 186, 9, 5, 65, 195, 66, 185, 64, 0, 223, 1, 0, 0, 0, 8, 201, 173, 78, 201, 2, 0, 229, 27, 75, 201, 2, 0, 0, 0, 0, 241, 0, 6, 188, 71, 213, 103, 103, 71, 213, 103, 247, 0, 0, 0, 3, 0, 0, 0, 26, 0, 21, 0, 26, 0, 21, 0, 26, 46, 155, 80, 0, 81, 4, 64, 100, 102, 193, 66, 185, 46, 208, 58, 216, 0, 8, 201, 173, 78, 201, 2, 0, 229, 27, 75, 201, 2, 0, 0, 0, 0, 241, 0, 6, 179, 71, 213, 103, 101, 71, 213, 103, 247, 0, 0, 0, 3, 0, 0, 0, 26, 0, 21, 0, 26, 0, 21, 0, 26, 145, 155, 80, 0, 81, 4, 64, 100, 102, 193, 66, 185, 46, 208, 58, 216, 128, 0, 221, 1, 0, 0, 0, 221, 134, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 0, 0, 129, 0, 6, 180, 71, 213, 103, 134, 71, 213, 103, 38, 3, 0, 0, 2, 0, 0, 0, 21, 0, 34, 0, 21, 0, 34, 0, 24, 187, 1, 218, 156, 201, 173, 78, 254, 255, 201, 2, 2, 0, 0, 0, 0, 0, 0, 128, 254, 11, 0, 0, 0, 0, 0, 0, 0, 78, 0, 1, 64, 80, 20, 0, 42, 35, 211, 203, 103, 92, 74, 192, 76, 7, 0, 20, 1, 96, 34, 3, 42, 104, 0, 222, 1, 0, 0, 0, 0, 167, 51, 204, 11, 128, 207, 118, 88, 75, 91, 213, 103, 19, 0, 9, 0}
	s = convert.Reverse(s)

	packet, err := Decode(s, net.IP([]byte{1, 1, 1, 1}))
	if err != nil {
		t.Errorf("Decoding packet failed: %v\n", err)
	}

	flowSet := []byte{0, 0, 0, 221, 134, 100, 249, 80, 201, 2, 0, 228, 27, 75, 201, 2, 0, 0, 0, 0, 129, 0, 6, 180, 71, 213, 103, 134, 71, 213, 103, 38, 3, 0, 0, 2, 0, 0, 0, 21, 0, 34, 0, 21, 0, 34, 0, 24, 187, 1, 218, 156, 201, 173, 78, 254, 255, 201, 2, 2, 0, 0, 0, 0, 0, 0, 128, 254, 11, 0, 0, 0, 0, 0, 0, 0, 78, 0, 1, 64, 80, 20, 0, 42, 35, 211, 203, 103, 92, 74, 192, 76, 7, 0, 20, 1, 96, 34, 3, 42}

	if !testEq(packet.FlowSets[0].Flows, flowSet) {
		t.Errorf("Decoded FlowSet is not the expected one. Got: %v, Expected: %v\n", packet.FlowSets[0].Flows, flowSet)
	}
}*/

func TestDecode2(t *testing.T) {
	s := []byte{
		8, 0, // Length
		44, 0, // Type

		8, 0, // Length
		43, 0, // Type

		8, 0, // Length
		42, 0, // Type

		8, 0, // Length
		41, 0, // Type

		4, 0, // Scope 1 Field Length
		1, 0, // Scope 1 Field Type = 1 = System

		16, 0, // OptionLength
		4, 0, // OptionScopeLength
		10, 1, // TemplateID

		30, 0, // Length
		1, 0, // FlowSetID

		0, 0, 0, 0, //Source ID
		167, 51, 204, 11, // Sequence Number
		128, 207, 118, 88, // UNIX secs
		75, 91, 213, 103, // sysUpTime
		1, 0, // Count
		9, 0} // Version
	s = convert.Reverse(s)

	_, err := Decode(s, net.IP([]byte{1, 1, 1, 1}))
	if err != nil {
		t.Errorf("Decoding packet failed: %v\n", err)
	}

}

func testEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
