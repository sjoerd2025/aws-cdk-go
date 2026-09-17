//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

func validateDirectS3Read_EnabledParameters(bucket awss3.IBucket) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}

