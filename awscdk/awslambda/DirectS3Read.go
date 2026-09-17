package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The DirectS3Read configuration for an S3 Files filesystem mount.
//
// Direct reads let Lambda read objects straight from the backing S3 bucket for
// higher throughput, instead of routing every read through the file system mount.
//
// Create one with a factory method:
//
// - `DirectS3Read.enabled(bucket)` — turn direct reads on and grant the execution
//   role read access to `bucket`.
// - `DirectS3Read.enabledWithoutGrant()` — turn direct reads on but add no S3
//   permissions; grant read access to the execution role yourself.
// - `DirectS3Read.auto()` — let the service decide based on the function's memory.
// - `DirectS3Read.disabled()` — always read through the mount.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   // Versioning is required — S3 Files relies on object versions for consistency.
//   bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
//   	Versioned: jsii.Boolean(true),
//   })
//
//   // S3 Files assumes this role to sync data between S3 and the file system.
//   role := iam.NewRole(this, jsii.String("S3FilesRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("elasticfilesystem.amazonaws.com")),
//   })
//
//   // S3 permissions: read/write access to the bucket and objects
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("s3:ListBucket*"),
//   	},
//   	Resources: []*string{
//   		bucket.bucketArn,
//   	},
//   }))
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("s3:AbortMultipartUpload"),
//   		jsii.String("s3:DeleteObject"),
//   		jsii.String("s3:GetObject*"),
//   		jsii.String("s3:List*"),
//   		jsii.String("s3:PutObject*"),
//   	},
//   	Resources: []*string{
//   		bucket.ArnForObjects(jsii.String("*")),
//   	},
//   }))
//
//   // EventBridge permissions: S3 Files creates rules prefixed "DO-NOT-DELETE-S3-Files"
//   // to detect S3 object changes and trigger data synchronization.
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("events:DeleteRule"),
//   		jsii.String("events:DisableRule"),
//   		jsii.String("events:EnableRule"),
//   		jsii.String("events:PutRule"),
//   		jsii.String("events:PutTargets"),
//   		jsii.String("events:RemoveTargets"),
//   	},
//   	Resources: []*string{
//   		fmt.Sprintf("arn:%v:events:*:*:rule/DO-NOT-DELETE-S3-Files*", cdk.Aws_PARTITION()),
//   	},
//   	Conditions: map[string]interface{}{
//   		"StringEquals": map[string]*string{
//   			"events:ManagedBy": jsii.String("elasticfilesystem.amazonaws.com"),
//   		},
//   	},
//   }))
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("events:DescribeRule"),
//   		jsii.String("events:ListRuleNamesByTarget"),
//   		jsii.String("events:ListRules"),
//   		jsii.String("events:ListTargetsByRule"),
//   	},
//   	Resources: []*string{
//   		fmt.Sprintf("arn:%v:events:*:*:rule/*", cdk.Aws_PARTITION()),
//   	},
//   }))
//
//   fileSystem := s3files.NewCfnFileSystem(this, jsii.String("S3FilesFs"), &CfnFileSystemProps{
//   	Bucket: bucket.bucketArn,
//   	RoleArn: role.roleArn,
//   })
//
//   sg := ec2.NewSecurityGroup(this, jsii.String("MountTargetSG"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   // Create a mount target in each private subnet so Lambda can reach the file system via NFS.
//   vpc.PrivateSubnets.forEach((subnet, i) =>
//     new s3files.CfnMountTarget(this, `MountTarget${i}`, {
//       fileSystemId: fileSystem.attrFileSystemId,
//       subnetId: subnet.subnetId,
//       securityGroups: [sg.securityGroupId],
//     }))
//
//   // The access point defines the POSIX identity and root path Lambda uses on the file system.
//   accessPoint := s3files.NewCfnAccessPoint(this, jsii.String("AccessPoint"), &CfnAccessPointProps{
//   	FileSystemId: fileSystem.attrFileSystemId,
//   	RootDirectory: &RootDirectoryProperty{
//   		Path: jsii.String("/export/lambda"),
//   		CreationPermissions: &CreationPermissionsProperty{
//   			OwnerGid: jsii.String("1001"),
//   			OwnerUid: jsii.String("1001"),
//   			Permissions: jsii.String("750"),
//   		},
//   	},
//   	PosixUser: &PosixUserProperty{
//   		Gid: jsii.String("1001"),
//   		Uid: jsii.String("1001"),
//   	},
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	Vpc: Vpc,
//   	Filesystem: lambda.FileSystem_FromS3FilesAccessPoint(accessPoint, jsii.String("/mnt/s3files"), &S3FilesOptions{
//   		// Enables direct reads and grants s3:GetObject/s3:GetObjectVersion on the bucket to the execution role.
//   		DirectS3Read: lambda.DirectS3Read_Enabled(bucket),
//   	}),
//   })
//
type DirectS3Read interface {
}

// The jsii proxy struct for DirectS3Read
type jsiiProxy_DirectS3Read struct {
	_ byte // padding
}

// Let the service decide whether to use direct S3 read based on the function's memory configuration: direct reads are active for functions with 512 MB or more of memory.
//
// No S3 read permissions are added; the execution role must already hold them for a
// service-initiated direct read to succeed, otherwise reads fall back to the mount.
func DirectS3Read_Auto() DirectS3Read {
	_init_.Initialize()

	var returns DirectS3Read

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DirectS3Read",
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Disable direct S3 read;
//
// all reads are routed through the S3 Files file system's
// high-performance storage.
func DirectS3Read_Disabled() DirectS3Read {
	_init_.Initialize()

	var returns DirectS3Read

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DirectS3Read",
		"disabled",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Enable direct S3 reads, bypassing the mount for higher throughput, and grant the function's execution role `s3:GetObject` and `s3:GetObjectVersion` on the bucket's objects so that direct reads can succeed.
//
// Unlike `auto()`, this enables direct reads regardless of the function's memory size,
// including functions with less than 512 MB of memory.
//
// If the bucket is encrypted with a customer-managed KMS key, also grant the execution
// role `kms:Decrypt` on that key yourself.
func DirectS3Read_Enabled(bucket awss3.IBucket) DirectS3Read {
	_init_.Initialize()

	if err := validateDirectS3Read_EnabledParameters(bucket); err != nil {
		panic(err)
	}
	var returns DirectS3Read

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DirectS3Read",
		"enabled",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}

// Enable direct S3 reads, bypassing the mount for higher throughput, without adding any S3 read permissions.
//
// Like `enabled()`, this enables direct reads regardless of the function's memory size,
// including functions with less than 512 MB of memory.
//
// Use this when the execution role already has `s3:GetObject`/`s3:GetObjectVersion` on the
// backing bucket (for example through a managed policy or a bucket policy). You are
// responsible for granting those permissions; without them, direct reads silently fall
// back to reading through the file system.
func DirectS3Read_EnabledWithoutGrant() DirectS3Read {
	_init_.Initialize()

	var returns DirectS3Read

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DirectS3Read",
		"enabledWithoutGrant",
		nil, // no parameters
		&returns,
	)

	return returns
}

