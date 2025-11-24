package xinterval

import "sort"

type Interval struct {
	Start int64
	End   int64
}

func Merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sortedIntervals := make([]Interval, len(intervals))
	copy(sortedIntervals, intervals)
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i].Start < sortedIntervals[j].Start
	})

	merged := []Interval{sortedIntervals[0]}

	for _, interval := range sortedIntervals[1:] {
		lastMerged := &merged[len(merged)-1]

		if IsOverlap(*lastMerged, interval) {
			lastMerged.End = max(lastMerged.End, interval.End)
		} else {
			merged = append(merged, interval)
		}
	}

	return merged
}

func IsOverlap(a, b Interval) bool {
	return a.Start <= b.End && b.Start <= a.End
}
