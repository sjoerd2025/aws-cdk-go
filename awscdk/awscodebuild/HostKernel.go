package awscodebuild


// The host operating system kernel used for builds in a CodeBuild project.
//
// The host kernel does not affect the build environment operating system,
// which is determined by the build image.
//
// Only applies to the `LINUX_CONTAINER`, `ARM_CONTAINER`, `LINUX_EC2` and `ARM_EC2`
// environment types. It is not applicable to Windows, Lambda or Mac environment types.
//
// Example:
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
//   		HostKernel: codebuild.HostKernel_LINUX_KERNEL_6,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/codebuild/latest/APIReference/API_ProjectEnvironment.html#CodeBuild-Type-ProjectEnvironment-hostKernel
//
type HostKernel string

const (
	// Runs on an Amazon Linux 2 host (kernel 4.x).
	HostKernel_LINUX_KERNEL_4 HostKernel = "LINUX_KERNEL_4"
	// Runs on an Amazon Linux 2023 host (kernel 6.x).
	HostKernel_LINUX_KERNEL_6 HostKernel = "LINUX_KERNEL_6"
	// Runs on the latest supported host kernel.
	HostKernel_LINUX_KERNEL_LATEST HostKernel = "LINUX_KERNEL_LATEST"
)

