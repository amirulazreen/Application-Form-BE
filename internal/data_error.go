package internal

type DBError struct {
	QueryFail    string
	InsertFail   string
	NameConflict string
}

var Errors = DBError{
	QueryFail:    "Database query error",
	InsertFail:   "Failed to insert record",
	NameConflict: "Organization of the same name already exists in the database",
}
