package types

import "github.com/cometbft/cometbft/crypto/tmhash"

func (m *VirtualBlock) Hash() (dAtA []byte, err error) {
	data, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	return tmhash.Sum(data), nil
}
