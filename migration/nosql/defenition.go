package nosql

import (
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/util/migration"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	Background = true
	StoreId    = "storeId"
	IsDeleted  = "isDeleted"
	IsActive   = "isActive"
	Ascending  = bsonx.Int32(int32(1))
	Descending = bsonx.Int32(int32(-1))

	Timestamp = "timestamp"

	Date       = "date"
	YearMonth  = "yearMonth"
	LocationId = "locationId"
	SearchType = "searchType"

	UserCollection         = "users"
	StoryCollection        = "stories"
	StoryChapterCollection = "stories_collection"
)

var (
	Script = make(map[int64]*migration.NoSqlScript)
)
