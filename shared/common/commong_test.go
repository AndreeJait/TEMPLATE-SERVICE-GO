package common

import (
	"testing"
	"time"
)

func TestBase_MarshalJSON(t *testing.T) {
	tm := JSONUnixTimestamp(time.Date(2019, 2, 2, 0, 0, 0, 0, time.Local))
	t.Run("when MarshalJson", func(t *testing.T) {
		_, _ = tm.MarshalJSON()
	})
}

func TestBase_UnmarshalJSON(t *testing.T) {
	var tm JSONUnixTimestamp
	t.Run("when UnmarshalJSON error", func(t *testing.T) {
		str := "{{"
		err := tm.UnmarshalJSON([]byte(str))
		if err == nil {
			t.Error("it should be error")
		}
	})

	t.Run("when UnmarshalJSON", func(t *testing.T) {
		str := "1570640400000"
		_ = tm.UnmarshalJSON([]byte(str))
	})
}

func TestBase_String(t *testing.T) {
	tm := JSONUnixTimestamp(time.Date(2019, 2, 2, 0, 0, 0, 0, time.Local))
	t.Run("when String", func(t *testing.T) {
		_ = tm.String()
	})
}

func TestBase_Timestamp(t *testing.T) {
	tm := JSONUnixTimestamp(time.Date(2019, 2, 2, 0, 0, 0, 0, time.Local))
	t.Run("when Timestamp", func(t *testing.T) {
		_ = tm.Timestamp()
	})
}

func TestBase_TimestampForSecond(t *testing.T) {
	tm := JSONUnixTimestamp(time.Date(2019, 2, 2, 0, 0, 0, 0, time.Local))
	t.Run("when TimestampForSecond", func(t *testing.T) {
		_ = tm.TimestampForSecond()
	})
}

func TestBase_Time(t *testing.T) {
	tm := JSONUnixTimestamp(time.Date(2019, 2, 2, 0, 0, 0, 0, time.Local))
	t.Run("when String", func(t *testing.T) {
		_ = tm.Time()
	})
}

func TestBase_UnmarshallBSON(t *testing.T) {
	tm := JSONUnixTimestamp(time.Date(2019, 2, 2, 0, 0, 0, 0, time.Local))
	t.Run("when String", func(t *testing.T) {
		str := "1570640400000"
		_ = tm.UnmarshalBSON([]byte(str))
	})
}

// func TestIndex(t *testing.T) {

// 	stringVarArray := []string{"string1", "string2"}
// 	t.Run("when found", func(t *testing.T) {
// 		res := Index(stringVarArray, "string2")
// 		assert.EqualValues(t, res, 1)
// 	})

// 	t.Run("when not found", func(t *testing.T) {
// 		res := Index(stringVarArray, "xxx")
// 		assert.EqualValues(t, res, -1)
// 	})

// }
