package srrule

type SRruler interface {
	UnmarshalRule(data []byte) (Rules, error)
	IfInRange(r string, z string) (bool, error)
}
