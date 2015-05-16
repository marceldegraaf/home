package storage

// TODO: Implement a Redis meter data storage backend.
//
// This will use a hash for data points (uuid => binary data)
// and will use a sorted set to create an index of uuids with
// Unix epoch timestamps (timestamp is the score, uuid the member)

type Point struct {
	UUID      string
	Timestamp int
	Data      string
}

func Initialize() error                 { return nil }
func Save(p *Point) error               { return nil }
func Range(t1, t2 int) ([]Point, error) { return nil, nil }
