package srrule

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

func checkRRule(r string, l *time.Location) (bool, error) {
	if strings.Contains(r, alwaysOn) {
		return true, nil
	}

	if strings.HasPrefix(r, ignorePrefix) {
		return false, nil
	} else if strings.HasPrefix(r, relPrefix) {
		inWeekDays := false
	Loop:
		for _, l := range strings.Split(strings.Trim(r, relPrefix), rulesSeparator) {
			if strings.HasPrefix(l, ignorePrefix) {
				continue
			}
			if strings.HasPrefix(l, wDayPrefix) {
				l = strings.TrimFunc(l, func(r rune) bool {
					return !unicode.IsNumber(r) && !unicode.IsMark(r)
				})

				for _, w := range strings.Split(l, ruleSeqSeparator) {
					if len(w) > 1 {
						st, err := strconv.Atoi(strings.Split(w, rangeSeparator)[0])
						if err != nil {
							return false, err
						}
						end, err := strconv.Atoi(strings.Split(w, rangeSeparator)[1])
						if err != nil {
							return false, err
						}

						for i := st; i <= end; i++ {
							if time.Now().Weekday() == time.Weekday(i) {
								inWeekDays = true
								break Loop
							}
						}
					} else {
						wd, err := strconv.Atoi(w)
						if err != nil {
							return false, err
						}
						test := time.Weekday(wd)
						if time.Now().Weekday() == test {
							inWeekDays = true
							break Loop
						}
					}
				}
			}
		}

		if !inWeekDays {
			return false, nil
		}
		for _, rul := range strings.Split(strings.Trim(r, relPrefix), rulesSeparator) {
			if strings.HasPrefix(rul, timePrefix) {
				rul = strings.TrimFunc(rul, func(r rune) bool {
					return !unicode.IsNumber(r) && !unicode.IsMark(r)
				})

				for _, t := range strings.Split(rul, ruleSeqSeparator) {
					if len(t) > 1 {
						start, err := time.ParseInLocation(timeLayout, strings.Split(t, rangeSeparator)[0], l)
						if err != nil {
							return false, err
						}
						end, err := time.ParseInLocation(timeLayout, strings.Split(t, rangeSeparator)[1], l)
						if err != nil {
							return false, err
						}
						nowTime := time.Date(0, 1, 1, time.Now().Hour(), time.Now().Minute(), 0, 0, l)
						if (nowTime.After(start) && nowTime.Before(end)) || nowTime.Equal(start) || nowTime.Equal(end) {
							return true, nil
						}
					} else {
						return false, FormatError
					}
				}
			}
		}
	} else if strings.HasPrefix(r, exactPrefix) {
		var kl []string
		for _, lk := range strings.Split(strings.Trim(r, relPrefix), rulesSeparator) {
			lk = strings.TrimFunc(lk, func(r rune) bool {
				return !unicode.IsNumber(r) && !unicode.IsMark(r)
			})
			kl = append(kl, lk)
		}

		if kl != nil && len(kl) > 1 {
			start, err := time.ParseInLocation(excTimeLayout, kl[0], l)
			if err != nil {
				return false, err
			}
			end, err := time.ParseInLocation(excTimeLayout, kl[1], l)
			if err != nil {
				return false, err
			}
			if (time.Now().After(start) && time.Now().Before(end)) || time.Now().Equal(start) || time.Now().Equal(end) {
				return true, nil
			}
		}

	} else {
		return false, FormatError
	}
	return false, nil
}
