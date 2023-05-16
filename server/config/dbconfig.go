package config

import "os"

func DbUrl() string {
	return os.Getenv("DB_URL")
}

func DbTable() string {
	return os.Getenv("DB_TABLE")
}

func DbGsiName() string {
	return os.Getenv("DB_GSI_NAME")
}

func DbGsiAttr() string {
	return os.Getenv("DB_GSI_ATTR")
}
