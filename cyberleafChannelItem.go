package cyberleafSDK

import (
	"go/ast"
	"sync"
)

// Represents a solution for an issue by an optimizer
type Solution struct {
	Optimizer   string `json:"optimizer"`   // Name of the optimizer
	Description string `json:"description"` // Description of the suggested solution
}

// Represents a single issue found by an analyser
type Issue struct {
	ID                string      `json:"id"`                  // Identifier for the issue
	Row               int         `json:"row"`                 // Line number where the issue was found
	Description       string      `json:"description"`         // Description of the issue
	Solutions         *[]Solution `json:"solutions,omitempty"` // List of solutions for this issue by different optimizers (omitted if empty)
	optimizationMutex sync.Mutex  `json:"-"`                   // Mutex to ensure thread safety when modifying Optimizations (excluded from JSON)
}

// Represents a full analysis result by one analyser
type Analysis struct {
	Analyser string   `json:"analyser"` // Name of the analyser
	Issues   []*Issue `json:"issues"`   // List of issues found by the analyser
}

// Represents the item shared and updated by the plugins in the execution chain
type CyberleafChannelItem struct {
	ID            int         `json:"-"`        // ID to sort the items (excluded from JSON)
	Target        string      `json:"target"`   // Path to the file being processed
	AST           *ast.File   `json:"-"`        // Abstract Syntax Tree of the file (excluded from JSON)
	Analysis      *[]Analysis `json:"analysis"` // List of analyses performed on the item
	analysisMutex sync.Mutex  `json:"-"`        // Mutex to ensure thread safety when modifying Analysis (excluded from JSON)
}

// Adds a new analysis in a thread-safe way
func (item *CyberleafChannelItem) AddAnalysis(newAnalysis Analysis) {
	item.analysisMutex.Lock()         // Lock the mutex to ensure mutual exclusion
	defer item.analysisMutex.Unlock() // Unlock the mutex when the function exits

	if item.Analysis == nil {
		item.Analysis = &[]Analysis{}
	}
	*item.Analysis = append(*item.Analysis, newAnalysis)
}

// Adds a new solution to an issue in a thread-safe way
func (issue *Issue) AddOptimization(newOptimization Solution) {
	issue.optimizationMutex.Lock()         // Lock the mutex to ensure mutual exclusion
	defer issue.optimizationMutex.Unlock() // Unlock the mutex when the function exits

	if issue.Solutions == nil {
		issue.Solutions = &[]Solution{}
	}
	*issue.Solutions = append(*issue.Solutions, newOptimization)
}
