package store

type StoreType string

const (
	StoreType_file StoreType = "file"
)

func (st StoreType) Validate() bool {
	switch st {
	case StoreType_file:
		return true
	default:
		return false
	}
}
