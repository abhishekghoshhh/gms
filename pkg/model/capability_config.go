package model

type CapabilitiesConfig struct {
	config map[string]string
}

func (capabilitiesConfig *CapabilitiesConfig) Get(key string) (string, bool) {
	val, ok := capabilitiesConfig.config[key]
	return val, ok
}

func NewCapabitiesConfig(entries ...entry) *CapabilitiesConfig {
	config := make(map[string]string)
	for _, pair := range entries {
		config[pair.key] = pair.value
	}
	return &CapabilitiesConfig{config}
}
