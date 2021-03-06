package main

import (
	"fmt"
	"github.com/satori/go.uuid"
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
	Create(todo *TODO) error
	Update(todo *TODO) error
	Delete(id string) error
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

		if len(filter.Query) != 0 && !strings.Contains(v.Name, filter.Query) {
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
		return nil, nil
	}

	return filtered[filter.Offset:filter.Limit], nil
}

func (dl *dataLayer) Create(t *TODO) error {
	t.ID = uuid.NewV4().String()

	dl.data = append(dl.data, t)

	return nil
}

func (dl *dataLayer) Update(t *TODO) error {
	var f *TODO

	for _, v := range dl.data {
		if v.ID == t.ID {
			f = v
			break
		}
	}

	if f == nil {
		return fmt.Errorf("no TODO found with ID %s", t.ID)
	}

	f.Completed = t.Completed
	t = f

	return nil
}

func (dl *dataLayer) Delete(id string) error {
	var f *TODO
	var ci int

	for i, v := range dl.data {
		if v.ID == id {
			ci = i
			f = v
			break
		}
	}

	if f == nil {
		return fmt.Errorf("no TODO found with ID %s", id)
	}

	ln := len(dl.data)

	if ci+1 > ln {
		dl.data = dl.data[ln-1:]
	} else {
		dl.data = append(dl.data[:ci], dl.data[ci+1:]...)
	}

	return nil
}
