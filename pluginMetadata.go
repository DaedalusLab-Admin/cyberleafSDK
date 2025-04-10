package cyberleafSDK

// Default plugin metadata type implementing the PluginMetadataInterface
type PluginMetadata struct {
	Name        string
	Description string
	Type        string
	Version     string
}

func (bpm PluginMetadata) GetName() string {
	return bpm.Name
}

func (bpm PluginMetadata) GetDescription() string {
	return bpm.Description
}

func (bpm PluginMetadata) GetType() string {
	return bpm.Type
}

func (bpm PluginMetadata) GetVersion() string {
	return bpm.Version
}
