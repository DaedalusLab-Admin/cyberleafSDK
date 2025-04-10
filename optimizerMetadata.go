package cyberleafSDK

// Optimizer plugin metadata type implementing the PluginMetadataInterface
type OptimizerMetadata struct {
	PluginMetadata    // embedding PluginMetadata
	RequiredAnalysers []string
	AcceptedAnalysers []string
}

func (pm OptimizerMetadata) GetName() string {
	return pm.Name
}

func (pm OptimizerMetadata) GetDescription() string {
	return pm.Description
}

func (pm OptimizerMetadata) GetType() string {
	return pm.Type
}

func (pm OptimizerMetadata) GetVersion() string {
	return pm.Version
}

// Implement GetRequiredAnalysers for OptimizerMetadata
func (pm OptimizerMetadata) GetRequiredAnalysers() []string {
	return pm.RequiredAnalysers
}

// Implement GetAcceptedAnalysers for OptimizerMetadata
func (pm OptimizerMetadata) GetAcceptedAnalysers() []string {
	return pm.AcceptedAnalysers
}
