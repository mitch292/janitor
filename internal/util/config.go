package util

// GetNestedConfigValueAccessor will generate the dot notation string used
// to access a nested config value
func GetNestedConfigValueAccessor(topLevelKey, section, keyName string) string {
	return topLevelKey + "." + section + "." + keyName
}
