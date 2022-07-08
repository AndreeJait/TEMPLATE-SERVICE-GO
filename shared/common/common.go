package common

import (
	"encoding/binary"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	JSONUnixTimestamp time.Time

	SearchType  string
	StatusPrice string
)

const (
	DefaultLang = "id"

	Timestamp = "timestamp"

	IsDeletedLabel = "isDeleted"
	Failed         = "FAILED"

	PUBLISHED = "PUBLISHED"

	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
)

var (
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

type (
	BaseMongo struct {
		Id          *primitive.ObjectID `json:"id" bson:"_id"`
		Version     int64               `json:"version" bson:"version"`
		CreatedDate JSONUnixTimestamp   `json:"createdDate" bson:"createdDate"`
		CreatedBy   string              `json:"createdBy" bson:"createdBy"`
		UpdatedDate JSONUnixTimestamp   `json:"updatedDate" bson:"updatedDate"`
		UpdatedBy   string              `json:"updatedBy" bson:"updatedBy"`
		StoreId     string              `json:"storeId" bson:"storeId"`
		IsDeleted   int                 `json:"isDeleted" bson:"isDeleted"`
	}

	MandatoryRequestDto struct {
	}

	BaseResponseDto struct {
		Code       string      `json:"code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Errors     []string    `json:"errors"`
		ServerTime int64       `json:"serverTime"`
	}
)

func (t JSONUnixTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *JSONUnixTimestamp) UnmarshalJSON(s []byte) error {
	r := strings.ReplaceAll(string(s), `"`, "")

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(0, q*int64(time.Millisecond))
	return nil
}

func (t JSONUnixTimestamp) String() string { return time.Time(t).String() }

func (t JSONUnixTimestamp) Timestamp() int64 {
	return time.Time(t).UnixNano() / int64(time.Millisecond)
}

func (t JSONUnixTimestamp) TimestampForSecond() int64 {
	return time.Time(t).UnixNano() / int64(time.Second)
}

func (t JSONUnixTimestamp) Time() time.Time {
	return time.Time(t)
}

func (t *JSONUnixTimestamp) UnmarshalBSON(s []byte) error {
	i := int64(binary.LittleEndian.Uint64(s))
	*(*time.Time)(t) = time.Unix(0, i*int64(time.Millisecond))

	return nil
}
