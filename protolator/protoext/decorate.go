/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoext

import (
	"gitee.com/zhaochuninhefei/fabric-config-gm/protolator/protoext/commonext"
	"gitee.com/zhaochuninhefei/fabric-config-gm/protolator/protoext/ledger/rwsetext"
	"gitee.com/zhaochuninhefei/fabric-config-gm/protolator/protoext/mspext"
	"gitee.com/zhaochuninhefei/fabric-config-gm/protolator/protoext/ordererext"
	"gitee.com/zhaochuninhefei/fabric-config-gm/protolator/protoext/peerext"
	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/common"
	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/ledger/rwset"
	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/msp"
	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/orderer"
	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/peer"
	"github.com/golang/protobuf/proto"
)

// Docorate will add additional capabilities to some protobuf messages that
// enable proper JSON marshalling and unmarshalling in protolator.
func Decorate(msg proto.Message) proto.Message {
	switch m := msg.(type) {
	case *common.BlockData:
		return &commonext.BlockData{BlockData: m}
	case *common.Config:
		return &commonext.Config{Config: m}
	case *common.ConfigSignature:
		return &commonext.ConfigSignature{ConfigSignature: m}
	case *common.ConfigUpdate:
		return &commonext.ConfigUpdate{ConfigUpdate: m}
	case *common.ConfigUpdateEnvelope:
		return &commonext.ConfigUpdateEnvelope{ConfigUpdateEnvelope: m}
	case *common.Envelope:
		return &commonext.Envelope{Envelope: m}
	case *common.Header:
		return &commonext.Header{Header: m}
	case *common.ChannelHeader:
		return &commonext.ChannelHeader{ChannelHeader: m}
	case *common.SignatureHeader:
		return &commonext.SignatureHeader{SignatureHeader: m}
	case *common.Payload:
		return &commonext.Payload{Payload: m}
	case *common.Policy:
		return &commonext.Policy{Policy: m}

	case *msp.MSPConfig:
		return &mspext.MSPConfig{MSPConfig: m}
	case *msp.MSPPrincipal:
		return &mspext.MSPPrincipal{MSPPrincipal: m}

	case *orderer.ConsensusType:
		return &ordererext.ConsensusType{ConsensusType: m}

	case *peer.ChaincodeAction:
		return &peerext.ChaincodeAction{ChaincodeAction: m}
	case *peer.ChaincodeActionPayload:
		return &peerext.ChaincodeActionPayload{ChaincodeActionPayload: m}
	case *peer.ChaincodeEndorsedAction:
		return &peerext.ChaincodeEndorsedAction{ChaincodeEndorsedAction: m}
	case *peer.ChaincodeProposalPayload:
		return &peerext.ChaincodeProposalPayload{ChaincodeProposalPayload: m}
	case *peer.ProposalResponsePayload:
		return &peerext.ProposalResponsePayload{ProposalResponsePayload: m}
	case *peer.TransactionAction:
		return &peerext.TransactionAction{TransactionAction: m}

	case *rwset.TxReadWriteSet:
		return &rwsetext.TxReadWriteSet{TxReadWriteSet: m}

	default:
		return msg
	}
}
