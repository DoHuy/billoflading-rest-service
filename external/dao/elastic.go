package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v6"
	"log"
	"vtp-apis/external/dao/models"
	"vtp-apis/external/domain"
)

const (
	chitietdon_index      = "chitiet_don"
	vandonhanhtrinh_index = "vandon_hanhtrinh"
)

type ElasticsearchDAO struct {
	client *elasticsearch.Client
}

func (s *ElasticsearchDAO) execQuery(ctx context.Context, query *bytes.Buffer, index string, result interface{}) error {
	// Perform the search request.
	res, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex(index),
		s.client.Search.WithBody(query),
		s.client.Search.WithTrackTotalHits(true),
		s.client.Search.WithPretty(),
	)
	if err != nil {
		return err
	}
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Printf("Error parsing the response body: %s", err)
			return err
		} else {
			// Print the response status and error information.
			fmt.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return errors.New(fmt.Sprintf("type: %v, reason: %v", e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"]))
		}
	}
	var r map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return err
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s", hit.(map[string]interface{})["_index"])
		t, _ := json.Marshal(hit.(map[string]interface{})["_source"])
		json.Unmarshal(t, &result)
	}
	defer res.Body.Close()
	return nil
}
func (s *ElasticsearchDAO) FetchByID(ctx context.Context, id int) (*models.Chitietdon, *models.Vandonhanhtrinh, error) {
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"order_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, nil, err
	}
	var vandonhanhtrinh models.Vandonhanhtrinh
	if err := s.execQuery(ctx, &buf, vandonhanhtrinh_index, &vandonhanhtrinh); err != nil {
		return nil, nil, err
	}
	// lay chi tiet don
	// Build the request body.
	query = map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"order_number": "551539242",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, nil, err
	}
	var order models.Chitietdon
	if err := s.execQuery(ctx, &buf, chitietdon_index, &order); err != nil {
		return nil, nil, err
	}
	return &order, &vandonhanhtrinh, nil
}

func (s *ElasticsearchDAO) UpdateOrderStatus(ctx context.Context, state domain.Trangthaivandon) error {
	panic("implement me")
}

func NewElasticsearchDAO(db *elasticsearch.Client) DAO {
	return &ElasticsearchDAO{client: db}
}
