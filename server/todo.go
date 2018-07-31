package main

import (
	"strings"
)

type TODO struct {
	ID        string
	Name      string
	Completed bool
}

type ListFilter struct {
	Query     string
	Completed bool
	Limit     int64
	Offset    int64
}

type DataLayer interface {
	List(filter ListFilter) ([]*TODO, error)
}

type dataLayer struct {
	data []*TODO
}

func NewDataLayer() DataLayer {
	return &dataLayer{}
}

func (dl *dataLayer) List(filter ListFilter) ([]*TODO, error) {
	if dl.data == nil {
		return nil, nil
	}

	var filtered []*TODO

	for _, v := range dl.data {
		if v.Completed != filter.Completed {
			continue
		}

		if !strings.Contains(v.Name, filter.Query) {
			continue
		}

		filtered = append(filtered, v)
	}

	max := int64(len(filtered))

	if filter.Limit < 0 {
		filter.Limit = 0
	} else if filter.Limit > max {
		filter.Limit = max
	}

	if filter.Offset < 0 {
		filter.Offset = 0
	} else if filter.Offset+1 > max {
		filter.Offset = max
	}

	return filtered, nil
}
