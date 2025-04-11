package cyberleafSDK

// Optimizer plugin metadata type implementing the PluginMetadataInterface
type OptimizerMetadata struct {
	PluginMetadata    // embedding PluginMetadata
	RequiredAnalysers []string
	AcceptedAnalysers []string
}

func (metadata OptimizerMetadata) GetName() string {
	return metadata.Name
}

func (metadata OptimizerMetadata) GetDescription() string {
	return metadata.Description
}

func (metadata OptimizerMetadata) GetType() string {
	return metadata.Type
}

func (metadata OptimizerMetadata) GetVersion() string {
	return metadata.Version
}

// Implement GetRequiredAnalysers for OptimizerMetadata
func (metadata OptimizerMetadata) GetRequiredAnalysers() []string {
	return metadata.RequiredAnalysers
}

// Implement GetAcceptedAnalysers for OptimizerMetadata
func (metadata OptimizerMetadata) GetAcceptedAnalysers() []string {
	return metadata.AcceptedAnalysers
}
