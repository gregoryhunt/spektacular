package plan

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/jumppad-labs/spektacular/internal/store"
	"github.com/jumppad-labs/spektacular/internal/workflow"
	"github.com/jumppad-labs/spektacular/templates"
)

// PlanFilePath returns the store-relative path for a plan file.
func PlanFilePath(name string) string {
	return "plans/" + name + "/plan.md"
}

// ContextFilePath returns the store-relative path for a plan's context file.
func ContextFilePath(name string) string {
	return "plans/" + name + "/context.md"
}

// ResearchFilePath returns the store-relative path for a plan's research file.
func ResearchFilePath(name string) string {
	return "plans/" + name + "/research.md"
}

// Steps returns the ordered step configs for a plan workflow.
func Steps() []workflow.StepConfig {
	return []workflow.StepConfig{
		{Name: "new", Src: []string{"start"}, Dst: "new", Callback: new()},
		{Name: "overview", Src: []string{"new"}, Dst: "overview", Callback: overview()},
		{Name: "discovery", Src: []string{"overview"}, Dst: "discovery", Callback: discovery()},
		{Name: "architecture", Src: []string{"discovery"}, Dst: "architecture", Callback: architecture()},
		{Name: "components", Src: []string{"architecture"}, Dst: "components", Callback: components()},
		{Name: "data_structures", Src: []string{"components"}, Dst: "data_structures", Callback: dataStructures()},
		{Name: "implementation_detail", Src: []string{"data_structures"}, Dst: "implementation_detail", Callback: implementationDetail()},
		{Name: "dependencies", Src: []string{"implementation_detail"}, Dst: "dependencies", Callback: dependencies()},
		{Name: "testing_approach", Src: []string{"dependencies"}, Dst: "testing_approach", Callback: testingApproach()},
		{Name: "milestones", Src: []string{"testing_approach"}, Dst: "milestones", Callback: milestones()},
		{Name: "phases", Src: []string{"milestones"}, Dst: "phases", Callback: phases()},
		{Name: "open_questions", Src: []string{"phases"}, Dst: "open_questions", Callback: openQuestions()},
		{Name: "out_of_scope", Src: []string{"open_questions"}, Dst: "out_of_scope", Callback: outOfScope()},
		{Name: "verification", Src: []string{"out_of_scope"}, Dst: "verification", Callback: verification()},
		{Name: "write_plan", Src: []string{"verification"}, Dst: "write_plan", Callback: writePlan()},
		{Name: "write_context", Src: []string{"write_plan"}, Dst: "write_context", Callback: writeContext()},
		{Name: "write_research", Src: []string{"write_context"}, Dst: "write_research", Callback: writeResearch()},
		{Name: "finished", Src: []string{"write_research"}, Dst: "finished", Callback: finished()},
	}
}

// new initializes state only — no document created yet.
func new() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "overview", nil
	}
}

func overview() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("overview", "discovery", "steps/plan/01-overview.md", data, out, st, cfg)
	}
}

func discovery() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("discovery", "architecture", "steps/plan/02-discovery.md", data, out, st, cfg)
	}
}

func architecture() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("architecture", "components", "steps/plan/03-architecture.md", data, out, st, cfg)
	}
}

func components() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("components", "data_structures", "steps/plan/04-components.md", data, out, st, cfg)
	}
}

func dataStructures() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("data_structures", "implementation_detail", "steps/plan/05-data_structures.md", data, out, st, cfg)
	}
}

func implementationDetail() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("implementation_detail", "dependencies", "steps/plan/06-implementation_detail.md", data, out, st, cfg)
	}
}

func dependencies() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("dependencies", "testing_approach", "steps/plan/07-dependencies.md", data, out, st, cfg)
	}
}

func testingApproach() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("testing_approach", "milestones", "steps/plan/08-testing_approach.md", data, out, st, cfg)
	}
}

func milestones() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("milestones", "phases", "steps/plan/09-milestones.md", data, out, st, cfg)
	}
}

func phases() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("phases", "open_questions", "steps/plan/10-phases.md", data, out, st, cfg)
	}
}

func openQuestions() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("open_questions", "out_of_scope", "steps/plan/11-open_questions.md", data, out, st, cfg)
	}
}

func outOfScope() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		return "", writeStepResult("out_of_scope", "verification", "steps/plan/12-out_of_scope.md", data, out, st, cfg)
	}
}

func verification() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		planName := getString(data, "name")
		planScaffold, err := renderTemplate("scaffold/plan.md", map[string]any{"name": planName})
		if err != nil {
			return "", err
		}
		contextScaffold, err := renderTemplate("scaffold/context.md", map[string]any{"name": planName})
		if err != nil {
			return "", err
		}
		researchScaffold, err := renderTemplate("scaffold/research.md", map[string]any{"name": planName})
		if err != nil {
			return "", err
		}
		return "", writeStepResult("verification", "finished", "steps/plan/13-verification.md", data, out, st, cfg,
			map[string]any{
				"plan_template":     planScaffold,
				"context_template":  contextScaffold,
				"research_template": researchScaffold,
			})
	}
}

// writePlan writes plan.md from the plan_template data key.
func writePlan() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		planName := getString(data, "name")
		content := getString(data, "plan_template")
		if content == "" {
			return "", fmt.Errorf("plan_template missing — pipe the filled plan.md via --stdin plan_template")
		}
		if !cfg.DryRun {
			if err := st.Write(PlanFilePath(planName), []byte(content)); err != nil {
				return "", err
			}
		}
		return "", writeStepResult("write_plan", "write_context", "steps/plan/14-write_plan.md", data, out, st, cfg)
	}
}

// writeContext writes context.md from the context_template data key.
func writeContext() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		planName := getString(data, "name")
		content := getString(data, "context_template")
		if content == "" {
			return "", fmt.Errorf("context_template missing — pipe the filled context.md via --stdin context_template")
		}
		if !cfg.DryRun {
			if err := st.Write(ContextFilePath(planName), []byte(content)); err != nil {
				return "", err
			}
		}
		return "", writeStepResult("write_context", "write_research", "steps/plan/15-write_context.md", data, out, st, cfg)
	}
}

// writeResearch writes research.md from the research_template data key.
func writeResearch() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		planName := getString(data, "name")
		content := getString(data, "research_template")
		if content == "" {
			return "", fmt.Errorf("research_template missing — pipe the filled research.md via --stdin research_template")
		}
		if !cfg.DryRun {
			if err := st.Write(ResearchFilePath(planName), []byte(content)); err != nil {
				return "", err
			}
		}
		return "", writeStepResult("write_research", "finished", "steps/plan/16-write_research.md", data, out, st, cfg)
	}
}

func finished() workflow.StepCallback {
	return func(data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config) (string, error) {
		if cfg.DryRun {
			return "", writeStepResult("finished", "", "steps/plan/17-finished.md", data, out, st, cfg)
		}
		planName := getString(data, "name")
		for _, p := range []string{PlanFilePath(planName), ContextFilePath(planName), ResearchFilePath(planName)} {
			if !st.Exists(p) {
				return "", fmt.Errorf("expected file %s not found — the preceding write step should have written it", p)
			}
		}
		return "", writeStepResult("finished", "", "steps/plan/17-finished.md", data, out, st, cfg)
	}
}

// writeStepResult renders the step template and writes the result to out.
func writeStepResult(name, nextStep, templatePath string, data workflow.Data, out workflow.ResultWriter, st store.Store, cfg workflow.Config, extra ...map[string]any) error {
	planName := getString(data, "name")
	planPath := filepath.Join(st.Root(), PlanFilePath(planName))
	contextPath := filepath.Join(st.Root(), ContextFilePath(planName))
	researchPath := filepath.Join(st.Root(), ResearchFilePath(planName))
	planDir := filepath.Dir(planPath)
	specPath := filepath.Join(st.Root(), "specs", planName+".md")

	vars := map[string]any{
		"step":          name,
		"title":         stepTitle(name),
		"plan_path":     planPath,
		"context_path":  contextPath,
		"research_path": researchPath,
		"plan_dir":      planDir,
		"plan_name":     planName,
		"spec_path":     specPath,
		"next_step":     nextStep,
		"config":        map[string]any{"command": cfg.Command},
	}
	for _, m := range extra {
		for k, v := range m {
			vars[k] = v
		}
	}

	instruction, err := renderTemplate(templatePath, vars)
	if err != nil {
		return err
	}
	return out.WriteResult(Result{
		Step:        name,
		PlanPath:    planPath,
		PlanName:    planName,
		Instruction: instruction,
	})
}

func getString(data workflow.Data, key string) string {
	v, ok := data.Get(key)
	if !ok {
		return ""
	}
	s, _ := v.(string)
	return s
}

func renderTemplate(templatePath string, data map[string]any) (string, error) {
	tmplBytes, err := templates.FS.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("loading template %s: %w", templatePath, err)
	}
	return mustache.Render(string(tmplBytes), data)
}

func stepTitle(name string) string {
	words := strings.Split(name, "_")
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}
