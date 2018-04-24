/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package scc

import (
	//import system chaincodes here
	"github.com/hyperledger/fabric/core/scc/cscc"
	"github.com/hyperledger/fabric/core/scc/lscc"
	"github.com/hyperledger/fabric/core/scc/qscc"
	"github.com/hyperledger/fabric/core/scc/vscc"
)

func builtInSystemChaincodes(p *Provider) []*SystemChaincode {
	return []*SystemChaincode{
		{
			Enabled:           true,
			Name:              "cscc",
			Path:              "github.com/hyperledger/fabric/core/scc/cscc",
			InitArgs:          nil,
			Chaincode:         cscc.New(p),
			InvokableExternal: true, // cscc is invoked to join a channel
		},
		{
			Enabled:           true,
			Name:              "lscc",
			Path:              "github.com/hyperledger/fabric/core/scc/lscc",
			InitArgs:          nil,
			Chaincode:         lscc.New(p),
			InvokableExternal: true, // lscc is invoked to deploy new chaincodes
			InvokableCC2CC:    true, // lscc can be invoked by other chaincodes
		},
		{
			Enabled:   true,
			Name:      "vscc",
			Path:      "github.com/hyperledger/fabric/core/scc/vscc",
			InitArgs:  nil,
			Chaincode: vscc.New(p),
		},
		{
			Enabled:           true,
			Name:              "qscc",
			Path:              "github.com/hyperledger/fabric/core/chaincode/qscc",
			InitArgs:          nil,
			Chaincode:         qscc.New(),
			InvokableExternal: true, // qscc can be invoked to retrieve blocks
			InvokableCC2CC:    true, // qscc can be invoked to retrieve blocks also by a cc
		},
	}
}

//DeploySysCCs is the hook for system chaincodes where system chaincodes are registered with the fabric
//note the chaincode must still be deployed and launched like a user chaincode will be
func (p *Provider) DeploySysCCs(chainID string) {
	for _, sysCC := range p.SysCCs {
		sysCC.deploySysCC(chainID)
	}
}

//DeDeploySysCCs is used in unit tests to stop and remove the system chaincodes before
//restarting them in the same process. This allows clean start of the system
//in the same process
func (p *Provider) DeDeploySysCCs(chainID string) {
	for _, sysCC := range p.SysCCs {
		sysCC.deDeploySysCC(chainID)
	}
}
