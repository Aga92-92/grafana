package expressions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kube-openapi/pkg/validation/spec"

	query "github.com/grafana/grafana/pkg/apis/query/v0alpha1"
	"github.com/grafana/grafana/pkg/apis/query/v0alpha1/definition"
)

func TestParseQueriesIntoQueryDataRequest(t *testing.T) {
	ref := func(path string) spec.Ref {
		return spec.Ref{} // dummy ref
	}
	typesFile := "types.json"
	current := GetQueryTypeDefinitions(ref)

	created := time.Now().UTC()
	resourceVersion := fmt.Sprintf("%d", created.UnixMilli())
	defs := definition.QueryTypeDefinitionList{
		TypeMeta: v1.TypeMeta{
			Kind:       "QueryTypeDefinitionList",
			APIVersion: query.APIVERSION,
		},
	}
	existing := definition.QueryTypeDefinitionList{}
	body, err := os.ReadFile(typesFile)
	if err == nil {
		_ = json.Unmarshal(body, &existing)
		defs.ListMeta = existing.ListMeta
	}

	// Check for changes in any existing values
	for _, item := range existing.Items {
		v, ok := current[item.Name]
		if ok {
			delete(current, item.Name)
			a, e1 := json.Marshal(v)
			b, e2 := json.Marshal(item.Spec)
			if e1 != nil || e2 != nil || !bytes.Equal(a, b) {
				item.ResourceVersion = resourceVersion
				defs.ListMeta.ResourceVersion = resourceVersion
				if item.Annotations == nil {
					item.Annotations = make(map[string]string)
				}
				item.Annotations["grafana.app/modifiedTime"] = created.Format(time.RFC3339)
				item.Spec = v // the current value
			}
			defs.Items = append(defs.Items, item)
		} else {
			defs.ListMeta.ResourceVersion = resourceVersion
		}
	}

	// New items added to the list
	for k, v := range current {
		defs.ListMeta.ResourceVersion = resourceVersion
		defs.Items = append(defs.Items, definition.QueryTypeDefinition{
			TypeMeta: v1.TypeMeta{
				Kind: "QueryTypeDefinition",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:              k,
				CreationTimestamp: v1.NewTime(created),
				ResourceVersion:   resourceVersion,
			},
			Spec: v,
		})
	}

	out, err := json.MarshalIndent(defs, "", "  ")
	require.NoError(t, err)
	fmt.Printf("%s\n", out)

	err = os.WriteFile(typesFile, out, 0644)
	require.NoError(t, err, "error writing file")
}
