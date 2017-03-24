package archive

import (
	"database/sql"
	"github.com/pborman/uuid"
	"time"
)

// Primer is tracking information about a base URL
type Primer struct {
	Id          string       `json:"id"`
	Created     time.Time    `json:"created"`
	Updated     time.Time    `json:"updated"`
	ShortTitle  string       `json:"shortTitle"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Subprimers  []*Subprimer `json:"subprimers"`
}

// Subprimers returns the list of listed urls for crawling associated with this primer
func (p *Primer) ReadSubprimers(db sqlQueryable) error {
	rows, err := db.Query(qPrimerSubprimers, p.Id)
	if err != nil {
		return err
	}

	defer rows.Close()
	urls := make([]*Subprimer, 0)
	for rows.Next() {
		c := &Subprimer{}
		if err := c.UnmarshalSQL(rows); err != nil {
			return err
		}
		urls = append(urls, c)
	}

	p.Subprimers = urls
	return nil
}

func (p *Primer) Read(db sqlQueryable) error {
	if p.Id != "" {
		row := db.QueryRow(qPrimerById, p.Id)
		return p.UnmarshalSQL(row)
	}
	return ErrNotFound
}

func (p *Primer) Save(db sqlQueryExecable) error {
	prev := &Primer{Id: p.Id}
	if err := prev.Read(db); err != nil {
		if err == ErrNotFound {
			p.Id = uuid.New()
			p.Created = time.Now().Round(time.Second)
			p.Updated = p.Created
			_, err := db.Exec(qPrimerInsert, p.SQLArgs()...)
			return err
		} else {
			return err
		}
	} else {
		p.Updated = time.Now().Round(time.Second)
		_, err := db.Exec(qPrimerUpdate, p.SQLArgs()...)
		return err
	}
	return nil
}

func (p *Primer) Delete(db sqlQueryExecable) error {
	_, err := db.Exec(qPrimerDelete, p.Id)
	return err
}

func (p *Primer) UnmarshalSQL(row sqlScannable) error {
	var (
		id, title, description, short string
		created, updated              time.Time
	)

	if err := row.Scan(&id, &created, &updated, &short, &title, &description); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	*p = Primer{
		Id:          id,
		Created:     created.In(time.UTC),
		Updated:     updated.In(time.UTC),
		ShortTitle:  short,
		Title:       title,
		Description: description,
	}

	return nil
}

func (p *Primer) SQLArgs() []interface{} {
	return []interface{}{
		p.Id,
		p.Created.In(time.UTC),
		p.Updated.In(time.UTC),
		p.ShortTitle,
		p.Title,
		p.Description,
	}
}
