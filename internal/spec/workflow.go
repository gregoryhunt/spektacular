package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/cbroglie/mustache"
)

// SpecStep defines a step in the spec creation workflow
type SpecStep struct {
	ID      string // a-h
	Title   string
	Section string
}

// GetSpecStepByID returns the workflow definition for a given step ID (a-h)
func GetSpecStepByID(stepID string) (SpecStep, error) {
	steps := getSpecSteps()
	for _, step := range steps {
		if step.ID == stepID {
			return step, nil
		}
	}
	return SpecStep{}, fmt.Errorf("invalid step ID %s (must be a-h)", stepID)
}

// RenderStepTemplate renders the mustache template for a step with context variables
func RenderStepTemplate(stepID string, specPath string) (string, error) {
	step, err := GetSpecStepByID(stepID)
	if err != nil {
		return "", err
	}

	// Calculate next step
	var nextStep string
	if stepID != "h" {
		nextStep = string(rune(stepID[0] + 1))
	}

	// Find template file
	templatePath := getTemplateFilePath(stepID)
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("loading template for step %s: %w", stepID, err)
	}

	// Render template with mustache
	data := map[string]interface{}{
		"step":      stepID,
		"title":     step.Title,
		"section":   step.Section,
		"spec_path": specPath,
		"next_step": nextStep,
	}

	output, err := mustache.Render(string(templateBytes), data)
	if err != nil {
		return "", fmt.Errorf("rendering template for step %s: %w", stepID, err)
	}

	return output, nil
}

func getTemplateFilePath(stepID string) string {
	// Try to find templates/ directory relative to cwd first
	cwd, err := os.Getwd()
	if err == nil {
		path := filepath.Join(cwd, "templates", "spec-steps", fmt.Sprintf("%s-*.md", stepID))
		matches, err := filepath.Glob(path)
		if err == nil && len(matches) > 0 {
			return matches[0]
		}
	}

	// Fallback: try relative to this file
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		baseDir := filepath.Dir(filename)
		path := filepath.Join(baseDir, "..", "..", "templates", "spec-steps", fmt.Sprintf("%s-*.md", stepID))
		matches, _ := filepath.Glob(path)
		if len(matches) > 0 {
			return matches[0]
		}
	}

	// Last resort: assume templates/ in cwd
	return filepath.Join("templates", "spec-steps", fmt.Sprintf("%s-*.md", stepID))
}

func getSpecSteps() []SpecStep {
	return []SpecStep{
		{ID: "a", Title: "Overview", Section: "overview"},
		{ID: "b", Title: "Requirements", Section: "requirements"},
		{ID: "c", Title: "Acceptance Criteria", Section: "acceptance_criteria"},
		{ID: "d", Title: "Constraints", Section: "constraints"},
		{ID: "e", Title: "Technical Approach", Section: "technical_approach"},
		{ID: "f", Title: "Success Metrics", Section: "success_metrics"},
		{ID: "g", Title: "Non-Goals", Section: "non_goals"},
		{ID: "h", Title: "Verification", Section: "verification"},
	}
}
