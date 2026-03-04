package spec

import "fmt"

// SpecStep defines a step in the spec creation workflow
type SpecStep struct {
	ID      string // a-h
	Title   string
	Section string
	Prompt  string
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

func getSpecSteps() []SpecStep {
	return []SpecStep{
		{
			ID:      "a",
			Title:   "Overview",
			Section: "overview",
			Prompt:  `Ask the user to describe this feature in 2-3 sentences:
• What is being built?
• What problem does it solve?
• Who benefits?

Write their response to the Overview section of the spec file.
Be specific — avoid generic phrases like 'improve the experience'.`,
		},
		{
			ID:      "b",
			Title:   "Requirements",
			Section: "requirements",
			Prompt: `Ask the user to list the specific, testable behaviours this feature must deliver.

Use active voice:
• 'Users can...'
• 'The system must...'

Each item should be independently verifiable. One behaviour per line.

Format the requirements as a markdown checklist and write them to the Requirements section:
- [ ] **Title** — description`,
		},
		{
			ID:      "c",
			Title:   "Acceptance Criteria",
			Section: "acceptance_criteria",
			Prompt: `Read all requirements from the spec file Requirements section.

For each requirement, ask the user: "What is the pass/fail condition that proves this is done?"

A good criterion:
• Describes an observable outcome
• Passes or fails — no subjective judgment
• Is traceable to this requirement

Example: "When X happens, Y is visible / saved / returned."

Write all criteria to the Acceptance Criteria section as a checklist.`,
		},
		{
			ID:      "d",
			Title:   "Constraints",
			Section: "constraints",
			Prompt: `Ask the user: Are there any hard constraints or boundaries the solution must operate within?

Examples:
• Must integrate with the existing authentication system
• Cannot introduce breaking changes to the public API
• Must support the current minimum supported runtime versions

Write their response to the Constraints section. If blank, write 'None.'`,
		},
		{
			ID:      "e",
			Title:   "Technical Approach",
			Section: "technical_approach",
			Prompt: `Ask the user: Do you have any technical direction already decided?

Examples:
• Key architectural decisions already made
• Preferred patterns or technologies
• Integration points with existing systems
• Known risks or areas of uncertainty

Write their response to the Technical Approach section. If blank, write 'None.'`,
		},
		{
			ID:      "f",
			Title:   "Success Metrics",
			Section: "success_metrics",
			Prompt: `Ask the user: How will you know this feature is working well after delivery?

Be specific:
• Quantitative: 'p99 latency < 200ms', 'error rate < 0.1%'
• Behavioral: 'users complete the flow without support intervention'

Write their response to the Success Metrics section. If blank, write 'None.'`,
		},
		{
			ID:      "g",
			Title:   "Non-Goals",
			Section: "non_goals",
			Prompt: `Ask the user: What is explicitly OUT of scope for this feature?

Examples:
• 'Mobile support is out of scope (tracked in #456)'
• 'Internationalisation will be addressed in a follow-up spec'

Write their response to the Non-Goals section. If blank, write 'None.'`,
		},
		{
			ID:      "h",
			Title:   "Verification",
			Section: "verification",
			Prompt: `Read the spec file in full.

Validate every section (Overview, Requirements, Acceptance Criteria, Constraints, Technical Approach, Success Metrics, Non-Goals) for:
• Completeness — all sections are filled
• Clarity — requirements are specific and testable
• Consistency — sections reference each other appropriately

Report any issues found. If there are gaps or unclear sections, ask the user for clarification.

Once all sections are validated and complete, confirm the spec is ready.`,
		},
	}
}
