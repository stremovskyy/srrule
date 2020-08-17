/*
 * Project: go-estimator-ms
 * Module: estimator
 * File: srrule.go
 *
 * Copyright (C) Megakit Kharkiv 2017-2019, Inc - All Rights Reserved
 *  @link https://www.megakit.pro
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Anton (karmadon) Stremovskyy <stremovskyy@gmail.com>
 */

package srrule

import (
	"encoding/json"
	"time"
)

type Rules []string

type SRrule struct {
}

func (s *SRrule) RangeInLoc(r string, loc *time.Location) (bool, error) {
	rs, err := s.UnmarshalRule([]byte(r))
	if err != nil {
		return false, err
	}

	return s.IfRulesInRangeLoc(&rs, loc)
}

func NewSRrule() *SRrule {
	return &SRrule{}
}

func (s *SRrule) UnmarshalRule(data []byte) (Rules, error) {
	var r Rules
	err := json.Unmarshal(data, &r)
	return r, err
}

func (s *SRrule) IfInRange(r string, z string) (bool, error) {
	l, err := time.LoadLocation(z)
	if err != nil {
		return false, err
	}

	return s.RangeInLoc(r, l)
}

func (s *SRrule) IfRulesInRangeLoc(r *Rules, loc *time.Location) (bool, error) {
	if r == nil {
		return false, nil
	}

	if len(*r) > 1 {
		for _, rl := range *r {
			in, err := checkRRule(rl, loc)
			if err != nil {
				return false, err
			}
			if in {
				return true, nil
			}
		}
		return false, nil
	} else if len(*r) == 1 {
		return checkRRule((*r)[0], loc)
	} else {
		return false, FormatError
	}
}
