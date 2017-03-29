package archive

import (
	"database/sql"
)

func ListSubprimers(db sqlQueryable, limit, offset int) ([]*Subprimer, error) {
	rows, err := db.Query(qSubprimersList, limit, offset)
	if err != nil {
		return nil, err
	}
	return UnmarshalBoundedSubprimers(rows, limit)
}

func UnmarshalBoundedSubprimers(rows *sql.Rows, limit int) ([]*Subprimer, error) {
	defer rows.Close()
	subprimers := make([]*Subprimer, limit)
	i := 0
	for rows.Next() {
		u := &Subprimer{}
		if err := u.UnmarshalSQL(rows); err != nil {
			return nil, err
		}
		subprimers[i] = u
		i++
	}

	return subprimers[:i], nil
}
