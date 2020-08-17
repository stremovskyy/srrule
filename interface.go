package srrule

import "time"

type SRruler interface {
	UnmarshalRule(data []byte) (Rules, error)
	IfInRange(r string, t time.Time, z string) (bool, error)
	RangeInLoc(r string, t *time.Time, loc *time.Location) (bool, error)
	IfRulesInRangeLoc(r *Rules, t *time.Time, loc *time.Location) (bool, error)
}
