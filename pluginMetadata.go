package cyberleafSDK

// Default plugin metadata type implementing the PluginMetadataInterface
type PluginMetadata struct {
	Name        string
	Description string
	Type        string
	Version     string
}

func (metadata PluginMetadata) GetName() string {
	return metadata.Name
}

func (metadata PluginMetadata) GetDescription() string {
	return metadata.Description
}

func (metadata PluginMetadata) GetType() string {
	return metadata.Type
}

func (metadata PluginMetadata) GetVersion() string {
	return metadata.Version
}
