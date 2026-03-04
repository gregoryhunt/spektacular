package spec

import (
	"fmt"
	"os"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/jumppad-labs/spektacular/internal/workflow"
	"github.com/jumppad-labs/spektacular/templates"
)

// Steps returns the ordered step configs for a spec workflow.
func Steps() []workflow.StepConfig {
	return []workflow.StepConfig{
		{Name: "overview", Src: []string{"new"}, Dst: "overview"},
		{Name: "requirements", Src: []string{"overview"}, Dst: "requirements"},
		{Name: "acceptance_criteria", Src: []string{"requirements"}, Dst: "acceptance_criteria"},
		{Name: "constraints", Src: []string{"acceptance_criteria"}, Dst: "constraints"},
		{Name: "technical_approach", Src: []string{"constraints"}, Dst: "technical_approach"},
		{Name: "success_metrics", Src: []string{"technical_approach"}, Dst: "success_metrics"},
		{Name: "non_goals", Src: []string{"success_metrics"}, Dst: "non_goals"},
		{Name: "verification", Src: []string{"non_goals"}, Dst: "verification"},
	}
}

// StepTitle converts a step name like "acceptance_criteria" to "Acceptance Criteria".
func StepTitle(name string) string {
	words := strings.Split(name, "_")
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

// RenderStep renders a step template with the standard variables.
func RenderStep(stepName, specPath, nextStepName, command string) (string, error) {
	tmplBytes, err := templates.FS.ReadFile("spec-steps/" + stepName + ".md")
	if err != nil {
		return "", fmt.Errorf("loading template for step %s: %w", stepName, err)
	}

	data := map[string]any{
		"step":      stepName,
		"title":     StepTitle(stepName),
		"section":   stepName,
		"spec_path": specPath,
		"next_step": nextStepName,
		"command":   command,
	}

	out, err := mustache.Render(string(tmplBytes), data)
	if err != nil {
		return "", fmt.Errorf("rendering template for step %s: %w", stepName, err)
	}
	return out, nil
}

// RenderScaffold writes the spec scaffold template to specPath.
func RenderScaffold(specPath, name string) error {
	tmplBytes, err := templates.FS.ReadFile("spec-scaffold.md")
	if err != nil {
		return fmt.Errorf("reading scaffold template: %w", err)
	}

	rendered, err := mustache.Render(string(tmplBytes), map[string]string{"name": name})
	if err != nil {
		return fmt.Errorf("rendering scaffold template: %w", err)
	}

	return os.WriteFile(specPath, []byte(rendered), 0644)
}

