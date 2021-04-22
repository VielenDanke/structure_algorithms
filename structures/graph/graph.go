package graph

import (
	"github.com/vielendanke/structure_algorithms/structures/common"
)

type adjacencyGraph struct {
	m map[string][]string
}

func NewGraph() common.Graph {
	return &adjacencyGraph{m: make(map[string][]string)}
}

func (a *adjacencyGraph) AddEdge(fVertex string, sVertex string) {
	f, fOk := a.m[fVertex]
	s, sOk := a.m[sVertex]
	if fOk {
		f = append(f, sVertex)
		a.m[fVertex] = f
	} else {
		a.m[fVertex] = append(make([]string, 0), sVertex)
	}
	if sOk {
		s = append(s, fVertex)
		a.m[sVertex] = s
	} else {
		a.m[sVertex] = append(make([]string, 0), fVertex)
	}
}

func (a *adjacencyGraph) AddVertex(val string) {
	_, ok := a.m[val]
	if ok {
		return
	}
	a.m[val] = make([]string, 0)
}

func (a *adjacencyGraph) RemoveEdge(fVertex string, sVertex string) {
	f, fOk := a.m[fVertex]
	s, sOk := a.m[sVertex]
	if fOk {
		for k, v := range f {
			if v == sVertex {
				a.m[fVertex] = append(f[:k], f[k+1:]...)
			}
		}
	}
	if sOk {
		for k, v := range s {
			if v == fVertex {
				a.m[sVertex] = append(s[:k], s[k+1:]...)
			}
		}
	}
}

func (a *adjacencyGraph) RemoveVertex(fVertex string) bool {
	fArr, ok := a.m[fVertex]
	if !ok {
		return false
	}
	for _, v := range fArr {
		temp, tOk := a.m[v]
		if tOk {
			for k, tv := range temp {
				if tv == fVertex {
					a.m[v] = append(temp[:k], temp[k+1:]...)
				}
			}
		}
	}
	delete(a.m, fVertex)
	return true
}