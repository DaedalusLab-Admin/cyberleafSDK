package cyberleafSDK

import "go/ast"

// A map where keys are plugin categories and values are lists of plugin names.
type PluginsMap map[string][]string

// TODO move in sdk and use it a input/output of plugins
type CyberleafChannelItem struct {
	Target       string    // Path to the file
	AST          *ast.File // Abstract Syntax Tree of the file
	From         string    // Name of the last plugin that processed this item
	Analysis     string    // Analysis result
	Optimization string    // Optimization result
}
