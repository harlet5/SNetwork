package db

import (
	"database/sql"
	"errors"
)

type Cat struct {
	CId   int
	CName string
}

func CreateCat(name string) {
	stmt, err := DB.Prepare("INSERT INTO Cats(CName) VALUES (?)")
	DbErrHandler(true, "Cat creation | prepare", err)
	_, err = stmt.Exec(name)
	DbErrHandler(true, "Cat creation | execute", err)
}

func GetCats() []Cat {
	var cats []Cat
	rows, err := DB.Query("SELECT * FROM Cats;")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Cats get | query", err)
		} else {
			DbErrHandler(true, "Cats get | query", err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		var cat Cat
		if err := rows.Scan(&cat.CId, &cat.CName); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				DbErrHandler(false, "Cats get | scan", err)
			} else {
				DbErrHandler(true, "Cats get | scan", err)
			}
		}
		cats = append(cats, cat)
	}
	return cats
}

func GetCatId(name string) (int, error) {
	var id int
	err := DB.QueryRow(`SELECT CId FROM Cats WHERE CName = ? `, name).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Cat get | query", err)
		} else {
			DbErrHandler(true, "Cat get | query", err)
		}
		return 0, err
	}
	return id, nil
}

func GetCatByThread(id string) ([]Cat, error) {
	var cats []Cat
	rows, err := DB.Query("SELECT CName FROM Cats INNER JOIN ThreadsInCats ON Cats.CId = TCcID WHERE TCtId = " + id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Cat thread | query", err)
		} else {
			DbErrHandler(true, "Cat thread | query", err)
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cat Cat
		if err := rows.Scan(&cat.CName); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				DbErrHandler(false, "Cat thread | scan", err)
			} else {
				DbErrHandler(true, "Cat thread | scan", err)
			}
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

func LinkThreadToCats(topId, catId int) error {
	stm, err := DB.Prepare("INSERT INTO ThreadsInCats (TCtId, TCcId) VALUES (?,?)")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Cat link | prepare", err)
		} else {
			DbErrHandler(true, "Cat link | prepare", err)
		}
		return err
	}
	_, err = stm.Exec(topId, catId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Cat link | execute", err)
		} else {
			DbErrHandler(true, "Cat link | execute", err)
		}
	}
	return nil
}
