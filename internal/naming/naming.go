package naming

import (
	"strings"

	"github.com/jinzhu/inflection"
)

var (
	// Replace all occurrences of these strings in the property name.
	propertyNameReplacements = map[string]string{
		"CloudFormation": "Cloudformation",
		"CloudFront":     "Cloudfront",
		"CloudWatch":     "Cloudwatch",
		"FSx":            "Fsx",
	}
)

// CloudFormationPropertyToTerraformAttribute converts a CloudFormation property name to a Terraform attribute name.
// For example `GlobalReplicationGroupDescription` -> `global_replication_group_description`.
func CloudFormationPropertyToTerraformAttribute(propertyName string) string {
	propertyName = strings.TrimSpace(propertyName)

	if propertyName == "" {
		return propertyName
	}

	for old, new := range propertyNameReplacements {
		propertyName = strings.ReplaceAll(propertyName, old, new)
	}

	attributeName := strings.Builder{}

	for i, ch := range []byte(propertyName) {
		isCap := isCapitalLetter(ch)
		isLow := isLowercaseLetter(ch)
		isDig := isNumeric(ch)

		if isCap {
			ch = toLowercaseLetter(ch)
		}

		if i < len(propertyName)-1 {
			nextCh := propertyName[i+1]
			nextIsCap := isCapitalLetter(nextCh)
			nextIsLow := isLowercaseLetter(nextCh)
			nextIsDig := isNumeric(nextCh)

			// Append underscore if case changes.
			if (isCap && nextIsLow) || (isLow && (nextIsCap || nextIsDig) || (isDig && (nextIsCap || nextIsLow))) {
				if isCap && nextIsLow {
					if prevIsCap := i > 0 && isCapitalLetter(propertyName[i-1]); prevIsCap {
						attributeName.WriteByte('_')
					}
				}
				attributeName.WriteByte(ch)
				if isLow || isDig {
					attributeName.WriteByte('_')
				}

				continue
			}
		}

		if isCap || isLow || isDig {
			attributeName.WriteByte(ch)
		} else {
			attributeName.WriteByte('_')
		}
	}

	return attributeName.String()
}

// Pluralize converts a name to its plural form.
// The inflection package is used as a first attempt to pluralize names,
// but exceptions to the rule are handled as follows:
//  - 'es' is appended to a name ending in 's'
//  - 's' is appended to a name ending in a number
func Pluralize(name string) string {
	if name == "" {
		return name
	}

	// Custom Rules
	inflection.AddIrregular("lens", "lenses")          // "lens" => "lenses"
	inflection.AddPlural("((e|n)f)s$", "${1}s_plural") // "efs" => "efs_plural" or "nfs" => "nfs_plural"
	inflection.AddPlural("(tion)s$", "${1}s_plural")   // "associations" => "associations_plural"
	inflection.AddPlural("(window)s$", "${1}s_plural") // "windows" => "windows_plural"

	pluralName := inflection.Plural(name)

	if pluralName == name {
		arr := []byte(pluralName)
		lastChar := arr[len(arr)-1]

		if isNumeric(lastChar) {
			pluralName += "s" // "s3" => "s3s"
		}
	}

	return pluralName
}

func isCapitalLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func isLowercaseLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func toLowercaseLetter(ch byte) byte {
	ch += 'a'
	ch -= 'A'
	return ch
}
