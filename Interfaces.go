package cyberleafSDK

// Define an interface that both
type PluginMetadataInterface interface {
	GetName() string
	GetDescription() string
	GetType() string
	GetVersion() string
}
