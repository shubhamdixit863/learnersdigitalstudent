package versiontrack

//stack for version tracking
type VersionTrack struct {
	items []interface{}
}

func NewVersionTrack() *VersionTrack {
	return &VersionTrack{}
}

func (vt *VersionTrack) Push(item interface{}) {
	vt.items = append(vt.items, item)
}

func (vt *VersionTrack) Pop() (interface{}, error) {
	if len(vt.items) == 0 {
		return nil, nil
	}
	vt.items = vt.items[:len(vt.items)-1]
	return vt.items[len(vt.items)-1], nil
}
func (vt *VersionTrack) Peek() interface{} {
	return vt.items[len(vt.items)-1]
}

func (vt *VersionTrack) IsEmpty() bool {
	return len(vt.items) == 0
}

func (vt *VersionTrack) Clear() {
	vt.items = nil
}
