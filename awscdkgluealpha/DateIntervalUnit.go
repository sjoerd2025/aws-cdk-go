package awscdkgluealpha


// Date interval unit for partition projection.
//
// Example:
//   var myDatabase Database
//
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("data"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	PartitionKeys: []Column{
//   		&Column{
//   			Name: jsii.String("date"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   	PartitionProjection: map[string]PartitionProjectionConfiguration{
//   		"date": glue.PartitionProjectionConfiguration_date(&DatePartitionProjectionConfigurationProps{
//   			"min": jsii.String("2020-01-01"),
//   			"max": jsii.String("2023-12-31"),
//   			"format": jsii.String("yyyy-MM-dd"),
//   			// `step` bundles interval + unit (supply both or neither). Optional at day
//   			// precision or coarser; required when the format is sub-day (e.g. hours).
//   			"step": &DateProjectionStep{
//   				"interval": jsii.Number(1),
//   				"intervalUnit": glue.DateIntervalUnit_DAYS,
//   			},
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/ug/partition-projection-supported-types.html#partition-projection-date-type
//
// Deprecated.
type DateIntervalUnit string

const (
	// Year interval.
	// Deprecated.
	DateIntervalUnit_YEARS DateIntervalUnit = "YEARS"
	// Month interval.
	// Deprecated.
	DateIntervalUnit_MONTHS DateIntervalUnit = "MONTHS"
	// Week interval.
	// Deprecated.
	DateIntervalUnit_WEEKS DateIntervalUnit = "WEEKS"
	// Day interval (default).
	// Deprecated.
	DateIntervalUnit_DAYS DateIntervalUnit = "DAYS"
	// Hour interval.
	// Deprecated.
	DateIntervalUnit_HOURS DateIntervalUnit = "HOURS"
	// Minute interval.
	// Deprecated.
	DateIntervalUnit_MINUTES DateIntervalUnit = "MINUTES"
	// Second interval.
	// Deprecated.
	DateIntervalUnit_SECONDS DateIntervalUnit = "SECONDS"
)

