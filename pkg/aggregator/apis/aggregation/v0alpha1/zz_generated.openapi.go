//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneService":          schema_aggregator_apis_aggregation_v0alpha1_DataPlaneService(ref),
		"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceCondition": schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceCondition(ref),
		"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceList":      schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceList(ref),
		"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceSpec":      schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceSpec(ref),
		"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceStatus":    schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceStatus(ref),
		"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.Service":                   schema_aggregator_apis_aggregation_v0alpha1_Service(ref),
	}
}

func schema_aggregator_apis_aggregation_v0alpha1_DataPlaneService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceSpec", "github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataPlaneServiceCondition describes the state of an DataPlaneService at a particular point",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of the condition.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is the status of the condition. Can be True, False, Unknown.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "Unique, one-word, CamelCase reason for the condition's last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human-readable message indicating details about last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneService"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneService", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"services": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type":       "map",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Services is a list of services that the proxied service provides.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.Service"),
									},
								},
							},
						},
					},
				},
				Required: []string{"group", "version"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.Service"},
	}
}

func schema_aggregator_apis_aggregation_v0alpha1_DataPlaneServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DataPlaneServiceStatus contains derived information about a remote DataPlaneService.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type":       "map",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Current service state of DataPlaneService.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceCondition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1.DataPlaneServiceCondition"},
	}
}

func schema_aggregator_apis_aggregation_v0alpha1_Service(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Service defines the type of service the proxied service provides.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Possible enum values:\n - `\"admission\"`\n - `\"customroute\"`\n - `\"data\"`\n - `\"diagnostics\"`\n - `\"stream\"`\n - `\"subresource\"`",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"admission", "customroute", "data", "diagnostics", "stream", "subresource"},
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource is used by the SubResourceServiceType to specify the resource to proxy.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path is used by the CustomRouteServiceType and SubResourceServiceType to specify the path to the endpoint.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type"},
			},
		},
	}
}
