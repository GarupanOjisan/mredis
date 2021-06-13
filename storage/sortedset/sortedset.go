package sortedset

import (
	"context"
	"fmt"

	"github.com/MauriceGit/skiplist"
)

type SortedSet struct {
	scoreRanking skiplist.SkipList
	memberScore  map[string]int64
}

type scoreRankingElement struct {
	member string
	score  int64
}

func (e *scoreRankingElement) ExtractKey() float64 {
	return float64(e.score)
}

func (e *scoreRankingElement) String() string {
	return e.member
}

func NewSortedSet() *SortedSet {
	return &SortedSet{
		scoreRanking: skiplist.New(),
		memberScore:  make(map[string]int64),
	}
}

func (s *SortedSet) Add(ctx context.Context, score int64, member string) error {
	if _, ok := s.memberScore[member]; !ok {
		s.memberScore[member] = score
		newElem := &scoreRankingElement{score: score, member: member}
		s.scoreRanking.Insert(newElem)
		return nil
	}

	s.memberScore[member] = score
	currentScore := s.memberScore[member]
	currentElem := &scoreRankingElement{score: currentScore, member: member}
	if e, ok := s.scoreRanking.Find(currentElem); ok {
		newElem := &scoreRankingElement{score: score, member: member}
		if s.scoreRanking.ChangeValue(e, newElem) {
			return nil
		}
		return fmt.Errorf("error update: score = %d, member = %s", score, member)
	}
	return nil
}

func (s *SortedSet) RankAsc(ctx context.Context, member string) (int64, error) {
	score, ok := s.memberScore[member]
	if !ok {
		return 0, fmt.Errorf("member not found(1): %s", member)
	}

	if e, ok := s.scoreRanking.Find(&scoreRankingElement{score: score, member: member}); ok {
		rank := int64(1)
		for node := s.scoreRanking.GetSmallestNode(); node != e; node = s.scoreRanking.Next(node) {
			rank++
		}
		return rank, nil
	}
	return 0, fmt.Errorf("member not found(2): member = %s, score = %d", member, score)
}
