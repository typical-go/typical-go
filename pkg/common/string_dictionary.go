package common

// StringDictionary is list of string key-value item
type StringDictionary []*StringKV

// StringKV contain key and value
type StringKV struct {
	Key   string
	Value string
}

// Add item
func (k *StringDictionary) Add(key, s string) *StringDictionary {
	*k = append(*k, &StringKV{Key: key, Value: s})
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
