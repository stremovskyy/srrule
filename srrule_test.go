/*
 * Project: go-estimator-ms
 * Module: estimator
 * File: srrule_test.go
 *
 * Copyright (C) Megakit Kharkiv 2017-2019, Inc - All Rights Reserved
 *  @link https://www.megakit.pro
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Anton (karmadon) Stremovskyy <stremovskyy@gmail.com>
 */
//["~w{1-5};t{10:12-17:35}~"]
package srrule

import (
	"testing"
	"time"
)

func BenchmarkIfInRange(b *testing.B) {
	r := NewSRrule()
	t := time.Now()

	for i := 0; i < b.N; i++ {
		_, _ = r.IfInRange("[\"~w{3};t{00:12-23:59}~\"]", t, "Europe/Kiev")
	}
}

func BenchmarkCheckRRule(b *testing.B) {
	l, _ := time.LoadLocation("Europe/Kiev")
	t := time.Now()

	for i := 0; i < b.N; i++ {
		_, _ = checkRRuleForTime("[\"~w{3};t{00:12-23:59}~\"]", &t, l)
	}
}

func BenchmarkRangeInLoc(b *testing.B) {
	r := NewSRrule()
	l, _ := time.LoadLocation("Europe/Kiev")
	t := time.Now()

	for i := 0; i < b.N; i++ {
		_, _ = r.RangeInLoc("[\"~w{3};t{00:12-23:59}~\"]", &t, l)
	}
}

func BenchmarkIfRulesInRangeLoc(b *testing.B) {
	r := NewSRrule()
	l, _ := time.LoadLocation("Europe/Kiev")
	rs, _ := r.UnmarshalRule([]byte("[\"~w{3};t{00:12-23:59}~\"]"))
	t := time.Now()

	for i := 0; i < b.N; i++ {
		_, _ = r.IfRulesInRangeLoc(&rs, &t, l)
	}
}

func TestIfInRange(t *testing.T) {
	type args struct {
		r string
		z string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "Sun,Mon,Sat whole day", args: args{"[\"~w{1-2,6};t{00:12-23:59}~\"]", "Europe/Kiev"}, want: true, wantErr: false},
		{name: "Always on", args: args{"[\"*\"]", "Europe/Kiev"}, want: true, wantErr: false},
		{name: "Wed whole day", args: args{"[\"~w{3};t{00:12-23:59}~\"]", "Europe/Kiev"}, want: false, wantErr: false},
		{name: "Office is Open", args: args{"[\"~w{1-5};t{10:12-17:35}~\"]", "Europe/Kiev"}, want: true, wantErr: false},
		{name: "Error in RRULE", args: args{"[\"\"~w{3};t{10:59-23:59}~\",\"-БЛА-БЛА-БЛА-\",\"~w{5};t{00:00-05:00}~\"]", "Europe/Kiev"}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewSRrule()
			tim := time.Now()

			got, err := r.IfInRange(tt.args.r, tim, tt.args.z)
			if (err != nil) != tt.wantErr {
				t.Errorf("IfInRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("IfInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkRRule(t *testing.T) {
	type args struct {
		r string
		l *time.Location
	}
	var tests []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tim := time.Now()

			got, err := checkRRuleForTime(tt.args.r, &tim, tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkRRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkRRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
