package config

import (
	"os"

	"github.com/spf13/cast"
)

const (
	// StatusPublished ...
	StatusPublished = "published"

	// EnumPostType ...
	EnumPostType = "post"

	// FollowType enum
	FollowType = "follow"
	// UnfollowType enum
	UnfollowType = "unfollow"
	// LikeType enum
	LikeType = "like"
	// CommentType enum
	CommentType = "comment"
	// DislikeType enum
	DislikeType = "dislike"
	// GiftType enum
	GiftType = "gift"

	// UniqueViolationCode ...
	UniqueViolationCode = "23505"

	// ForeignKeyViolationCode ..
	ForeignKeyViolationCode = "23503"

	// PointsForLike ...
	PointsForLike = 1

	// PointsForDislike ...
	PointsForDislike = 1

	// PointsForComment ...
	PointsForComment = 2

	// PointsForPosting ...
	PointsForPosting = 2

)

// Config ...
type Config struct {
	Environment         string // develop, staging, production
	PostgresHost        string
	PostgresPort        int
	PostgresDatabase    string
	PostgresUser        string
	PostgresPassword    string
	LogLevel            string
	RPCPort             string
	TodoServiceHost 	string
	TodoServicePort		int
	KafkaURL                string

}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DB", "test_userdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "muhammad"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12345"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9001"))
	c.KafkaURL = cast.ToString(getOrReturnDefault("KAFKA_URL", "127.0.0.1:9092"))
	c.TodoServiceHost = cast.ToString(getOrReturnDefault("TODO_SERVICE_HOST", "localhost"))
	c.TodoServicePort = cast.ToInt(getOrReturnDefault("TODO_SERVICE_PORT", 9002))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
