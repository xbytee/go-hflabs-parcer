package usecase

import (
	"GoHFLabsParcer/config"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tanaikech/go-gdoctableapp"
)

func CheckTable(conf *config.Config, client *http.Client) error {
	gInit := gdoctableapp.New()
	resp, err := gInit.Docs(conf.DocumentID).GetTables().Do(client)
	if err != nil {
		return err
	}

	status, desc, err := ParceTable(conf.URLFromParce)
	if err != nil {
		return err
	}

	if resp.Tables == nil {
		if err := createAndAppend(status, desc, conf.DocumentID, client); err != nil {
			return err
		}
	}

	if resp.Tables != nil {
		if err := deliteTable(conf.DocumentID, client); err != nil {
			logrus.Errorf("error delite table: %v", err)
			return err
		}

		if err := createAndAppend(status, desc, conf.DocumentID, client); err != nil {
			return err
		}
	}

	return nil
}

func createAndAppend(status []string, desc []string, docID string, client *http.Client) error {
	if err := createTable(docID, client); err != nil {
		logrus.Errorf("error create table: %v", err)
		return err
	}

	for i := 0; i < len(status); i++ {
		if err := appendTable(status[i], desc[i], docID, client); err != nil {
			logrus.Errorf("error add elements to table: %v", err)
		}
	}
	return nil
}

func createTable(docID string, client *http.Client) error {
	gInit := gdoctableapp.New()

	obj := &gdoctableapp.CreateTableRequest{
		Rows:    1,
		Columns: 2,
		Index:   1,
		Values: [][]interface{}{
			[]interface{}{"HTTP-код ответа", "Описание"},
		},
	}

	_, err := gInit.Docs(docID).CreateTable(obj).Do(client)
	if err != nil {
		return err
	}
	return nil
}

func appendTable(status string, desc string, docID string, client *http.Client) error {
	gInit := gdoctableapp.New()

	obj := &gdoctableapp.AppendRowRequest{
		Values: [][]interface{}{[]interface{}{status, desc}},
	}
	_, err := gInit.Docs(docID).TableIndex(0).AppendRow(obj).Do(client)
	if err != nil {
		return err
	}
	return nil
}

func deliteTable(docID string, client *http.Client) error {
	gInit := gdoctableapp.New()
	_, err := gInit.Docs(docID).TableIndex(0).DeleteTable().Do(client)
	if err != nil {
		return err
	}
	return nil
}
