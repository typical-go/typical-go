package common

// StringDictionary is list of string key-value item
type StringDictionary []*StringKV

// StringKV contain key and value
type StringKV struct {
	Key   string
	Value string
}

// Append item
func (k *StringDictionary) Append(item ...*StringKV) *StringDictionary {
	*k = append(*k, item...)
	return k
}

// Add item
func (k *StringDictionary) Add(key, s string) *StringDictionary {
	k.Append(&StringKV{Key: key, Value: s})
	return k
}

// Get keystring by key
func (k *StringDictionary) Get(key string) *StringKV {
	for _, kv := range *k {
		if key == kv.Key {
			return kv
		}
	}
	return nil
}

// Exist return true if key exist
func (k *StringDictionary) Exist(key string) bool {
	return k.Get(key) == nil
}

// Slice of Key-String
func (k *StringDictionary) Slice() []*StringKV {
	return *k
}
