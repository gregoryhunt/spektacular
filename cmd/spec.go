package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jumppad-labs/spektacular/internal/spec"
	"github.com/spf13/cobra"
)

var specCmd = &cobra.Command{
	Use:   "spec --name <name> --step <a-h>",
	Short: "Get instructions for a spec workflow step (a=overview, b=requirements, ...h=verify)",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		stepStr, _ := cmd.Flags().GetString("step")

		if name == "" || stepStr == "" {
			return fmt.Errorf("--name and --step are required")
		}

		if len(stepStr) != 1 || stepStr[0] < 'a' || stepStr[0] > 'h' {
			return fmt.Errorf("step must be a-h (a=overview, b=requirements, c=acceptance criteria, d=constraints, e=technical approach, f=success metrics, g=non-goals, h=verification)")
		}

		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting working directory: %w", err)
		}

		// Format spec path
		specPath := filepath.Join(cwd, ".spektacular", "specs", name+".md")

		// Render the template with mustache variables
		instruction, err := spec.RenderStepTemplate(stepStr, specPath)
		if err != nil {
			return fmt.Errorf("rendering step template: %w", err)
		}

		// Get step metadata for JSON output
		step, err := spec.GetSpecStepByID(stepStr)
		if err != nil {
			return err
		}

		// Output as JSON
		output := map[string]interface{}{
			"step":         stepStr,
			"section":      step.Section,
			"total_steps":  8,
			"spec_path":    specPath,
			"spec_name":    name,
			"instruction":  instruction,
		}

		jsonBytes, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return fmt.Errorf("marshaling output: %w", err)
		}

		fmt.Println(string(jsonBytes))
		return nil
	},
}

func init() {
	specCmd.Flags().StringP("name", "n", "", "Spec name (required)")
	specCmd.Flags().StringP("step", "s", "", "Step letter a-h (required)")
	specCmd.MarkFlagRequired("name")
	specCmd.MarkFlagRequired("step")
}
