package cyberleafSDK

import (
	"go/ast"
)

// Represents the optimization of a file by an optimizer
type Optimization struct {
	Optimizer   string
	Description string
}

// Represents a single issue found by an analyzer
type Issue struct {
	ID            string
	Description   string
	Optimizations []Optimization
}

// Represents the analysis of a file by an analyzer
type Analysis struct {
	Analyser string  // Name of the analyser
	Issues   []Issue // List of issues found by the analyser
}

// Represents the item shared and updated by the plugins in the execution chain
type CyberleafChannelItem struct {
	Target   string     // Path to the file being processed
	AST      *ast.File  // Abstract Syntax Tree of the file
	From     string     // Name of the last plugin that processed this item
	Analysis []Analysis // Contains the results of the execution chain
}
