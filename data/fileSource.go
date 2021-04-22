package data

import (
	"encoding/csv"
	"go-api/models"
	"io"
	"os"
	"strconv"
)

type FileSource struct {
	sourceFile string
	Data       *[]models.CardBack
}

func NewFileSource(sourceFile string) (*FileSource, error) {
	data, err := initializeData(sourceFile)
	if err != nil {
		return nil, err
	}
	return &FileSource{sourceFile, data}, nil
}

func initializeData(filePath string) (*[]models.CardBack, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	return readCSVData(csvFile)
}

func readCSVData(r io.Reader) (*[]models.CardBack, error) {
	data := csv.NewReader(r)
	cardBacks := []models.CardBack{}

	isHeader := true
	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if isHeader {
			isHeader = !isHeader
			continue
		}

		cardBack := models.CardBack{}
		cardBack.ID, err = strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		cardBack.SortCategory, err = strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		cardBack.Text = record[2]
		cardBack.Name = record[3]
		cardBack.Image = record[4]
		cardBack.Slug = record[5]
		cardBacks = append(cardBacks, cardBack)
	}

	return &cardBacks, nil
}

func (fs *FileSource) Save() error {
	file, err := os.Create(fs.sourceFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return writeCSVData(fs.Data, file)
}

func writeCSVData(data *[]models.CardBack, w io.Writer) error {
	records := [][]string{{"id", "sortCategory", "text", "name", "image", "slug"}}
	for _, record := range *data {
		r := []string{
			strconv.Itoa(record.ID),
			strconv.Itoa(record.SortCategory),
			record.Text, record.Name, record.Image,
			record.Slug,
		}
		records = append(records, r)
	}

	csvW := csv.NewWriter(w)
	return csvW.WriteAll(records)
}
