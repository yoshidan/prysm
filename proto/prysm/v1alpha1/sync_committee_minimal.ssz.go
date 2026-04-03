//go:build minimal

package eth

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// Minimal preset constants for SyncCommittee
// SYNC_COMMITTEE_SIZE = 32 (vs 512 for mainnet)
// SyncCommitteeBits = 4 bytes (32 bits) vs 64 bytes (512 bits) for mainnet
// SyncCommittee SSZ size = 32*48 + 48 = 1584 bytes (vs 24624 for mainnet)
// SyncAggregate SSZ size = 4 + 96 = 100 bytes (vs 160 for mainnet)

// MarshalSSZ ssz marshals the SyncAggregate object
func (s *SyncAggregate) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncAggregate object to a target array
func (s *SyncAggregate) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SyncCommitteeBits' - 4 bytes for minimal (32 bits)
	if size := len(s.SyncCommitteeBits); size != 4 {
		err = ssz.ErrBytesLengthFn("--.SyncCommitteeBits", size, 4)
		return
	}
	dst = append(dst, s.SyncCommitteeBits...)

	// Field (1) 'SyncCommitteeSignature'
	if size := len(s.SyncCommitteeSignature); size != 96 {
		err = ssz.ErrBytesLengthFn("--.SyncCommitteeSignature", size, 96)
		return
	}
	dst = append(dst, s.SyncCommitteeSignature...)

	return
}

// UnmarshalSSZ ssz unmarshals the SyncAggregate object
func (s *SyncAggregate) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 100 {
		return ssz.ErrSize
	}

	// Field (0) 'SyncCommitteeBits' - 4 bytes for minimal
	if cap(s.SyncCommitteeBits) == 0 {
		s.SyncCommitteeBits = make([]byte, 0, len(buf[0:4]))
	}
	s.SyncCommitteeBits = append(s.SyncCommitteeBits, buf[0:4]...)

	// Field (1) 'SyncCommitteeSignature'
	if cap(s.SyncCommitteeSignature) == 0 {
		s.SyncCommitteeSignature = make([]byte, 0, len(buf[4:100]))
	}
	s.SyncCommitteeSignature = append(s.SyncCommitteeSignature, buf[4:100]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncAggregate object
func (s *SyncAggregate) SizeSSZ() (size int) {
	size = 100
	return
}

// HashTreeRoot ssz hashes the SyncAggregate object
func (s *SyncAggregate) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncAggregate object with a hasher
func (s *SyncAggregate) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'SyncCommitteeBits' - 4 bytes for minimal
	if size := len(s.SyncCommitteeBits); size != 4 {
		err = ssz.ErrBytesLengthFn("--.SyncCommitteeBits", size, 4)
		return
	}
	hh.PutBytes(s.SyncCommitteeBits)

	// Field (1) 'SyncCommitteeSignature'
	if size := len(s.SyncCommitteeSignature); size != 96 {
		err = ssz.ErrBytesLengthFn("--.SyncCommitteeSignature", size, 96)
		return
	}
	hh.PutBytes(s.SyncCommitteeSignature)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the SyncCommittee object
func (s *SyncCommittee) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncCommittee object to a target array
func (s *SyncCommittee) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Pubkeys' - 32 for minimal
	if size := len(s.Pubkeys); size != 32 {
		err = ssz.ErrVectorLengthFn("--.Pubkeys", size, 32)
		return
	}
	for ii := 0; ii < 32; ii++ {
		if size := len(s.Pubkeys[ii]); size != 48 {
			err = ssz.ErrBytesLengthFn("--.Pubkeys[ii]", size, 48)
			return
		}
		dst = append(dst, s.Pubkeys[ii]...)
	}

	// Field (1) 'AggregatePubkey'
	if size := len(s.AggregatePubkey); size != 48 {
		err = ssz.ErrBytesLengthFn("--.AggregatePubkey", size, 48)
		return
	}
	dst = append(dst, s.AggregatePubkey...)

	return
}

// UnmarshalSSZ ssz unmarshals the SyncCommittee object
func (s *SyncCommittee) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 1584 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkeys' - 32 for minimal, 32*48 = 1536 bytes
	s.Pubkeys = make([][]byte, 32)
	for ii := 0; ii < 32; ii++ {
		if cap(s.Pubkeys[ii]) == 0 {
			s.Pubkeys[ii] = make([]byte, 0, len(buf[0:1536][ii*48:(ii+1)*48]))
		}
		s.Pubkeys[ii] = append(s.Pubkeys[ii], buf[0:1536][ii*48:(ii+1)*48]...)
	}

	// Field (1) 'AggregatePubkey'
	if cap(s.AggregatePubkey) == 0 {
		s.AggregatePubkey = make([]byte, 0, len(buf[1536:1584]))
	}
	s.AggregatePubkey = append(s.AggregatePubkey, buf[1536:1584]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncCommittee object
func (s *SyncCommittee) SizeSSZ() (size int) {
	size = 1584
	return
}

// HashTreeRoot ssz hashes the SyncCommittee object
func (s *SyncCommittee) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncCommittee object with a hasher
func (s *SyncCommittee) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkeys' - 32 for minimal
	{
		if size := len(s.Pubkeys); size != 32 {
			err = ssz.ErrVectorLengthFn("--.Pubkeys", size, 32)
			return
		}
		subIndx := hh.Index()
		for _, i := range s.Pubkeys {
			if len(i) != 48 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}
		hh.Merkleize(subIndx)
	}

	// Field (1) 'AggregatePubkey'
	if size := len(s.AggregatePubkey); size != 48 {
		err = ssz.ErrBytesLengthFn("--.AggregatePubkey", size, 48)
		return
	}
	hh.PutBytes(s.AggregatePubkey)

	hh.Merkleize(indx)
	return
}
