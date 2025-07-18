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

package hnsw

import (
	"fmt"
	"os"
	"strings"

	"github.com/weaviate/weaviate/entities/schema/config"
	vectorIndexCommon "github.com/weaviate/weaviate/entities/vectorindex/common"
)

const (
	// Set these defaults if the user leaves them blank
	DefaultCleanupIntervalSeconds = 5 * 60
	DefaultMaxConnections         = 32
	DefaultEFConstruction         = 128
	DefaultEF                     = -1 // indicates "let Weaviate pick"
	DefaultDynamicEFMin           = 100
	DefaultDynamicEFMax           = 500
	DefaultDynamicEFFactor        = 8
	DefaultSkip                   = false
	DefaultFlatSearchCutoff       = 40000

	FilterStrategySweeping = "sweeping"
	FilterStrategyAcorn    = "acorn"

	DefaultFilterStrategy = FilterStrategySweeping

	// Fail validation if those criteria are not met
	MinmumMaxConnections  = 4
	MaximumMaxConnections = 2047
	MinmumEFConstruction  = 4
)

// UserConfig bundles all values settable by a user in the per-class settings
type UserConfig struct {
	Skip                   bool              `json:"skip"`
	CleanupIntervalSeconds int               `json:"cleanupIntervalSeconds"`
	MaxConnections         int               `json:"maxConnections"`
	EFConstruction         int               `json:"efConstruction"`
	EF                     int               `json:"ef"`
	DynamicEFMin           int               `json:"dynamicEfMin"`
	DynamicEFMax           int               `json:"dynamicEfMax"`
	DynamicEFFactor        int               `json:"dynamicEfFactor"`
	VectorCacheMaxObjects  int               `json:"vectorCacheMaxObjects"`
	FlatSearchCutoff       int               `json:"flatSearchCutoff"`
	Distance               string            `json:"distance"`
	PQ                     PQConfig          `json:"pq"`
	BQ                     BQConfig          `json:"bq"`
	SQ                     SQConfig          `json:"sq"`
	RQ                     RQConfig          `json:"rq"`
	FilterStrategy         string            `json:"filterStrategy"`
	Multivector            MultivectorConfig `json:"multivector"`
}

// IndexType returns the type of the underlying vector index, thus making sure
// the schema.VectorIndexConfig interface is implemented
func (u UserConfig) IndexType() string {
	return "hnsw"
}

func (u UserConfig) DistanceName() string {
	return u.Distance
}

func (u UserConfig) IsMultiVector() bool {
	return u.Multivector.Enabled
}

// SetDefaults in the user-specifyable part of the config
func (u *UserConfig) SetDefaults() {
	u.MaxConnections = DefaultMaxConnections
	u.EFConstruction = DefaultEFConstruction
	u.CleanupIntervalSeconds = DefaultCleanupIntervalSeconds
	u.VectorCacheMaxObjects = vectorIndexCommon.DefaultVectorCacheMaxObjects
	u.EF = DefaultEF
	u.DynamicEFFactor = DefaultDynamicEFFactor
	u.DynamicEFMax = DefaultDynamicEFMax
	u.DynamicEFMin = DefaultDynamicEFMin
	u.Skip = DefaultSkip
	u.FlatSearchCutoff = DefaultFlatSearchCutoff
	u.Distance = vectorIndexCommon.DefaultDistanceMetric
	u.PQ = PQConfig{
		Enabled:        DefaultPQEnabled,
		BitCompression: DefaultPQBitCompression,
		Segments:       DefaultPQSegments,
		Centroids:      DefaultPQCentroids,
		TrainingLimit:  DefaultPQTrainingLimit,
		Encoder: PQEncoder{
			Type:         DefaultPQEncoderType,
			Distribution: DefaultPQEncoderDistribution,
		},
	}
	u.BQ = BQConfig{
		Enabled: DefaultBQEnabled,
	}
	u.SQ = SQConfig{
		Enabled:       DefaultSQEnabled,
		TrainingLimit: DefaultSQTrainingLimit,
		RescoreLimit:  DefaultSQRescoreLimit,
	}
	u.RQ = RQConfig{
		Enabled:      DefaultRQEnabled,
		Bits:         DefaultRQBits,
		RescoreLimit: DefaultRQRescoreLimit,
	}
	if strategy := os.Getenv("HNSW_DEFAULT_FILTER_STRATEGY"); strategy == FilterStrategyAcorn {
		u.FilterStrategy = FilterStrategyAcorn
	} else {
		u.FilterStrategy = FilterStrategySweeping
	}
	u.Multivector = MultivectorConfig{
		Aggregation: DefaultMultivectorAggregation,
		Enabled:     DefaultMultivectorEnabled,
		MuveraConfig: MuveraConfig{
			Enabled:      DefaultMultivectorMuveraEnabled,
			KSim:         DefaultMultivectorKSim,
			DProjections: DefaultMultivectorDProjections,
			Repetitions:  DefaultMultivectorRepetitions,
		},
	}
}

// ParseAndValidateConfig from an unknown input value, as this is not further
// specified in the API to allow of exchanging the index type
func ParseAndValidateConfig(input interface{}, isMultiVector bool) (config.VectorIndexConfig, error) {
	uc := UserConfig{}
	uc.SetDefaults()

	if input == nil {
		return uc, nil
	}

	asMap, ok := input.(map[string]interface{})
	if !ok || asMap == nil {
		return uc, fmt.Errorf("input must be a non-nil map")
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "maxConnections", func(v int) {
		uc.MaxConnections = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "cleanupIntervalSeconds", func(v int) {
		uc.CleanupIntervalSeconds = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "efConstruction", func(v int) {
		uc.EFConstruction = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "ef", func(v int) {
		uc.EF = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "dynamicEfFactor", func(v int) {
		uc.DynamicEFFactor = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "dynamicEfMax", func(v int) {
		uc.DynamicEFMax = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "dynamicEfMin", func(v int) {
		uc.DynamicEFMin = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "vectorCacheMaxObjects", func(v int) {
		uc.VectorCacheMaxObjects = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalIntFromMap(asMap, "flatSearchCutoff", func(v int) {
		uc.FlatSearchCutoff = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalBoolFromMap(asMap, "skip", func(v bool) {
		uc.Skip = v
	}); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalStringFromMap(asMap, "distance", func(v string) {
		uc.Distance = v
	}); err != nil {
		return uc, err
	}

	if err := parsePQMap(asMap, &uc.PQ); err != nil {
		return uc, err
	}

	if err := parseBQMap(asMap, &uc.BQ); err != nil {
		return uc, err
	}

	if err := parseSQMap(asMap, &uc.SQ); err != nil {
		return uc, err
	}

	if err := parseRQMap(asMap, &uc.RQ); err != nil {
		return uc, err
	}

	if err := vectorIndexCommon.OptionalStringFromMap(asMap, "filterStrategy", func(v string) {
		uc.FilterStrategy = v
	}); err != nil {
		return uc, err
	}

	if err := parseMultivectorMap(asMap, &uc.Multivector, isMultiVector); err != nil {
		return uc, err
	}

	return uc, uc.validate()
}

func (u *UserConfig) validate() error {
	var errMsgs []string
	if u.MaxConnections < MinmumMaxConnections {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"maxConnections must be a positive integer with a minimum of %d",
			MinmumMaxConnections,
		))
	}

	if u.MaxConnections > MaximumMaxConnections {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"maxConnections must be less than %d",
			MaximumMaxConnections+1,
		))
	}

	if u.EFConstruction < MinmumEFConstruction {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"efConstruction must be a positive integer with a minimum of %d",
			MinmumMaxConnections,
		))
	}

	if u.FilterStrategy != FilterStrategySweeping && u.FilterStrategy != FilterStrategyAcorn {
		errMsgs = append(errMsgs, "filterStrategy must be either 'sweeping' or 'acorn'")
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf("invalid hnsw config: %s",
			strings.Join(errMsgs, ", "))
	}

	enabled := 0
	if u.PQ.Enabled {
		enabled++
	}
	if u.BQ.Enabled {
		enabled++
	}
	if u.SQ.Enabled {
		enabled++
	}
	if u.RQ.Enabled {
		enabled++
	}
	if enabled > 1 {
		return fmt.Errorf("invalid hnsw config: more than a single compression methods enabled")
	}

	if u.Multivector.MuveraConfig.Enabled && u.Multivector.MuveraConfig.KSim > 10 {
		return fmt.Errorf("invalid hnsw config: ksim must be less than 10")
	}

	return nil
}

func NewDefaultUserConfig() UserConfig {
	uc := UserConfig{}
	uc.SetDefaults()
	return uc
}

func NewDefaultMultiVectorUserConfig() UserConfig {
	uc := UserConfig{}
	uc.SetDefaults()
	uc.Multivector = MultivectorConfig{Enabled: true}
	return uc
}
