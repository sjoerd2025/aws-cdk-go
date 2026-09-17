package awscdkgluealpha


// The set of pre-installed Python libraries available to a Python shell job running Python 3.9.
// See: https://docs.aws.amazon.com/glue/latest/dg/add-job-python.html#python-shell-supported-library
//
// Deprecated.
type LibrarySet string

const (
	// Include the common analytics libraries for Python 3.9 (e.g. pandas, numpy, scikit-learn, awswrangler).
	// Deprecated.
	LibrarySet_ANALYTICS LibrarySet = "ANALYTICS"
	// Do not install the common library set.
	//
	// Use this when your libraries are custom or conflict
	// with the pre-installed ones.
	// Deprecated.
	LibrarySet_NONE LibrarySet = "NONE"
)

