package model

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MomentCollectionName = "moment"
const MomentIndexName = "meowchat_moment.moment-alias"

var _ MomentModel = (*customMomentModel)(nil)

type (
	// MomentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMomentModel.
	MomentModel interface {
		momentModel
		FindManyByCommunityId(ctx context.Context, id string, count, skip int64) ([]*Moment, int64, error)
		FindManyByUserId(ctx context.Context, id string, count, skip int64) ([]*Moment, int64, error)
		Search(ctx context.Context, communityId, keyword string, count, skip int64) ([]*Moment, int64, error)
	}

	customMomentModel struct {
		*defaultMomentModel
		es *elasticsearch.Client
	}
)

// NewMomentModel returns a model for the mongo.
func NewMomentModel(url, db string, c cache.CacheConf, es config.ElasticsearchConf) MomentModel {
	conn := monc.MustNewModel(url, db, MomentCollectionName, c)
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: es.Addresses,
		Username:  es.Username,
		Password:  es.Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return &customMomentModel{
		defaultMomentModel: newDefaultMomentModel(conn),
		es:                 esClient,
	}
}

func (m *customMomentModel) FindManyByCommunityId(ctx context.Context, communityId string, count, skip int64) ([]*Moment, int64, error) {
	data := make([]*Moment, 0)
	opt := &options.FindOptions{
		Skip:  &skip,
		Limit: &count,
		Sort:  bson.M{"createAt": -1},
	}
	err := m.conn.Find(ctx, &data, bson.M{"communityId": communityId}, opt)
	if err != nil {
		return nil, 0, err
	}
	total, err := m.conn.CountDocuments(ctx, bson.M{"communityId": communityId})
	if err != nil {
		return nil, 0, err
	}
	return data, total, err
}

func (m *customMomentModel) FindManyByUserId(ctx context.Context, userId string, count, skip int64) ([]*Moment, int64, error) {
	data := make([]*Moment, 0)
	opt := &options.FindOptions{
		Skip:  &skip,
		Limit: &count,
		Sort:  bson.M{"createAt": -1},
	}
	err := m.conn.Find(ctx, &data, bson.M{"userId": userId}, opt)
	if err != nil {
		return nil, 0, err
	}
	total, err := m.conn.CountDocuments(ctx, bson.M{"userId": userId})
	if err != nil {
		return nil, 0, err
	}
	return data, total, err
}

func (m *customMomentModel) Search(ctx context.Context, communityId, keyword string, count, skip int64) ([]*Moment, int64, error) {
	search := m.es.Search
	query := map[string]any{
		"from": skip,
		"size": count,
		"query": map[string]any{
			"bool": map[string]any{
				"must": []any{
					map[string]any{
						"term": map[string]any{
							"communityId": communityId,
						},
					},
					map[string]any{
						"multi_match": map[string]any{
							"query":  keyword,
							"fields": []string{"title", "text"},
						},
					},
				},
			},
		},
		"sort": map[string]any{
			"_score": map[string]any{
				"order": "desc",
			},
			"createAt": map[string]any{
				"order": "desc",
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, 0, err
	}
	res, err := search(
		search.WithIndex(MomentIndexName),
		search.WithContext(ctx),
		search.WithBody(&buf),
	)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, 0, err
		} else {
			logx.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r map[string]any
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}
	hits := r["hits"].(map[string]any)["hits"].([]any)
	total := int64(r["hits"].(map[string]any)["total"].(map[string]any)["value"].(float64))
	moments := make([]*Moment, 0, 10)
	for i := range hits {
		hit := hits[i].(map[string]any)
		moment := &Moment{}
		source := hit["_source"].(map[string]any)
		if source["createAt"], err = time.Parse("2006-01-02T15:04:05Z07:00", source["createAt"].(string)); err != nil {
			return nil, 0, err
		}
		if source["updateAt"], err = time.Parse("2006-01-02T15:04:05Z07:00", source["updateAt"].(string)); err != nil {
			return nil, 0, err
		}
		hit["_source"] = source
		err := mapstructure.Decode(hit["_source"], moment)
		if err != nil {
			return nil, 0, err
		}
		oid := hit["_id"].(string)
		id, err := primitive.ObjectIDFromHex(oid)
		if err != nil {
			return nil, 0, err
		}
		moment.ID = id
		moments = append(moments, moment)
	}
	return moments, total, nil
}
