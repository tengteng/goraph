package gm

import "testing"

func Test_JSON_GetVertices(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	l := g.GetVerticesSize()
	if l != 8 {
		test.Error("In testgraph1, it should have 8 vertices but", l)
	}
}

func Test_JSON_GetVerticesSize(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	r := g.GetVerticesSize()
	if r != 8 {
		test.Error("In testgraph1, it should have 8 vertices but", r)
	}
}

func Test_JSON_GetEdges(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	edges := g.GetEdges()
	check := false
	for _, edge := range edges {
		if edge.Weight == 200 {
			check = true
		}
	}
	if !check {
		test.Error("In testgraph1, there must be an edge of weight 200")
	}
}

func Test_JSON_GetEdgesSize(test *testing.T) {
	g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
	r1 := g1.GetEdgesSize()
	if r1 != 30 {
		test.Error("In testgraph1, it should have 30 edges but", r1)
	}
	g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
	r2 := g2.GetEdgesSize()
	if r2 != 24 {
		test.Error("In testgraph2, it should have 24 edges but", r2)
	}
}

func Test_JSON_CreateAndAddToGraph(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	_ = g.CreateAndAddToGraph("X")
	s := g.GetVerticesSize()
	if s != 9 {
		test.Error("In testgraph1, Created X vertex so it should now contain 9 vertices but", s)
	}
}

func Test_JSON_AddVertex(test *testing.T) {
	g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
	a := NewVertex("Google")
	g1.AddVertex(a)
	r1 := g1.GetVerticesSize()
	if r1 != 9 {
		test.Error("In testgraph1, it should have 9 vertices but", r1)
	}
}

func Test_JSON_AddEdge(test *testing.T) {
	g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
	e := NewEdge(g1.Vertices["A"], g1.Vertices["B"], 1.0)
	g1.AddEdge(e)
	r1 := g1.GetEdgesSize()
	if r1 != 31 {
		test.Error("In testgraph1, it should have 31 edges but", r1)
	}
}

func Test_JSON_GetOutVertices(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	d := g.Vertices["D"]
	l := d.GetOutVerticesSize()
	if l != 5 {
		test.Error("In testgraph1, D should have 5 outgoing vertices but", l)
	}

	_, ok1 := d.GetOutVertices()["F"]
	_, ok2 := d.GetOutVertices()["B"]
	if !ok1 || !ok2 {
		test.Error("In testgraph1, F and B should exist as outgoing vertices of D but", d.GetOutVertices()["F"], d.GetOutVertices()["B"])
	}

	g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
	testCases := []struct {
		vtx      string
		outedges int
	}{
		{"S", 3},
		{"A", 4},
		{"B", 4},
		{"C", 2},
		{"D", 0},
		{"E", 4},
		{"F", 3},
		{"T", 4},
	}

	for _, testCase := range testCases {
		v := g2.Vertices[testCase.vtx]
		n := v.GetOutVerticesSize()
		if n != testCase.outedges {
			test.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.outedges, n)
		}
	}
}

func Test_JSON_GetInVertices(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	d := g.Vertices["D"]
	l := d.GetInVerticesSize()
	if l != 5 {
		test.Error("In testgraph1, D should have 5 outgoing edges but", l)
	}

	s := g.Vertices["S"]
	se := s.GetInVerticesSize()
	if se != 3 {
		test.Error("In testgraph1, S only have 3 incoming vertices but", se)
	}

	g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
	testCases := []struct {
		vtx     string
		inedges int
	}{
		{"S", 3},
		{"A", 3},
		{"B", 3},
		{"C", 1},
		{"D", 5},
		{"E", 4},
		{"F", 2},
		{"T", 3},
	}

	for _, testCase := range testCases {
		v := g2.Vertices[testCase.vtx]
		n := v.GetInVerticesSize()
		if n != testCase.inedges {
			test.Errorf("In testgraph2, %+v, Expected '%#v'. But %#v", testCase.vtx, testCase.inedges, n)
		}
	}
}

func Test_JSON_ImmediateDominate(test *testing.T) {
	testCases2 := []struct {
		vts []string
		imd bool
	}{
		{[]string{"S", "A"}, true},
		{[]string{"A", "S"}, true},
		{[]string{"A", "T"}, true},
		{[]string{"D", "E"}, false},
		{[]string{"E", "D"}, true},
		{[]string{"F", "S"}, false},
		{[]string{"B", "C"}, false},
		{[]string{"D", "A"}, false},
		{[]string{"E", "F"}, true},
		{[]string{"D", "A"}, false},
		{[]string{"T", "C"}, false},
	}
	for _, testCase := range testCases2 {
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		r := g2.ImmediateDominate(g2.Vertices[testCase.vts[0]], g2.Vertices[testCase.vts[1]])
		if r != testCase.imd {
			test.Errorf("testgraph2: %+v does not go to %#v with one edge", testCase.vts[0], testCase.vts[1])
		}
	}
}

func Test_JSON_GetEdgeWeight(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 14.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},
		{[]string{"T", "A"}, 44.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"C", "E"}, 24.0},
		{[]string{"B", "E"}, 18.0},
		{[]string{"D", "T"}, 16.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},
		{[]string{"S", "C"}, 200.0},
		{[]string{"S", "A"}, 100.0},
	}
	for _, testCase := range testCases {
		wgt := g.GetEdgeWeight(g.Vertices[testCase.vertices[0]], g.Vertices[testCase.vertices[1]])
		if wgt != testCase.weight {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, wgt)
		}
	}

	g3 := FromJSON("../../files/testgraph.json", "testgraph.003")
	testCases3 := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 20.0}, // Updated (Added 6)
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},
		{[]string{"T", "A"}, 44.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"C", "E"}, 24.0},
		{[]string{"B", "E"}, 18.0},
		{[]string{"D", "T"}, 16.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},
		{[]string{"S", "C"}, 200.0},
		{[]string{"S", "A"}, 100.0},
	}
	for _, testCase := range testCases3 {
		wgt := g3.GetEdgeWeight(g3.Vertices[testCase.vertices[0]], g3.Vertices[testCase.vertices[1]])
		if wgt != testCase.weight {
			test.Errorf("In testgraph3, Expected '%#v'. But %#v", testCase.weight, wgt)
		}
	}
}

func Test_JSON_UpdateWeight(test *testing.T) {
	g := FromJSON("../../files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 914.0},
		{[]string{"A", "B"}, 95.0},
		{[]string{"A", "D"}, 920.0},
		{[]string{"A", "T"}, 944.0},
		{[]string{"T", "A"}, 944.0},
		{[]string{"D", "E"}, 92.0},
		{[]string{"E", "D"}, 92.0},
		{[]string{"C", "E"}, 924.0},
		{[]string{"B", "E"}, 918.0},
		{[]string{"D", "T"}, 916.0},
		{[]string{"T", "D"}, 916.0},
		{[]string{"F", "E"}, 96.0},
		{[]string{"E", "F"}, 96.0},
		{[]string{"E", "T"}, 919.0},
		{[]string{"S", "C"}, 9200.0},
		{[]string{"S", "A"}, 9100.0},
	}
	for _, testCase := range testCases {
		g.UpdateWeight(g.Vertices[testCase.vertices[0]], g.Vertices[testCase.vertices[1]], testCase.weight)
		wgt := g.GetEdgeWeight(g.Vertices[testCase.vertices[0]], g.Vertices[testCase.vertices[1]])
		if wgt != testCase.weight {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, wgt)
		}
	}
}

func Test_JSON_Connect(test *testing.T) {
	g := NewGraph()
	a := NewVertex("Google")
	b := NewVertex("Apple")
	c := NewVertex("Samsung")
	g.Connect(a, b, 0.0)
	g.Connect(a, c, 0.0)
	l := a.GetOutVerticesSize()
	if l != 2 {
		test.Error("Google should have 2 outgoing edges but", l)
	}
	le := g.GetEdgesSize()
	if le != 2 {
		test.Error("This graph should have 2 edges but", le)
	}
}

func Test_JSON_DeleteEdge(test *testing.T) {
	g := NewGraph()
	a := g.CreateAndAddToGraph("Google")
	b := g.CreateAndAddToGraph("Apple")
	c := g.CreateAndAddToGraph("Samsung")
	g.Connect(a, b, 0.0)
	g.Connect(a, b, 1.0)
	g.Connect(a, b, 2.0)
	g.Connect(a, b, 3.0)
	g.Connect(a, c, 4.0)
	g.Connect(a, c, 5.0)
	g.Connect(a, c, 6.0)
	g.Connect(b, c, 7.0)

	g.DeleteEdge(a, c)
	if g.GetEdgesSize() != 2 {
		test.Error("Should only have 2 edges:", g.GetEdgesSize())
	}

	g.DeleteEdge(b, c)
	if g.GetEdgesSize() != 1 {
		test.Error("Should only have 1 edges:", g.GetEdgesSize())
	}

	g.DeleteEdge(a, b)
	if g.GetEdgesSize() != 0 {
		test.Error("Should only have 0 edge:", g.GetEdgesSize())
	}

	testCases1 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 29},
		{[]string{"S", "A"}, 29},
		{[]string{"S", "C"}, 29},
		{[]string{"A", "B"}, 29},
		{[]string{"A", "D"}, 29},
		{[]string{"A", "S"}, 29},
		{[]string{"A", "T"}, 29},
		{[]string{"T", "A"}, 29},
		{[]string{"D", "E"}, 29},
		{[]string{"E", "D"}, 29},
		{[]string{"C", "E"}, 29},
		{[]string{"B", "E"}, 29},
		{[]string{"D", "T"}, 29},
		{[]string{"T", "D"}, 29},
		{[]string{"F", "E"}, 29},
		{[]string{"E", "F"}, 29},
		{[]string{"E", "T"}, 29},
	}
	for _, testCase := range testCases1 {
		g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
		o := g1.GetEdgesSize()
		g1.DeleteEdge(g1.Vertices[testCase.vts[0]], g1.Vertices[testCase.vts[1]])
		// g1.DeleteEdge(g1.Vertices[testCase.vts[1]), g1.Vertices[testCase.vts[0]))
		n := g1.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("In testgraph1, %+v is deleted. Expected '%#v' edges left. But %#v, originally %#v", testCase.vts, testCase.edgnum, n, o)
		}
	}

	testCases2 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 23},
		{[]string{"S", "A"}, 23},
		{[]string{"S", "C"}, 23},
		{[]string{"A", "B"}, 23},
		{[]string{"A", "D"}, 23},
		{[]string{"A", "S"}, 23},
		{[]string{"A", "T"}, 23},
		{[]string{"T", "A"}, 23},
		{[]string{"D", "E"}, 24},
		{[]string{"E", "D"}, 23},
		{[]string{"C", "E"}, 23},
		{[]string{"B", "E"}, 23},
		{[]string{"D", "T"}, 24},
		{[]string{"T", "D"}, 23},
		{[]string{"F", "E"}, 23},
		{[]string{"E", "F"}, 23},
		{[]string{"E", "T"}, 23},
	}
	for _, testCase := range testCases2 {
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		g2.DeleteEdge(g2.Vertices[testCase.vts[0]], g2.Vertices[testCase.vts[1]])
		// g2.DeleteEdge(g2.Vertices[testCase.vts[1]), g2.Vertices[testCase.vts[0]))
		n := g2.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("In testgraph2, %+v is deleted. Expected '%#v' edges left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}

	testCases1b := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 28},
		{[]string{"S", "A"}, 28},
		{[]string{"S", "C"}, 28},
		{[]string{"A", "B"}, 28},
		{[]string{"A", "D"}, 28},
		{[]string{"A", "S"}, 28},
		{[]string{"A", "T"}, 28},
		{[]string{"T", "A"}, 28},
		{[]string{"D", "E"}, 28},
		{[]string{"E", "D"}, 28},
		{[]string{"C", "E"}, 28},
		{[]string{"B", "E"}, 28},
		{[]string{"D", "T"}, 28},
		{[]string{"T", "D"}, 28},
		{[]string{"F", "E"}, 28},
		{[]string{"E", "F"}, 28},
		{[]string{"E", "T"}, 28},
	}
	for _, testCase := range testCases1b {
		g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
		g1.DeleteEdge(g1.Vertices[testCase.vts[0]], g1.Vertices[testCase.vts[1]])
		g1.DeleteEdge(g1.Vertices[testCase.vts[1]], g1.Vertices[testCase.vts[0]])
		n := g1.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("(Bi-direction) In testgraph1, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}

	testCases2b := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "B"}, 22},
		{[]string{"S", "A"}, 22},
		{[]string{"S", "C"}, 22},
		{[]string{"A", "B"}, 22},
		{[]string{"A", "D"}, 23},
		{[]string{"A", "S"}, 22},
		{[]string{"F", "D"}, 23},
		{[]string{"T", "A"}, 22},
		{[]string{"D", "E"}, 23},
		{[]string{"E", "D"}, 23},
		{[]string{"C", "E"}, 23},
		{[]string{"B", "E"}, 22},
		{[]string{"D", "T"}, 23},
		{[]string{"T", "D"}, 23},
		{[]string{"F", "E"}, 22},
		{[]string{"E", "F"}, 22},
		{[]string{"E", "T"}, 22},
	}
	for _, testCase := range testCases2b {
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		g2.DeleteEdge(g2.Vertices[testCase.vts[0]], g2.Vertices[testCase.vts[1]])
		g2.DeleteEdge(g2.Vertices[testCase.vts[1]], g2.Vertices[testCase.vts[0]])
		n := g2.GetEdgesSize()
		if n != testCase.edgnum {
			test.Errorf("(Bi-direction) In testgraph2, %+v is deleted. Expected '%#v' vertices left. But %#v", testCase.vts, testCase.edgnum, n)
		}
	}
}

func Test_JSON_DeleteVertex(test *testing.T) {
	testCases1 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "A", "D"}, 10},
		{[]string{"A", "T", "C"}, 12},
		{[]string{"D", "E", "S"}, 6},
		{[]string{"F", "S", "T"}, 12},
		{[]string{"B", "C", "D"}, 10},
		{[]string{"D", "A", "E"}, 6},
		{[]string{"E", "F", "A"}, 8},
		{[]string{"D", "A", "T"}, 10},
	}
	for _, testCase := range testCases1 {
		g1 := FromJSON("../../files/testgraph.json", "testgraph.001")
		for _, v := range testCase.vts {
			g1.DeleteVertex(g1.Vertices[v])
		}

		n := g1.GetEdgesSize()
		v := g1.GetVerticesSize()
		if n != testCase.edgnum {
			test.Errorf("testgraph1: deleted %+v, expected '%#v' edges to be left. But %#v edges left and %#v vertices left. %+v", testCase.vts, testCase.edgnum, n, v, g1.Vertices[testCase.vts[0]])
		}
	}

	testCases2 := []struct {
		vts    []string
		edgnum int
	}{
		{[]string{"S", "A", "D"}, 9},
		{[]string{"A", "T", "C"}, 9},
		{[]string{"D", "E", "S"}, 6},
		{[]string{"F", "S", "T"}, 8},
		{[]string{"B", "C", "D"}, 10},
		{[]string{"D", "A", "E"}, 6},
		{[]string{"E", "F", "A"}, 6},
		{[]string{"D", "A", "T"}, 9},
	}
	for _, testCase := range testCases2 {
		g2 := FromJSON("../../files/testgraph.json", "testgraph.002")
		for _, v := range testCase.vts {
			g2.DeleteVertex(g2.Vertices[v])
		}
		n := g2.GetEdgesSize()
		v := g2.GetVerticesSize()
		if n != testCase.edgnum {
			test.Errorf("testgraph2: deleted %+v, expected '%#v' edges to be left. But %#v edges left and %#v vertices left.", testCase.vts, testCase.edgnum, n, v)
		}
	}
}
