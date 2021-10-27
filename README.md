# dsg-lab-data-processing

## Preparation

Add the entrypoint and rating CSVs to `data/raw/`

## Running the code

```
go run cmd/rawdataextractor/main.go <input-entrypoint-path> <input-rating-path> <output-persons-path> <output-ratings-path>
```

```
go mod download
go run cmd/rawdataextractor/main.go data/raw/entrypoint.csv data/raw/rating.csv data/processed/persons.csv data/processed/ratings.csv
```
