package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"rm/common"
	"rm/utils/apollo"
)

var (
	sampleNameRe  *regexp.Regexp
	runPlanNameRe *regexp.Regexp
)

func init() {
	sampleNameRe, _ = regexp.Compile(`^[0-9A-Za-z-]{1,20}$`)
	runPlanNameRe, _ = regexp.Compile(`^[A-Za-z0-9-]{1,20}$`)
}

var (
	Validators = map[string]ICustomValidator{
		"sampleNameValidator":            newCustomValidator("sampleNameValidator", sampleInfoNameValidator),
		"platformKeyValidator":           newCustomValidator("platformKeyValidator", platformKeyValidator),
		"index1LengthKeyValidator":       newCustomValidator("index1LengthKeyValidator", index1LengthKeyValidator),
		"index2LengthKeyValidator":       newCustomValidator("index2LengthKeyValidator", index2LengthKeyValidator),
		"runPlanNameValidator":           newCustomValidator("runPlanNameValidator", runPlanNameValidator),
		"runPlanTypeValidator":           newCustomValidator("runPlanTypeValidator", runPlanTypeValidator),
		"runPlanThroughputModeValidator": newCustomValidator("runPlanThroughputModeValidator", runPlanThroughputModeValidator),
		"runPlanReadlengthModeValidator": newCustomValidator("runPlanReadlengthModeValidator", runPlanReadlengthModeValidator),
		"runPlanReferenceValidator":      newCustomValidator("runPlanReferenceValidator", runPlanReferenceValidator),
		"runPlanAnalysisValidator":       newCustomValidator("runPlanAnalysisValidator", runPlanAnalysisValidator),
	}
)

func sampleInfoNameValidator(fl validator.FieldLevel) bool {
	if name, ok := fl.Field().Interface().(string); ok {
		if sampleNameRe.MatchString(name) {
			return true
		}
	}
	return false
}

func platformKeyValidator(fl validator.FieldLevel) bool {
	if key, ok := fl.Field().Interface().(int32); ok {
		if _, ok1 := apollo.BKCommonConfig.GetPlatform(key); ok1 {
			return true
		}
	}
	return false
}

func index1LengthKeyValidator(fl validator.FieldLevel) bool {
	if key, ok := fl.Field().Interface().(int32); ok {
		if _, ok1 := apollo.BKCommonConfig.GetIndex1(key); ok1 || key == 0 {
			return true
		}
	}
	return false
}

func index2LengthKeyValidator(fl validator.FieldLevel) bool {
	if key, ok := fl.Field().Interface().(int32); ok {
		if _, ok1 := apollo.BKCommonConfig.GetIndex2(key); ok1 || key == 0 {
			return true
		}
	}
	return false
}

func runPlanNameValidator(fl validator.FieldLevel) bool {
	if name, ok := fl.Field().Interface().(string); ok {
		if runPlanNameRe.MatchString(name) {
			return true
		}
	}
	return false
}

func runPlanTypeValidator(fl validator.FieldLevel) bool {
	if name, ok := fl.Field().Interface().(int32); ok {
		if name == int32(common.NormalPlan) || name == int32(common.CustomPlan) {
			return true
		}
	}
	return false
}

func runPlanThroughputModeValidator(fl validator.FieldLevel) bool {
	if key, ok := fl.Field().Interface().(int32); ok {
		if _, ok1 := apollo.BKRunPlanConfig.GetThroughputDefinition(key); ok1 {
			return true
		}
	}
	return false
}

func runPlanReadlengthModeValidator(fl validator.FieldLevel) bool {
	if key, ok := fl.Field().Interface().(int32); ok {
		if _, ok1 := apollo.BKRunPlanConfig.GetReadlengthModeDefinition(key); ok1 {
			return true
		}
	}
	return false
}

func runPlanReferenceValidator(fl validator.FieldLevel) bool {
	if key, ok := fl.Field().Interface().(int32); ok {
		if _, ok1 := apollo.BKRunPlanConfig.GetReference(key); ok1 || key == 0 {
			return true
		}
	}
	return false
}

func runPlanAnalysisValidator(fl validator.FieldLevel) bool {
	if name, ok := fl.Field().Interface().(string); ok {
		if _, ok1 := common.MetadataConfig.GetAnalysis(name); ok1 {
			return true
		}
	}
	return false
}
