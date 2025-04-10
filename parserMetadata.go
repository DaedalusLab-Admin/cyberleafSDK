package cyberleafSDK

// Parser plugin metadata type implementing the PluginMetadataInterface
type ParserMetadata struct {
	PluginMetadata         // embedding PluginMetadata
	AcceptedFileExtensions []string
}

func (pm ParserMetadata) GetName() string {
	return pm.Name
}

func (pm ParserMetadata) GetDescription() string {
	return pm.Description
}

func (pm ParserMetadata) GetType() string {
	return pm.Type
}

func (pm ParserMetadata) GetVersion() string {
	return pm.Version
}

func (pm ParserMetadata) GetAcceptedFileExtensions() []string {
	return pm.AcceptedFileExtensions
}
