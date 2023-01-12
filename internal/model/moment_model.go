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
		FindMany(ctx context.Context, id string, count, skip int64) ([]*Moment, error)
		Search(ctx context.Context, communityId, keyword string, count, skip int64) ([]*Moment, error)
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

func (m *customMomentModel) FindMany(ctx context.Context, communityId string, count, skip int64) ([]*Moment, error) {
	oid, err := primitive.ObjectIDFromHex(communityId)
	if err != nil {
		return nil, ErrInvalidObjectId
	}
	data := make([]*Moment, 0, 20)
	opt := &options.FindOptions{
		Skip:  &skip,
		Limit: &count,
		Sort:  bson.M{"createAt": -1},
	}
	err = m.conn.Find(ctx, &data, bson.M{"communityId": oid}, opt)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (m *customMomentModel) Search(ctx context.Context, communityId, keyword string, count, skip int64) ([]*Moment, error) {
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
			"createAt": map[string]any{
				"order": "desc",
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}
	res, err := search(
		search.WithIndex(MomentIndexName),
		search.WithContext(ctx),
		search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
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
		return nil, err
	}
	hits := r["hits"].(map[string]any)["hits"].([]any)
	moments := make([]*Moment, 0, 10)
	for i := range hits {
		hit := hits[i].(map[string]any)
		moment := &Moment{}
		source := hit["_source"].(map[string]any)
		if source["createAt"], err = time.Parse("2006-01-02T15:04:05Z07:00", source["createAt"].(string)); err != nil {
			return nil, err
		}
		if source["updateAt"], err = time.Parse("2006-01-02T15:04:05Z07:00", source["updateAt"].(string)); err != nil {
			return nil, err
		}
		hit["_source"] = source
		err := mapstructure.Decode(hit["_source"], moment)
		if err != nil {
			return nil, err
		}
		oid := hit["_id"].(string)
		id, err := primitive.ObjectIDFromHex(oid)
		if err != nil {
			return nil, err
		}
		moment.ID = id
		moments = append(moments, moment)
	}
	return moments, nil
}
