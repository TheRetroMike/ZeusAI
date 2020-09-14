package protowire

import (
	"github.com/kaspanet/kaspad/app/appmessage"
)

func (x *KaspadMessage_GetBlockRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.GetBlockRequestMessage{
		Hash:                          x.GetBlockRequest.Hash,
		SubnetworkID:                  x.GetBlockRequest.SubnetworkId,
		IncludeBlockHex:               x.GetBlockRequest.IncludeBlockHex,
		IncludeBlockVerboseData:       x.GetBlockRequest.IncludeBlockVerboseData,
		IncludeTransactionVerboseData: x.GetBlockRequest.IncludeTransactionVerboseData,
	}, nil
}

func (x *KaspadMessage_GetBlockRequest) fromAppMessage(message *appmessage.GetBlockRequestMessage) error {
	x.GetBlockRequest = &GetBlockRequestMessage{
		Hash:                          message.Hash,
		SubnetworkId:                  message.SubnetworkID,
		IncludeBlockHex:               message.IncludeBlockHex,
		IncludeBlockVerboseData:       message.IncludeBlockVerboseData,
		IncludeTransactionVerboseData: message.IncludeTransactionVerboseData,
	}
	return nil
}

func (x *KaspadMessage_GetBlockResponse) toAppMessage() (appmessage.Message, error) {
	var err *appmessage.RPCError
	if x.GetBlockResponse.Error != nil {
		err = &appmessage.RPCError{Message: x.GetBlockResponse.Error.Message}
	}
	var blockVerboseData *appmessage.BlockVerboseData
	if x.GetBlockResponse.BlockVerboseData != nil {
		appBlockVerboseData, err := x.GetBlockResponse.BlockVerboseData.toAppMessage()
		if err != nil {
			return nil, err
		}
		blockVerboseData = appBlockVerboseData
	}
	return &appmessage.GetBlockResponseMessage{
		BlockHex:         x.GetBlockResponse.BlockHex,
		BlockVerboseData: blockVerboseData,
		Error:            err,
	}, nil
}

func (x *KaspadMessage_GetBlockResponse) fromAppMessage(message *appmessage.GetBlockResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	var blockVerboseData *BlockVerboseData
	if message.BlockVerboseData != nil {
		protoBlockVerboseData := &BlockVerboseData{}
		err := protoBlockVerboseData.fromAppMessage(message.BlockVerboseData)
		if err != nil {
			return err
		}
		blockVerboseData = protoBlockVerboseData
	}
	x.GetBlockResponse = &GetBlockResponseMessage{
		BlockHex:         message.BlockHex,
		BlockVerboseData: blockVerboseData,
		Error:            err,
	}
	return nil
}

func (x *BlockVerboseData) toAppMessage() (*appmessage.BlockVerboseData, error) {
	transactionVerboseData := make([]*appmessage.TransactionVerboseData, len(x.TransactionVerboseData))
	for i, data := range x.TransactionVerboseData {
		inputs := make([]*appmessage.TransactionVerboseInput, len(data.TransactionVerboseInputs))
		for j, item := range data.TransactionVerboseInputs {
			scriptSig := &appmessage.ScriptSig{
				Asm: item.ScriptSig.Asm,
				Hex: item.ScriptSig.Hex,
			}
			inputs[j] = &appmessage.TransactionVerboseInput{
				TxID:        item.TxId,
				OutputIndex: item.OutputIndex,
				ScriptSig:   scriptSig,
				Sequence:    item.Sequence,
			}
		}
		outputs := make([]*appmessage.TransactionVerboseOutput, len(data.TransactionVerboseOutputs))
		for j, item := range data.TransactionVerboseOutputs {
			scriptPubKey := &appmessage.ScriptPubKeyResult{
				Asm:     item.ScriptPubKey.Asm,
				Hex:     item.ScriptPubKey.Hex,
				Type:    item.ScriptPubKey.Type,
				Address: item.ScriptPubKey.Address,
			}
			outputs[j] = &appmessage.TransactionVerboseOutput{
				Value:        item.Value,
				Index:        item.Index,
				ScriptPubKey: scriptPubKey,
			}
		}
		transactionVerboseData[i] = &appmessage.TransactionVerboseData{
			Hex:                       data.Hex,
			TxID:                      data.TxId,
			Hash:                      data.Hash,
			Size:                      data.Size,
			Version:                   data.Version,
			LockTime:                  data.LockTime,
			SubnetworkID:              data.SubnetworkId,
			Gas:                       data.Gas,
			PayloadHash:               data.PayloadHash,
			Payload:                   data.Payload,
			TransactionVerboseInputs:  inputs,
			TransactionVerboseOutputs: outputs,
			BlockHash:                 data.BlockHash,
			AcceptedBy:                data.AcceptedBy,
			IsInMempool:               data.IsInMempool,
			Time:                      data.Time,
			BlockTime:                 data.BlockTime,
		}
	}
	return &appmessage.BlockVerboseData{
		Hash:                   x.Hash,
		Confirmations:          x.Confirmations,
		Size:                   x.Size,
		BlueScore:              x.BlueScore,
		IsChainBlock:           x.IsChainBlock,
		Version:                x.Version,
		VersionHex:             x.VersionHex,
		HashMerkleRoot:         x.HashMerkleRoot,
		AcceptedIDMerkleRoot:   x.AcceptedIDMerkleRoot,
		UTXOCommitment:         x.UtxoCommitment,
		TxIDs:                  x.TransactionHex,
		TransactionVerboseData: transactionVerboseData,
		Time:                   x.Time,
		Nonce:                  x.Nonce,
		Bits:                   x.Bits,
		Difficulty:             x.Difficulty,
		ParentHashes:           x.ParentHashes,
		SelectedParentHash:     x.SelectedParentHash,
		ChildHashes:            x.ChildHashes,
		AcceptedBlockHashes:    x.AcceptedBlockHashes,
	}, nil
}

func (x *BlockVerboseData) fromAppMessage(message *appmessage.BlockVerboseData) error {
	transactionVerboseData := make([]*TransactionVerboseData, len(message.TransactionVerboseData))
	for i, data := range message.TransactionVerboseData {
		inputs := make([]*TransactionVerboseInput, len(data.TransactionVerboseInputs))
		for j, item := range data.TransactionVerboseInputs {
			scriptSig := &ScriptSig{
				Asm: item.ScriptSig.Asm,
				Hex: item.ScriptSig.Hex,
			}
			inputs[j] = &TransactionVerboseInput{
				TxId:        item.TxID,
				OutputIndex: item.OutputIndex,
				ScriptSig:   scriptSig,
				Sequence:    item.Sequence,
			}
		}
		outputs := make([]*TransactionVerboseOutput, len(data.TransactionVerboseOutputs))
		for j, item := range data.TransactionVerboseOutputs {
			scriptPubKey := &ScriptPubKeyResult{
				Asm:     item.ScriptPubKey.Asm,
				Hex:     item.ScriptPubKey.Hex,
				Type:    item.ScriptPubKey.Type,
				Address: item.ScriptPubKey.Address,
			}
			outputs[j] = &TransactionVerboseOutput{
				Value:        item.Value,
				Index:        item.Index,
				ScriptPubKey: scriptPubKey,
			}
		}
		transactionVerboseData[i] = &TransactionVerboseData{
			Hex:                       data.Hex,
			TxId:                      data.TxID,
			Hash:                      data.Hash,
			Size:                      data.Size,
			Version:                   data.Version,
			LockTime:                  data.LockTime,
			SubnetworkId:              data.SubnetworkID,
			Gas:                       data.Gas,
			PayloadHash:               data.PayloadHash,
			Payload:                   data.Payload,
			TransactionVerboseInputs:  inputs,
			TransactionVerboseOutputs: outputs,
			BlockHash:                 data.BlockHash,
			AcceptedBy:                data.AcceptedBy,
			IsInMempool:               data.IsInMempool,
			Time:                      data.Time,
			BlockTime:                 data.BlockTime,
		}
	}
	*x = BlockVerboseData{
		Hash:                   message.Hash,
		Confirmations:          message.Confirmations,
		Size:                   message.Size,
		BlueScore:              message.BlueScore,
		IsChainBlock:           message.IsChainBlock,
		Version:                message.Version,
		VersionHex:             message.VersionHex,
		HashMerkleRoot:         message.HashMerkleRoot,
		AcceptedIDMerkleRoot:   message.AcceptedIDMerkleRoot,
		UtxoCommitment:         message.UTXOCommitment,
		TransactionHex:         message.TxIDs,
		TransactionVerboseData: transactionVerboseData,
		Time:                   message.Time,
		Nonce:                  message.Nonce,
		Bits:                   message.Bits,
		Difficulty:             message.Difficulty,
		ParentHashes:           message.ParentHashes,
		SelectedParentHash:     message.SelectedParentHash,
		ChildHashes:            message.ChildHashes,
		AcceptedBlockHashes:    message.AcceptedBlockHashes,
	}
	return nil
}