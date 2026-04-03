//go:build !minimal

package eth

import (
	"github.com/OffchainLabs/go-bitfield"
	"github.com/OffchainLabs/prysm/v7/consensus-types/primitives"
)

// ToAttestationElectra converts the attestation to an AttestationElectra.
func (a *SingleAttestation) ToAttestationElectra(committee []primitives.ValidatorIndex) *AttestationElectra {
	cb := primitives.NewAttestationCommitteeBits()
	cb.SetBitAt(uint64(a.CommitteeId), true)

	ab := bitfield.NewBitlist(uint64(len(committee)))
	for i, ix := range committee {
		if a.AttesterIndex == ix {
			ab.SetBitAt(uint64(i), true)
			break
		}
	}

	return &AttestationElectra{
		AggregationBits: ab,
		Data:            a.Data,
		Signature:       a.Signature,
		CommitteeBits:   cb,
	}
}
