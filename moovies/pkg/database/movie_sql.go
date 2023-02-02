package database

const createMovieTableSQL = `
CREATE TABLE IF NOT EXISTS "movies" (
	"id"     							uuid 						PRIMARY KEY,
	"title"   						varchar(64) 		NOT NULL,
	"year"                int 		        NOT NULL,
	"rate"                int             NOT NULL
)`

const insertMovieSQL = `
INSERT 
INTO "movies" ("id", "title", "year", "rate")
VALUES ($1, $2, $3, $4)`

const readMovieByTitleSQL = `
SELECT * 
FROM "movies" 
WHERE "title"= $1`

const deleteMovieByRateSQL = `
DELETE FROM "movies"
WHERE "rate">$1`
