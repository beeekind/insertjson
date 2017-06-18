package insertjson

import (
	"bytes"
)

// Property inserts an unescaped key-value pair into a byte encoded
// JSON object as the first property of the object.
func Property(key string, value string, dst []byte) []byte {
	toInsert := []byte(`"` + key + `":"` + value + `",`)

	// build a slice large enough to contain both the original JSON object and
	// the new JSON property we are about to insert
	result := make([]byte, len(dst) + len(toInsert))

	// iterate over the result slice until we reach the opening character of
	// a JSON object: '{' . Then insert all elements of toInsert before continuing the
	// iteration, finally

	// n represents an index of dst
	n := 0

	// inserted represents whether or not we have already inserted toInsert, so that we
	// may handle reaching a second '{' character in the object.
	inserted := false

	// i represents an index of result
	for i := 0; i < len(result); i++ {
		result[i] = dst[n]
		n += 1

		// if we have just entered the JSON, insert desired pairing
		if  i < len(dst) && bytes.Compare([]byte{ dst[i] }, []byte("{")) == 0 && !inserted {
			for _, elem := range toInsert {
				i += 1
				result[i] = elem
			}

			inserted = true
		}
	}

	return result
}