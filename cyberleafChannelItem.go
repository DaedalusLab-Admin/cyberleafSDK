package cyberleafSDK

import (
	"go/ast"
	"sync"
)

// Represents ta single optimization for an issue by an optimizer
type Optimization struct {
	Optimizer   string
	Description string
	AST         *ast.File // Abstract Syntax Tree of the file
	// TODO check if other fields are needed
}

// Represents a single issue found by an analyser
type Issue struct {
	ID                string
	Description       string
	Optimizations     *[]Optimization // List of possible optimizations for this issue
	optimizationMutex sync.Mutex      // Mutex to ensure thread safety when modifying Optimizations
	// TODO check if other fields are needed
}

type Analysis struct {
	Analyser string   // Name of the analyser
	Issues   []*Issue // List of issues found by the analyser
}

// Represents the item shared and updated by the plugins in the execution chain
type CyberleafChannelItem struct {
	Target        string      // Path to the file being processed
	AST           *ast.File   // Abstract Syntax Tree of the file
	Analysis      *[]Analysis // List of analyses performed on the item
	analysisMutex sync.Mutex  // Mutex to ensure thread safety when modifying Analysis
}

func (item *CyberleafChannelItem) AddAnalysis(newAnalysis Analysis) {
	item.analysisMutex.Lock()         // Lock the mutex to ensure mutual exclusion
	defer item.analysisMutex.Unlock() // Unlock the mutex when the function exits

	// Append analysi
	if item.Analysis == nil {
		item.Analysis = &[]Analysis{}
	}
	*item.Analysis = append(*item.Analysis, newAnalysis)
}

func (issue *Issue) AddOptimization(newOptimization Optimization) {
	issue.optimizationMutex.Lock()         // Lock the mutex to ensure mutual exclusion
	defer issue.optimizationMutex.Unlock() // Unlock the mutex when the function exits

	// Append optimization
	if issue.Optimizations == nil {
		issue.Optimizations = &[]Optimization{}
	}
	*issue.Optimizations = append(*issue.Optimizations, newOptimization)
}
