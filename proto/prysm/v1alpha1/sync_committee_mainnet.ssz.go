//go:build !minimal

package eth

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the SyncAggregate object
func (s *SyncAggregate) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncAggregate object to a target array
func (s *SyncAggregate) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SyncCommitteeBits'
	if size := len(s.SyncCommitteeBits); size != 64 {
		err = ssz.ErrBytesLengthFn("--.SyncCommitteeBits", size, 64)
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
	if size != 160 {
		return ssz.ErrSize
	}

	// Field (0) 'SyncCommitteeBits'
	if cap(s.SyncCommitteeBits) == 0 {
		s.SyncCommitteeBits = make([]byte, 0, len(buf[0:64]))
	}
	s.SyncCommitteeBits = append(s.SyncCommitteeBits, buf[0:64]...)

	// Field (1) 'SyncCommitteeSignature'
	if cap(s.SyncCommitteeSignature) == 0 {
		s.SyncCommitteeSignature = make([]byte, 0, len(buf[64:160]))
	}
	s.SyncCommitteeSignature = append(s.SyncCommitteeSignature, buf[64:160]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncAggregate object
func (s *SyncAggregate) SizeSSZ() (size int) {
	size = 160
	return
}

// HashTreeRoot ssz hashes the SyncAggregate object
func (s *SyncAggregate) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncAggregate object with a hasher
func (s *SyncAggregate) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'SyncCommitteeBits'
	if size := len(s.SyncCommitteeBits); size != 64 {
		err = ssz.ErrBytesLengthFn("--.SyncCommitteeBits", size, 64)
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

	// Field (0) 'Pubkeys'
	if size := len(s.Pubkeys); size != 512 {
		err = ssz.ErrVectorLengthFn("--.Pubkeys", size, 512)
		return
	}
	for ii := 0; ii < 512; ii++ {
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
	if size != 24624 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkeys'
	s.Pubkeys = make([][]byte, 512)
	for ii := 0; ii < 512; ii++ {
		if cap(s.Pubkeys[ii]) == 0 {
			s.Pubkeys[ii] = make([]byte, 0, len(buf[0:24576][ii*48:(ii+1)*48]))
		}
		s.Pubkeys[ii] = append(s.Pubkeys[ii], buf[0:24576][ii*48:(ii+1)*48]...)
	}

	// Field (1) 'AggregatePubkey'
	if cap(s.AggregatePubkey) == 0 {
		s.AggregatePubkey = make([]byte, 0, len(buf[24576:24624]))
	}
	s.AggregatePubkey = append(s.AggregatePubkey, buf[24576:24624]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncCommittee object
func (s *SyncCommittee) SizeSSZ() (size int) {
	size = 24624
	return
}

// HashTreeRoot ssz hashes the SyncCommittee object
func (s *SyncCommittee) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncCommittee object with a hasher
func (s *SyncCommittee) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkeys'
	{
		if size := len(s.Pubkeys); size != 512 {
			err = ssz.ErrVectorLengthFn("--.Pubkeys", size, 512)
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
