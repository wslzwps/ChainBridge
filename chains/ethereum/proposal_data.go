// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

// constructErc20ProposalData returns the bytes to construct a proposal suitable for Erc20
func ConstructErc20ProposalData(amount []byte, recipient []byte) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(amount, 32)...) // amount (uint256)
	fmt.Printf("Amount as integer: %d\n", new(big.Int).SetBytes(amount))
	recipientLen := big.NewInt(int64(len(recipient))).Bytes()
	fmt.Printf("Recipient length as integer: %d\n", new(big.Int).SetBytes(recipientLen))
	data = append(data, common.LeftPadBytes(recipientLen, 32)...) // length of recipient (uint256)
	data = append(data, recipient...)

	// Print recipient as hex
	fmt.Printf("Recipient as hex: %x\n", recipient)

	// Print final data as hex
	fmt.Printf("Final data as hex: %x\n", data) // recipient ([]byte)

	return data
}

// constructErc721ProposalData returns the bytes to construct a proposal suitable for Erc721
func ConstructErc721ProposalData(tokenId []byte, recipient []byte, metadata []byte) []byte {
	var data []byte
	data = append(data, common.LeftPadBytes(tokenId, 32)...) // tokenId ([]byte)

	recipientLen := big.NewInt(int64(len(recipient))).Bytes()
	data = append(data, common.LeftPadBytes(recipientLen, 32)...) // length of recipient
	data = append(data, recipient...)                             // recipient ([]byte)

	metadataLen := big.NewInt(int64(len(metadata))).Bytes()
	data = append(data, common.LeftPadBytes(metadataLen, 32)...) // length of metadata (uint256)
	data = append(data, metadata...)                             // metadata ([]byte)
	return data
}

// constructGenericProposalData returns the bytes to construct a generic proposal
func ConstructGenericProposalData(metadata []byte) []byte {
	var data []byte

	metadataLen := big.NewInt(int64(len(metadata)))
	data = append(data, math.PaddedBigBytes(metadataLen, 32)...) // length of metadata (uint256)
	data = append(data, metadata...)                             // metadata ([]byte)
	return data
}
