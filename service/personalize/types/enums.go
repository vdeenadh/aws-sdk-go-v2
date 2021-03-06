// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type RecipeProvider string

// Enum values for RecipeProvider
const (
	RecipeProviderService RecipeProvider = "SERVICE"
)

// Values returns all known values for RecipeProvider. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecipeProvider) Values() []RecipeProvider {
	return []RecipeProvider{
		"SERVICE",
	}
}

type TrainingMode string

// Enum values for TrainingMode
const (
	TrainingModeFull   TrainingMode = "FULL"
	TrainingModeUpdate TrainingMode = "UPDATE"
)

// Values returns all known values for TrainingMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TrainingMode) Values() []TrainingMode {
	return []TrainingMode{
		"FULL",
		"UPDATE",
	}
}
