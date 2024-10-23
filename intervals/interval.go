package intervals

import "sort"

type Interval struct {
	Start int64
	End   int64
}

func Merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sortedIntervals := make([]Interval, len(intervals))
	copy(sortedIntervals, intervals)
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i].Start < sortedIntervals[j].Start
	})
	mergedIntervals := []Interval{sortedIntervals[0]}
	for _, interval := range sortedIntervals[1:] {
		lastMergedInterval := &mergedIntervals[len(mergedIntervals)-1]
		if interval.Start <= lastMergedInterval.End {
			if interval.End > lastMergedInterval.End {
				lastMergedInterval.End = interval.End
			}
		} else {
			mergedIntervals = append(mergedIntervals, interval)
		}
	}
	return mergedIntervals
}

func IsOverlap(a, b Interval) bool {
	return a.Start <= b.End && b.Start <= a.End
}
