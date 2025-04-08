package cyberleafSDK

// A map where keys are plugin categories and values are lists of plugin names.
// TODO must be accessible from Plugins
type PluginsMap map[string][]string

type PluginMetadata struct {
	// The name of the plugin
	Name string
	// The description of the plugin
	Description string
	// The type of the plugin
	Type string
	// The version of the plugin
	Version string
	// Required plugins for this plugin
	RequiredPlugins PluginsMap
	// Accepted analysers for this plugin
	AcceptedAnalysers []string
}
