//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/test/helper"
	"github.com/weaviate/weaviate/test/helper/sample-schema/companies"
)

func testText2VecGoogle(rest, grpc, gcpProject, vectorizerName string) func(t *testing.T) {
	return func(t *testing.T) {
		helper.SetupClient(rest)
		// Data
		data := companies.Companies
		className := "VectorizerTest"
		class := companies.BaseClass(className)
		tests := []struct {
			name  string
			model string
		}{
			{
				name:  "textembedding-gecko@latest",
				model: "textembedding-gecko@latest",
			},
			{
				name:  "textembedding-gecko-multilingual@latest",
				model: "textembedding-gecko-multilingual@latest",
			},
			{
				name:  "text-embedding-005",
				model: "text-embedding-005",
			},
			{
				name:  "text-multilingual-embedding-002",
				model: "text-multilingual-embedding-002",
			},
			{
				name:  "gemini-embedding-001",
				model: "gemini-embedding-001",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Define class
				class.VectorConfig = map[string]models.VectorConfig{
					"description": {
						Vectorizer: map[string]interface{}{
							vectorizerName: map[string]interface{}{
								"properties":         []interface{}{"description"},
								"vectorizeClassName": false,
								"projectId":          gcpProject,
								"modelId":            tt.model,
							},
						},
						VectorIndexType: "flat",
					},
				}
				// create schema
				helper.CreateClass(t, class)
				defer helper.DeleteClass(t, class.Class)
				// create objects
				t.Run("create objects", func(t *testing.T) {
					companies.InsertObjects(t, rest, class.Class)
				})
				t.Run("check objects existence", func(t *testing.T) {
					for _, company := range data {
						t.Run(company.ID.String(), func(t *testing.T) {
							obj, err := helper.GetObject(t, class.Class, company.ID, "vector")
							require.NoError(t, err)
							require.NotNil(t, obj)
							require.Len(t, obj.Vectors, 1)
							require.IsType(t, []float32{}, obj.Vectors["description"])
							assert.True(t, len(obj.Vectors["description"].([]float32)) > 0)
						})
					}
				})
				t.Run("search tests", func(t *testing.T) {
					companies.PerformAllSearchTests(t, rest, grpc, class.Class)
				})
			})
		}
	}
}
