package domain

import "time"

type CommonConfig struct {
	SysConfig SysConfig
	Repo      RepoDB
	Cache     Cache
	StorageS3 StorageS3
}

type SysConfig struct {
	DBConfig string `mapstructure:"pg_config"`

	// ClientOrigin   string `mapstructure:"client_origin"`
	ServiceSchema  string `mapstructure:"service_schema"`
	ServiceIP      string `mapstructure:"service_ip"`
	ServicePort    string `mapstructure:"service_port"`
	ServicePathAPI string `mapstructure:"service_path_api"`
	ServicePathUI  string `mapstructure:"service_path_ui"`

	TokenSecret    string        `mapstructure:"token_secret"`
	TokenExpiresIn time.Duration `mapstructure:"token_expired_in"`
	TokenMaxAge    int           `mapstructure:"token_maxage"`
	FileMaxSize    int           `mapstructure:"file_max_size"`
	FileTmpPath    string        `mapstructure:"file_tmp_path"`

	EmailFrom string `mapstructure:"email_from"`
	SMTPHost  string `mapstructure:"smtp_host"`
	SMTPPass  string `mapstructure:"smtp_pass"`
	SMTPPort  int    `mapstructure:"smtp_port"`
	SMTPUser  string `mapstructure:"smtp_user"`

	MetricsScrapeSec time.Duration `mapstructure:"metrics_scrape_sec"`

	S3Endpoint        string `mapstructure:"s3_endpoint"`
	S3AccessKeyID     string `mapstructure:"s3_access_key_id"`
	S3SecretAccessKey string `mapstructure:"s3_secret_access_key"`
	S3EncryptPasswd   string `mapstructure:"s3_encrypt_passwd"`
	S3EnableSSL       bool   `mapstructure:"s3_enable_ssl"`
	S3Bucket          string `mapstructure:"s3_bucket"`
}

func NewCommonConfig(config SysConfig, repo RepoDB, cache Cache, s3 StorageS3) *CommonConfig {
	return &CommonConfig{
		SysConfig: config,
		Repo:      repo,
		StorageS3: s3,
		Cache:     cache,
	}
}
