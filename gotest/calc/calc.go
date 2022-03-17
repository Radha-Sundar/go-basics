package calc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var letters = []rune("1234567890ABCDEFGHJKLMNPQRSTUVWXYZ")

var numbers = []rune("1234567890")

func init() {
	rand.Seed(Now().UnixNano())
}

func RandStrN(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func RandIntN(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(numbers))]
	}

	return string(b)
}

func StringPt(v string) *string {
	return &v
}

func GetBasicAuthHeader(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func JsonForPrint(data interface{}) string {
	b, _ := json.Marshal(data) //nolint:errcheck //simplify
	return strings.ReplaceAll(string(b), "\"", "'")
}

func StructToMap(i interface{}) url.Values {
	values := url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()

	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		tag := typ.Field(i).Tag.Get("json")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(tag, v)
	}

	return values
}

func Now() time.Time {
	return time.Now().UTC()
}

func Trunc(s string, length int) string {
	if len(s) > length {
		s = s[:length]
	}

	return s
}

func Sanitize(s string, allowedCharsPattern string, length int) string {
	pattern := fmt.Sprintf("[^%v]+", allowedCharsPattern)
	s = regexp.MustCompile(pattern).ReplaceAllString(s, ``)

	return Trunc(s, length)
}
