// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package shed

var (
	// LevelDB key value for storing the schema.
	keySchema = []byte{0}
	// LevelDB key prefix for all field type.
	// LevelDB keys will be constructed by appending name values to this prefix.
	keyPrefixFields byte = 1
	// LevelDB key prefix from which indexing keys start.
	// Every index has its own key prefix and this value defines the first one.
	keyPrefixIndexStart byte = 2 // Q: or maybe a higher number like 7, to have more space for potential specific perfixes
)

// schema is used to serialize known database structure information.
type schema struct {
	Fields  map[string]fieldSpec `json:"fields"`  // keys are field names
	Indexes map[byte]indexSpec   `json:"indexes"` // keys are index prefix bytes
}

// fieldSpec holds information about a particular field.
// It does not need Name field as it is contained in the
// schema.Field map key.
type fieldSpec struct {
	Type string `json:"type"`
}

// indexSpec holds information about a particular index.
// It does not contain index type, as indexes do not have type.
type indexSpec struct {
	Name string `json:"name"`
}

// schemaFieldKey retrieves the complete LevelDB key for
// a particular field from the schema definition.
func (db *DB) schemaFieldKey(name, fieldType string) (key []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RenameIndex changes the schema so that an existing index name is changed
// while preserving its data by keeping the same internal key prefix.
// Renaming indexes is useful when encoding functions can be backward compatible
// to avoid data migrations.
func (db *DB) RenameIndex(name, newName string) (renamed bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// schemaIndexPrefix retrieves the complete LevelDB prefix for
// a particular index.
func (db *DB) schemaIndexPrefix(name string) (id byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// getSchema retrieves the complete schema from
// the database.
func (db *DB) getSchema() (s schema, err error) {
	_ = "STUB: not implemented"
	return *new(schema), nil
}

// putSchema stores the complete schema to
// the database.
func (db *DB) putSchema(s schema) (err error) { _ = "STUB: not implemented"; return nil }
