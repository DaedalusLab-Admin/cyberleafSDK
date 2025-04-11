package cyberleafSDK

// Parser plugin metadata type implementing the PluginMetadataInterface
type ParserMetadata struct {
	PluginMetadata         // embedding PluginMetadata
	AcceptedFileExtensions []string
}

func (metadata ParserMetadata) GetName() string {
	return metadata.Name
}

func (metadata ParserMetadata) GetDescription() string {
	return metadata.Description
}

func (metadata ParserMetadata) GetType() string {
	return metadata.Type
}

func (metadata ParserMetadata) GetVersion() string {
	return metadata.Version
}

func (metadata ParserMetadata) GetAcceptedFileExtensions() []string {
	return metadata.AcceptedFileExtensions
}
