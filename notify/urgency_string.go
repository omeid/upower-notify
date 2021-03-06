// generated by stringer -type=Urgency; DO NOT EDIT

package notify

import "fmt"

const _Urgency_name = "LowNormalCritical"

var _Urgency_index = [...]uint8{0, 3, 9, 17}

func (i Urgency) String() string {
	if i+1 >= Urgency(len(_Urgency_index)) {
		return fmt.Sprintf("Urgency(%d)", i)
	}
	return _Urgency_name[_Urgency_index[i]:_Urgency_index[i+1]]
}
