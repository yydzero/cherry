package controllers

type StatsController struct {
	CherryController
}

const (
	STATS_QUERY = `	SELECT g.open_id, g.name, a.articles, a.last_modified, a.crawled_at
				   	FROM cherry_gzh g
					LEFT JOIN (
						SELECT open_id, count(open_id) AS articles, max(last_modified) AS last_modified,
						       max(crawled_at) As crawled_at
						FROM cherry_articles GROUP BY open_id
					) a
					ON a.open_id = g.open_id`

	STATS_QUERY_ROW = `	SELECT g.open_id, g.name, a.articles, a.last_modified, a.crawled_at
				   	FROM cherry_gzh g
					LEFT JOIN (
						SELECT open_id, count(open_id) AS articles, max(last_modified) AS last_modified,
						       max(crawled_at) As crawled_at
						FROM cherry_articles GROUP BY open_id
					) a
					ON a.open_id = g.open_id
					WHERE g.open_id = ?`
)


type Stats struct {
	OpenId	string
	Name 	string
	Articles		int
	LastModified 	uint64
	CrawledAt 		uint64
}

// Signup will register new user
func (this *StatsController) Get() {
	var stats []Stats
	_, err := o.Raw(STATS_QUERY).QueryRows(&stats)
	if err != nil {
		this.Fail(err.Error())
	}

	this.Resource(stats)
}
