// Code generated by "stringer -type=SignatureType"; DO NOT EDIT.

package transaction

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SignatureTypeSingle-1]
	_ = x[SignatureTypeMulti-2]
}

const _SignatureType_name = "SignatureTypeSingleSignatureTypeMulti"

var _SignatureType_index = [...]uint8{0, 19, 37}

func (i SignatureType) String() string {
	i -= 1
	if i >= SignatureType(len(_SignatureType_index)-1) {
		return "SignatureType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SignatureType_name[_SignatureType_index[i]:_SignatureType_index[i+1]]
}
