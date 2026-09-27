module github.com/SarthakStha/blog-aggregator

go 1.27.0

replace github.com/SarthakStha/blog-aggregator/internal/config => ./internal/config

replace github.com/SarthakStha/blog-aggregator/internal/database => ./internal/database

replace github.com/SarthakStha/blog-aggregator/internal/feed => ./internal/feed

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.12.3
)
