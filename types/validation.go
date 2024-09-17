package types

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	tmcrypto "github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/dymensionxyz/dymint/fraud"
)

var MaxDrift = 10 * time.Minute

type ErrTimeFraud struct {
	drift           time.Duration
	proposerAddress []byte
	headerHash      [32]byte
	headerHeight    uint64
	headerTime      time.Time
	currentTime     time.Time
}

func NewErrTimeFraud(block *Block, currentTime time.Time) error {
	drift := time.Unix(int64(block.Header.Time), 0).Sub(currentTime)

	return ErrTimeFraud{
		drift:           drift,
		proposerAddress: block.Header.ProposerAddress,
		headerHash:      block.Header.Hash(),
		headerHeight:    block.Header.Height,
		headerTime:      time.Unix(int64(block.Header.Time), 0),
		currentTime:     currentTime,
	}
}

func (e ErrTimeFraud) Error() string {
	return fmt.Sprintf(
		`Sequencer %X posted a block with header hash %X, at height %dwith a time drift of %s, 
when the maximum allowed limit is %s. Sequencer reported block time was %s while the node local time is %s`,
		e.proposerAddress, e.headerHash, e.headerHeight, e.drift, MaxDrift, e.headerTime, e.currentTime,
	)
}

func (e ErrTimeFraud) Unwrap() error {
	return fraud.ErrFraud
}

func ValidateProposedTransition(state *State, block *Block, commit *Commit, proposerPubKey tmcrypto.PubKey) error {
	if err := block.ValidateWithState(state); err != nil {
		return fmt.Errorf("block: %w", err)
	}

	if err := commit.ValidateWithHeader(proposerPubKey, &block.Header); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}

// ValidateBasic performs basic validation of a block.
func (b *Block) ValidateBasic() error {
	currentTime := time.Now().UTC()
	if currentTime.Add(MaxDrift).Before(time.Unix(0, int64(b.Header.Time))) {
		return NewErrTimeFraud(b, currentTime)
	}

	err := b.Header.ValidateBasic()
	if err != nil {
		return err
	}

	err = b.Data.ValidateBasic()
	if err != nil {
		return err
	}

	err = b.LastCommit.ValidateBasic()
	if err != nil {
		return err
	}

	if b.Header.DataHash != [32]byte(GetDataHash(b)) {
		return ErrInvalidHeaderDataHash
	}
	return nil
}

func (b *Block) ValidateWithState(state *State) error {
	err := b.ValidateBasic()
	if err != nil {
		return err
	}
	if b.Header.Version.App != state.Version.Consensus.App ||
		b.Header.Version.Block != state.Version.Consensus.Block {
		return errors.New("b version mismatch")
	}

	if b.Header.Height != state.NextHeight() {
		return NewErrFraudHeightMismatch(state.NextHeight(), b.Header.Height, b)
	}

	if !bytes.Equal(b.Header.AppHash[:], state.AppHash[:]) {
		return NewErrFraudAppHashMismatch(state.AppHash, b.Header.AppHash, b)
	}

	if !bytes.Equal(b.Header.LastResultsHash[:], state.LastResultsHash[:]) {
		return NewErrLastResultsHashMismatch(state.LastResultsHash, b)
	}

	return nil
}

// ValidateBasic performs basic validation of a header.
func (h *Header) ValidateBasic() error {
	if len(h.ProposerAddress) == 0 {
		return errors.New("no proposer address")
	}

	return nil
}

// ValidateBasic performs basic validation of block data.
// Actually it's a placeholder, because nothing is checked.
func (d *Data) ValidateBasic() error {
	return nil
}

// ValidateBasic performs basic validation of a commit.
func (c *Commit) ValidateBasic() error {
	if c.Height > 0 {
		if len(c.Signatures) != 1 {
			return errors.New("there should be 1 signature")
		}
		if len(c.Signatures[0]) > tmtypes.MaxSignatureSize {
			return errors.New("signature is too big")
		}
	}
	return nil
}

func (c *Commit) ValidateWithHeader(proposerPubKey tmcrypto.PubKey, header *Header) error {
	if err := c.ValidateBasic(); err != nil {
		return err
	}
	abciHeaderPb := ToABCIHeaderPB(header)
	abciHeaderBytes, err := abciHeaderPb.Marshal()
	if err != nil {
		return err
	}
	// commit is validated to have single signature
	if !proposerPubKey.VerifySignature(abciHeaderBytes, c.Signatures[0]) {
		return ErrInvalidSignature
	}
	return nil
}
