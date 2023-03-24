package yaml

import (
	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestOnBranchProtectionRuleCondition(t *testing.T) {
	tests := []struct {
		yaml string
		want OnConditions
	}{
		{
			yaml: `
  branch_protection_rule:
    types: [created, deleted]
`,
			want: OnConditions{
				BranchProtectionRule: &BranchProtectionRuleCondition{
					Types: []string{"created", "deleted"},
				},
			},
		},
	}

	for i, test := range tests {
		got := new(OnConditions)
		if err := yaml.Unmarshal([]byte(test.yaml), got); err != nil {
			t.Log(test.yaml)
			t.Error(err)
			return
		}
		if diff := cmp.Diff(got, &test.want); diff != "" {
			t.Log(test.yaml)
			t.Errorf("Unexpected parsing results for test %v", i)
			t.Log(diff)
		}
	}
}

func TestOnBranchProtectionRuleCondition_Error(t *testing.T) {
	err := yaml.Unmarshal([]byte("[[]]"), new(Concurrency))
	if err == nil || err.Error() != "failed to unmarshal concurrency" {
		t.Errorf("Expect error, got %s", err)
	}
}
