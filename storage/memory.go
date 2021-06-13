package storage

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/garupanojisan/mredis/storage/sortedset"
)

type Memory struct {
	values map[string]interface{}
}

var (
	ErrMissingKey = errors.New("error missing key")
)

var sharedInstance *Memory

func GetInstance() *Memory {
	if sharedInstance == nil {
		sharedInstance = &Memory{
			values: make(map[string]interface{}),
		}
	}
	return sharedInstance
}

func (m *Memory) Set(ctx context.Context, key string, value []byte) error {
	m.values[key] = value
	return nil
}

func (m *Memory) Get(ctx context.Context, key string) ([]byte, error) {
	if v, ok := m.values[key]; ok {
		if b, ok := v.([]byte); ok {
			return b, nil
		}
		return nil, errors.New("unknown error: invalid value type")
	}
	return nil, fmt.Errorf("%w: key = %v", ErrMissingKey, key)
}

func (m *Memory) Push(ctx context.Context, key string, value []byte) error {
	if stack, ok := m.values[key]; ok {
		if bs, ok := stack.([][]byte); ok {
			bs = append(bs, value)
			m.values[key] = bs
			return nil
		}
		return fmt.Errorf("unknown value type: key = %s", key)
	}
	m.values[key] = [][]byte{value}
	return nil
}

func (m *Memory) Pop(ctx context.Context, key string) ([]byte, error) {
	if stack, ok := m.values[key]; ok {
		if bs, ok := stack.([][]byte); ok {
			if len(bs) == 0 {
				return nil, nil
			}
			b := bs[len(bs)-1]
			bs = bs[:len(bs)-1]
			m.values[key] = bs
			return b, nil
		}
		return nil, fmt.Errorf("unknown value type: key = %s", key)
	}
	return nil, fmt.Errorf("%w: key = %v", ErrMissingKey, key)
}

func (m *Memory) Enqueue(ctx context.Context, key string, value []byte) error {
	if stack, ok := m.values[key]; ok {
		if bs, ok := stack.([][]byte); ok {
			bs = append(bs, value)
			m.values[key] = bs
			return nil
		}
		return fmt.Errorf("unknown value type: key = %s", key)
	}
	m.values[key] = [][]byte{value}
	return nil
}

func (m *Memory) Dequeue(ctx context.Context, key string) ([]byte, error) {
	if stack, ok := m.values[key]; ok {
		if bs, ok := stack.([][]byte); ok {
			if len(bs) == 0 {
				return nil, nil
			}
			b := bs[0]
			bs = bs[1:]
			m.values[key] = bs
			return b, nil
		}
		return nil, fmt.Errorf("unknown value type: key = %s", key)
	}
	return nil, fmt.Errorf("%w: key = %v", ErrMissingKey, key)
}

func (m *Memory) Keys(ctx context.Context, pattern string) ([]string, error) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	var keys []string
	for key := range m.values {
		if regex.MatchString(key) {
			keys = append(keys, key)
		}
	}
	return keys, nil
}

func (m *Memory) ZAdd(ctx context.Context, key string, score int64, member string) error {
	if v, ok := m.values[key]; ok {
		if s, ok := v.(*sortedset.SortedSet); ok {
			if err := s.Add(ctx, score, member); err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("unknown value type: key = %s, member = %s, score = %d", key, member, score)
	}

	s := sortedset.NewSortedSet()
	if err := s.Add(ctx, score, member); err != nil {
		return err
	}
	m.values[key] = s

	return nil
}

func (m *Memory) ZRank(ctx context.Context, key, member string) (int64, error) {
	if v, ok := m.values[key]; ok {
		if s, ok := v.(*sortedset.SortedSet); ok {
			return s.RankAsc(ctx, member)
		}
		return 0, fmt.Errorf("unknown value type: key = %s", key)
	}
	return 0, fmt.Errorf("%w: key = %v", ErrMissingKey, key)
}
