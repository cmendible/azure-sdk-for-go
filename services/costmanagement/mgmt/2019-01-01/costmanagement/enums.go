package costmanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ExecutionStatus enumerates the values for execution status.
type ExecutionStatus string

const (
	// Completed ...
	Completed ExecutionStatus = "Completed"
	// DataNotAvailable ...
	DataNotAvailable ExecutionStatus = "DataNotAvailable"
	// Failed ...
	Failed ExecutionStatus = "Failed"
	// InProgress ...
	InProgress ExecutionStatus = "InProgress"
	// NewDataNotAvailable ...
	NewDataNotAvailable ExecutionStatus = "NewDataNotAvailable"
	// Queued ...
	Queued ExecutionStatus = "Queued"
	// Timeout ...
	Timeout ExecutionStatus = "Timeout"
)

// PossibleExecutionStatusValues returns an array of possible values for the ExecutionStatus const type.
func PossibleExecutionStatusValues() []ExecutionStatus {
	return []ExecutionStatus{Completed, DataNotAvailable, Failed, InProgress, NewDataNotAvailable, Queued, Timeout}
}

// ExecutionType enumerates the values for execution type.
type ExecutionType string

const (
	// OnDemand ...
	OnDemand ExecutionType = "OnDemand"
	// Scheduled ...
	Scheduled ExecutionType = "Scheduled"
)

// PossibleExecutionTypeValues returns an array of possible values for the ExecutionType const type.
func PossibleExecutionTypeValues() []ExecutionType {
	return []ExecutionType{OnDemand, Scheduled}
}

// FormatType enumerates the values for format type.
type FormatType string

const (
	// Csv ...
	Csv FormatType = "Csv"
)

// PossibleFormatTypeValues returns an array of possible values for the FormatType const type.
func PossibleFormatTypeValues() []FormatType {
	return []FormatType{Csv}
}

// GranularityType enumerates the values for granularity type.
type GranularityType string

const (
	// Daily ...
	Daily GranularityType = "Daily"
	// Hourly ...
	Hourly GranularityType = "Hourly"
)

// PossibleGranularityTypeValues returns an array of possible values for the GranularityType const type.
func PossibleGranularityTypeValues() []GranularityType {
	return []GranularityType{Daily, Hourly}
}

// QueryColumnType enumerates the values for query column type.
type QueryColumnType string

const (
	// QueryColumnTypeDimension ...
	QueryColumnTypeDimension QueryColumnType = "Dimension"
	// QueryColumnTypeTag ...
	QueryColumnTypeTag QueryColumnType = "Tag"
)

// PossibleQueryColumnTypeValues returns an array of possible values for the QueryColumnType const type.
func PossibleQueryColumnTypeValues() []QueryColumnType {
	return []QueryColumnType{QueryColumnTypeDimension, QueryColumnTypeTag}
}

// RecurrenceType enumerates the values for recurrence type.
type RecurrenceType string

const (
	// RecurrenceTypeAnnually ...
	RecurrenceTypeAnnually RecurrenceType = "Annually"
	// RecurrenceTypeDaily ...
	RecurrenceTypeDaily RecurrenceType = "Daily"
	// RecurrenceTypeMonthly ...
	RecurrenceTypeMonthly RecurrenceType = "Monthly"
	// RecurrenceTypeWeekly ...
	RecurrenceTypeWeekly RecurrenceType = "Weekly"
)

// PossibleRecurrenceTypeValues returns an array of possible values for the RecurrenceType const type.
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return []RecurrenceType{RecurrenceTypeAnnually, RecurrenceTypeDaily, RecurrenceTypeMonthly, RecurrenceTypeWeekly}
}

// SortDirection enumerates the values for sort direction.
type SortDirection string

const (
	// Ascending ...
	Ascending SortDirection = "Ascending"
	// Descending ...
	Descending SortDirection = "Descending"
)

// PossibleSortDirectionValues returns an array of possible values for the SortDirection const type.
func PossibleSortDirectionValues() []SortDirection {
	return []SortDirection{Ascending, Descending}
}

// StatusType enumerates the values for status type.
type StatusType string

const (
	// Active ...
	Active StatusType = "Active"
	// Inactive ...
	Inactive StatusType = "Inactive"
)

// PossibleStatusTypeValues returns an array of possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return []StatusType{Active, Inactive}
}

// TimeframeType enumerates the values for timeframe type.
type TimeframeType string

const (
	// BillingMonthToDate ...
	BillingMonthToDate TimeframeType = "BillingMonthToDate"
	// Custom ...
	Custom TimeframeType = "Custom"
	// MonthToDate ...
	MonthToDate TimeframeType = "MonthToDate"
	// TheLastBillingMonth ...
	TheLastBillingMonth TimeframeType = "TheLastBillingMonth"
	// TheLastMonth ...
	TheLastMonth TimeframeType = "TheLastMonth"
	// TheLastWeek ...
	TheLastWeek TimeframeType = "TheLastWeek"
	// TheLastYear ...
	TheLastYear TimeframeType = "TheLastYear"
	// WeekToDate ...
	WeekToDate TimeframeType = "WeekToDate"
	// YearToDate ...
	YearToDate TimeframeType = "YearToDate"
)

// PossibleTimeframeTypeValues returns an array of possible values for the TimeframeType const type.
func PossibleTimeframeTypeValues() []TimeframeType {
	return []TimeframeType{BillingMonthToDate, Custom, MonthToDate, TheLastBillingMonth, TheLastMonth, TheLastWeek, TheLastYear, WeekToDate, YearToDate}
}