package srrule

import "time"

type SRruler interface {
	UnmarshalRule(data []byte) (Rules, error)
	IfInRange(r string, z string) (bool, error)
	RangeInLoc(r string, loc *time.Location) (bool, error)
	IfRulesInRangeLoc(r *Rules, loc *time.Location) (bool, error)
}
