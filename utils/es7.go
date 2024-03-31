package utils

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
)

type Elastic struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	IndexName string `json:"index_name"`
	User      string `json:"user"`
	Password  string `json:"password"`
}

var (
	esClient *elastic.Client
	ctx      = context.Background()
)

// 实例化es客户端
func InitEs(e *Elastic) (*elastic.Client, error) {
	var err error
	esClient, err = elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%d", e.Host, e.Port)), elastic.SetSniff(false), elastic.SetBasicAuth(e.User, e.Password))
	return esClient, err
}

// 判断索引是否存在
func ExistIndex(indexName string) (bool, error) {
	exists, err := esClient.IndexExists(indexName).Do(ctx)
	return exists, err
}

// 创建索引
// indexName 索引名称,mapping 映射的结构体
func CreateIndex(indexName string, mapping string) (*elastic.IndicesCreateResult, error) {
	createIndex, err := esClient.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	return createIndex, err
}

// 向索引写入单条数据
func AddDocToIndex(indexName string, id int, doc interface{}) error {
	_, err := esClient.Index().
		Index(indexName).
		Id(strconv.Itoa(id)).
		BodyJson(doc).
		Do(ctx)
	return err
}

// 根据文档id查询数据
func SearchDocByDocID(indexName string, id int) (*elastic.GetResult, error) {
	result, err := esClient.Get().
		Index(indexName).
		Id(strconv.Itoa(id)).
		Do(ctx)
	return result, err
}

// 精确查询,term是精确查询，字段类型keyword 不能是text
func TermQuery(indexName, field, value string, offset, limit int) (*elastic.SearchResult, error) {
	termQuery := elastic.NewTermQuery(field, value)
	result, err := esClient.Search().
		Index(indexName).
		Query(termQuery).
		From(offset).Size(limit).
		Pretty(true).
		Do(ctx)
	return result, err
}

// 通过文档ID更改信息
func UpdateByDocId(indexName string, id int, doc interface{}) error {
	_, err := esClient.Update().
		Index(indexName).
		Id(strconv.Itoa(id)).
		Doc(doc).
		Do(ctx)

	return err
}

// 词项多条件精确查询
func TermsQuery(indexName, field string, offset, limit int, values ...interface{}) (*elastic.SearchResult, error) {
	termQuery := elastic.NewTermsQuery(field, values...)
	result, err := esClient.Search().
		Index(indexName).
		Query(termQuery).
		From(offset).Size(limit).
		Pretty(true).
		Do(ctx)
	return result, err
}

// 词项的区间查询
func RangeQuery(indexName, field string, offset, limit int, gte, lte interface{}) (*elastic.SearchResult, error) {
	rangeQuery := elastic.NewRangeQuery(field).Gte(gte).Lte(lte)
	result, err := esClient.Search().
		Index(indexName).
		Query(rangeQuery).
		From(offset).Size(limit).
		Pretty(true).
		Do(ctx)
	return result, err
}

// 高亮搜索
func SearchWithHighlight(indexName, field, msg string, offset, limit int) (*elastic.SearchResult, error) {
	query := elastic.NewMatchQuery(field, msg)
	highlight := elastic.NewHighlight().Field(field)
	highlight.PreTags("<span color='red'>")
	highlight.PostTags("</span>")
	result, err := esClient.Search().
		Index(indexName).
		Query(query).
		Highlight(highlight).
		From(offset).Size(limit).
		Pretty(true).
		Do(ctx)
	return result, err
}

// 同时满足两个字段查询
func SearchWithBothFields(indexName string, field1, field2 string, field1Value, field2Value interface{}) (*elastic.SearchResult, error) {
	termQuery1 := elastic.NewTermQuery(field1, field1Value)
	termQuery2 := elastic.NewTermQuery(field2, field2Value)
	boolQuery := elastic.NewBoolQuery().Must(termQuery1, termQuery2)
	result, err := esClient.Search().
		Index(indexName).
		Query(boolQuery).
		Do(ctx)
	return result, err
}

// 两个条件可以只满足一个
func SearchWithMixedFields(indexName string, field1, field2 string, field1Value, field2Value interface{}) (*elastic.SearchResult, error) {
	termQuery1 := elastic.NewTermQuery(field1, field1Value)
	termQuery2 := elastic.NewTermQuery(field2, field2Value)
	//可以不匹配
	notMatchQuery2 := elastic.NewBoolQuery().MustNot(termQuery2)
	boolQuery := elastic.NewBoolQuery().Must(termQuery1).Should(termQuery2).Should(notMatchQuery2)
	result, err := esClient.Search().
		Index(indexName).
		Query(boolQuery).
		Do(ctx)
	return result, err
}

// 根据症状、医院、科室、医生名查询医生列表
func SearchDoctor(indexName, field1, field2, field3, field4, value string, offset, limit int) (*elastic.SearchResult, error) {
	query := elastic.NewBoolQuery().
		Should(
			elastic.NewWildcardQuery(field1, "*"+value+"*"),
			elastic.NewWildcardQuery(field2, "*"+value+"*"),
			elastic.NewWildcardQuery(field3, "*"+value+"*"),
			elastic.NewWildcardQuery(field4, "*"+value+"*"),
		).
		MinimumShouldMatch("1")

	result, err := esClient.Search().
		Index(indexName).
		Query(query).
		From(offset).Size(limit).
		Do(ctx)
	return result, err
}
