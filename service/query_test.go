package services

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/couchbase-examples/golang-quickstart/models"
)

// fakeRows replays raw JSON rows the way *gocb.QueryResult would.
type fakeRows struct {
	rows   []string
	pos    int
	err    error
	closed bool
}

func (f *fakeRows) Next() bool {
	if f.pos >= len(f.rows) {
		return false
	}
	f.pos++
	return true
}

func (f *fakeRows) Row(valuePtr interface{}) error {
	return json.Unmarshal([]byte(f.rows[f.pos-1]), valuePtr)
}

func (f *fakeRows) Err() error   { return f.err }
func (f *fakeRows) Close() error { f.closed = true; return nil }

func TestCollectRowsDecodesEachRowIntoFreshValue(t *testing.T) {
	rows := &fakeRows{rows: []string{
		`{"name":"First","iata":"FA","icao":"FST"}`,
		`{"name":"Second","icao":"SND"}`,
	}}

	got, err := collectRows[models.Airline](rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []models.Airline{
		{Name: "First", IATA: "FA", ICAO: "FST"},
		{Name: "Second", ICAO: "SND"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
	if !rows.closed {
		t.Error("expected rows to be closed")
	}
}

func TestCollectRowsReturnsEmptySliceForNoRows(t *testing.T) {
	got, err := collectRows[models.Destination](&fakeRows{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("expected empty non-nil slice, got %#v", got)
	}
}

func TestCollectRowsReturnsStreamError(t *testing.T) {
	streamErr := errors.New("query failed mid-stream")
	rows := &fakeRows{rows: []string{`{"destinationairport":"JFK"}`}, err: streamErr}

	got, err := collectRows[models.Destination](rows)
	if !errors.Is(err, streamErr) {
		t.Fatalf("expected stream error, got %v", err)
	}
	if got != nil {
		t.Fatalf("expected no partial results, got %+v", got)
	}
}

func TestCollectRowsReturnsDecodeError(t *testing.T) {
	rows := &fakeRows{rows: []string{`{"destinationairport":42}`}}

	if _, err := collectRows[models.Destination](rows); err == nil {
		t.Fatal("expected decode error")
	}
}
