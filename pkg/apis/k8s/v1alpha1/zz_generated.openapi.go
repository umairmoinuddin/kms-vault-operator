// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.KMSVaultSecret":       schema_pkg_apis_k8s_v1alpha1_KMSVaultSecret(ref),
		"github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.KMSVaultSecretSpec":   schema_pkg_apis_k8s_v1alpha1_KMSVaultSecretSpec(ref),
		"github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.KMSVaultSecretStatus": schema_pkg_apis_k8s_v1alpha1_KMSVaultSecretStatus(ref),
	}
}

func schema_pkg_apis_k8s_v1alpha1_KMSVaultSecret(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KMSVaultSecret is the Schema for the kmsvaultsecrets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.KMSVaultSecretSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.KMSVaultSecretStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.KMSVaultSecretSpec", "github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.KMSVaultSecretStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_k8s_v1alpha1_KMSVaultSecretSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KMSVaultSecretSpec defines the desired state of KMSVaultSecret",
				Properties: map[string]spec.Schema{
					"path": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.Secret"),
						},
					},
				},
				Required: []string{"path", "secret"},
			},
		},
		Dependencies: []string{
			"github.com/patoarvizu/kms-vault-operator/pkg/apis/k8s/v1alpha1.Secret"},
	}
}

func schema_pkg_apis_k8s_v1alpha1_KMSVaultSecretStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KMSVaultSecretStatus defines the observed state of KMSVaultSecret",
				Properties: map[string]spec.Schema{
					"created": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
