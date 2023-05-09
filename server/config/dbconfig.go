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

func UsernamesEnabled() bool {
	if os.Getenv("CFG_USERNAMES_ENABLED") == "true" {
		return true
	} else {
		return false
	}
}
