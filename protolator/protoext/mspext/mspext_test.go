/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mspext_test

import (
	"gitee.com/zhaochuninhefei/fabric-config-gm/protolator"
	"gitee.com/zhaochuninhefei/fabric-config-gm/protolator/protoext/mspext"
)

// ensure structs implement expected interfaces
var (
	_ protolator.VariablyOpaqueFieldProto = &mspext.MSPConfig{}
	_ protolator.DecoratedProto           = &mspext.MSPConfig{}

	_ protolator.VariablyOpaqueFieldProto = &mspext.MSPPrincipal{}
	_ protolator.DecoratedProto           = &mspext.MSPPrincipal{}
)
