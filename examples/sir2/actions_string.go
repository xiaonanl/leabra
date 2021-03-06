// Code generated by "stringer -type=Actions"; DO NOT EDIT.

package main

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Store-0]
	_ = x[Ignore-1]
	_ = x[Recall-2]
	_ = x[ActionsN-3]
}

const _Actions_name = "StoreIgnoreRecallActionsN"

var _Actions_index = [...]uint8{0, 5, 11, 17, 25}

func (i Actions) String() string {
	if i < 0 || i >= Actions(len(_Actions_index)-1) {
		return "Actions(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Actions_name[_Actions_index[i]:_Actions_index[i+1]]
}

func (i *Actions) FromString(s string) error {
	for j := 0; j < len(_Actions_index)-1; j++ {
		if s == _Actions_name[_Actions_index[j]:_Actions_index[j+1]] {
			*i = Actions(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Actions")
}
