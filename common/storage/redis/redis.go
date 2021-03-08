package redis

import (
	"context"
	"reflect"
	"time"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	ConnectionURL string
}

type Storage struct {
	*redis.Client
}

func New(config *Config) (*Storage, errors.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	options, err := redis.ParseURL(config.ConnectionURL)
	if err != nil {
		return nil, errors.New("error parsing redis url", errors.InternalServerError)
	}

	options.MaxRetries = 3
	options.ReadTimeout = 15 * time.Second
	options.WriteTimeout = 15 * time.Second
	options.DialTimeout = 15 * time.Second

	client := redis.NewClient(options)
	if client.Ping(ctx).Err() != nil {
		return nil, errors.New("error connectiong to redis", errors.InternalServerError)
	}

	return &Storage{ client }, errors.Nil()
}

func Marshal(data interface{}) (map[string]interface{}, errors.Error) {
	dataType := reflect.TypeOf(data)
	value := reflect.ValueOf(data)

	if dataType == nil {
		return nil, errors.New("error marshaling, data supplied is nil", errors.InternalServerError)
	}

	switch dataType.Kind() {
	case reflect.Struct:
		// If the value is a struct we do nothing
	case reflect.Ptr:
		// If the value is a pointer we get the underlying value
		dataType = dataType.Elem()
		value = value.Elem()
	default:
		return nil, errors.New("failed marshalling, data supplied wasn't either a struct or a struct ptr", errors.InternalServerError)
	}

	redisMap := make(map[string]interface{}, dataType.NumField())

	// We traverse through the fields and check for redis tags and it's values
	for i := 0; i < dataType.NumField(); i++ {

		// Checks if the field's package path is "". Exported fields have empty package path's
		// Also check's if we can get the value of the field out with CanInterface()
		if dataType.Field(i).PkgPath != "" || !value.Field(i).CanInterface() {
			continue
		}

		tag := dataType.Field(i).Tag.Get("redis")

		// If the tag is an empty string, implicitly or explicitly we continue
		if tag == "" {
			continue
		}
		redisMap[tag] = value.Field(i).Interface()

	}

	return redisMap, errors.Nil()
}

func Unmarshal(input map[string]string, output interface{}) errors.Error {
	decoderConfig := &mapstructure.DecoderConfig{
		TagName:          "redis",
		Result:           output,
		WeaklyTypedInput: true,
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return errors.New("error unmarshaling redis map", errors.InternalServerError)
	}

	err = decoder.Decode(input)
	if err != nil {
		return errors.New("error unmarshaling redis map", errors.InternalServerError)
	}

	return errors.Nil()
}
