package config

import "os"

func GoogleAccessToken() string {
	return os.Getenv("GOOGLE_ACCESS_TOKEN")
}

func GoogleRefreshToken() string {
	return os.Getenv("GOOGLE_REFRESH_TOKEN")
}

func GoogleClientID() string {
	return os.Getenv("GOOGLE_CLIENT_ID")
}

func GoogleClientSecret() string {
	return os.Getenv("GOOGLE_CLIENT_SECRET")
}

func AmazonSmtpUsername() string {
	return os.Getenv("AMZN_SMTP_USERNAME")
}

func AmazonSmtpPassword() string {
	return os.Getenv("AMZN_SMTP_PASSWORD")
}
