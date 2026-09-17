//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateDirectS3Read_EnabledParameters(bucket awss3.IBucket) error {
	return nil
}

