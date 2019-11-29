package coll

// KeyStrings is collection of key-string pair
type KeyStrings []KeyString

// Append item
func (k *KeyStrings) Append(item ...KeyString) *KeyStrings {
	*k = append(*k, item...)
	return k
}

// Add item
func (k *KeyStrings) Add(key, s string) *KeyStrings {
	k.Append(KeyString{Key: key, String: s})
	return k
}
