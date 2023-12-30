package types

const (
	COVERAGE_TYPE_UNCOVERED_NEW_CODE       = "uncovered_new_code"
	COVERAGE_TYPE_GAINED_COVERAGE_NEW_CODE = "gained_coverage_new_code"

	COVERAGE_TYPE_LOST_BASELINE_COVERAGE   = "lost_baseline_coverage"
	COVERAGE_TYPE_GAINED_BASELINE_COVERAGE = "gained_baseline_coverage"

	COVERAGE_TYPE_UNCOVERED_INCLUDED_CODE       = "uncovered_included_code"
	COVERAGE_TYPE_GAINED_COVERAGE_INCLUDED_CODE = "gained_coverage_included_code"

	COVERAGE_TYPE_EXCLUDED_UNCOVERED_BASELINE_CODE = "excluded_uncovered_baseline_code"
	COVERAGE_TYPE_EXCLUDED_COVERED_BASELINE_CODE   = "excluded_covered_baseline_code"

	COVERAGE_TYPE_DELETED_UNCOVERED_BASELINE_CODE = "deleted_uncovered_baseline_code"
	COVERAGE_TYPE_DELETED_COVERED_BASELINE_CODE   = "deleted_covered_baseline_code"
)

var (
	COVERAGE_TYPES = []string{
		COVERAGE_TYPE_UNCOVERED_NEW_CODE,
		COVERAGE_TYPE_LOST_BASELINE_COVERAGE,
		COVERAGE_TYPE_UNCOVERED_INCLUDED_CODE,
		COVERAGE_TYPE_GAINED_BASELINE_COVERAGE,
		COVERAGE_TYPE_GAINED_COVERAGE_INCLUDED_CODE,
		COVERAGE_TYPE_GAINED_COVERAGE_NEW_CODE,
		COVERAGE_TYPE_EXCLUDED_UNCOVERED_BASELINE_CODE,
		COVERAGE_TYPE_EXCLUDED_COVERED_BASELINE_CODE,
		COVERAGE_TYPE_DELETED_UNCOVERED_BASELINE_CODE,
		COVERAGE_TYPE_DELETED_COVERED_BASELINE_CODE,
	}
)
