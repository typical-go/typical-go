package coll

// KeyStrings is collection of key-string pair
type KeyStrings []*KeyString

// Append item
func (k *KeyStrings) Append(item ...*KeyString) *KeyStrings {
	*k = append(*k, item...)
	return k
}

// Add item
func (k *KeyStrings) Add(key, s string) *KeyStrings {
	k.Append(&KeyString{Key: key, String: s})
	return k
}

// Get keystring by key
func (k *KeyStrings) Get(key string) *KeyString {
	for _, ks := range *k {
		if key == ks.Key {
			return ks
		}
	}
	return nil
}

// Exist return true if key exist
func (k *KeyStrings) Exist(key string) bool {
	return k.Get(key) == nil
}

// Slice of Key-String
func (k *KeyStrings) Slice() []*KeyString {
	return *k
}
