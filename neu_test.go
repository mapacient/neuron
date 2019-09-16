package edge

import (
	"testing"
)

// test open or create
func TestOpen(t *testing.T) {
	const name = "test_neu"
	edge, err := Open(name, 1024, EDGE_RDWR|EDGE_CREAT)
	if err != nil {
		t.Fatal(err)
	}
	addr, err := edge.Attach()
	if err != nil {
		panic(err)
	}
	t.Log("addr:", addr)
	if err = edge.Detach(); err != nil {
		t.Fatal(err)
	}
	if err = edge.Close(); err != nil {
		t.Fatal(err)
	}
}

// test open only
func TestOpen2(t *testing.T) {
	const name = "test_neu"
	edge, err := Open(name, 1024, EDGE_RDWR)
	if err != nil {
		t.Skip("I think edge neu is nil")
		return
	}
	t.Error("failed, should not open successfully.")
	if err = edge.Close(); err != nil {
		t.Fatal(err)
	}
}

// test create and open
func TestOpen3(t *testing.T) {
	const name = "test_neu"
	edge, err := Open(name, 1024, EDGE_RDWR|EDGE_CREAT)
	if err != nil {
		t.Fatal(err)
	}
	defer edge.Close()
	edge2, err := Open(name, 0, EDGE_RDWR)
	if err != nil {
		t.Fatal(err)
	}
	addr, err := edge2.Attach()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("addr2:", addr)
	if err = edge2.Detach(); err != nil {
		t.Fatal(err)
	}
	if err = edge2.Close(); err != nil {
		t.Fatal(err)
	}
}

// test create and create only
func TestOpen4(t *testing.T) {
	const name = "test_neu"
	edge, err := Open(name, 1024, EDGE_RDWR|EDGE_CREAT)
	if err != nil {
		panic(err)
	}
	defer edge.Close()
	edge2, err := Open(name, 0, EDGE_RDWR|EDGE_CREAT|EDGE_EXCL)
	if err != nil {
		t.Skip("I think already have the edge")
	}
	t.Error("failed, should not create successfully.")
	if err = edge2.Close(); err != nil {
		t.Fatal(err)
	}
}
