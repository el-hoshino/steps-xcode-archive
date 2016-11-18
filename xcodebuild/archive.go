package xcodebuild

import (
	"os"

	"os/exec"

	"github.com/bitrise-io/go-utils/cmdex"
)

func (xb Model) archiveCmdSlice() []string {
	slice := []string{toolName}

	if xb.projectAction != "" {
		slice = append(slice, xb.projectAction)
	}
	if xb.projectPath != "" {
		slice = append(slice, xb.projectPath)
	}
	if xb.scheme != "" {
		slice = append(slice, "-scheme", xb.scheme)
	}
	if xb.configuration != "" {
		slice = append(slice, "-configuration", xb.configuration)
	}

	if xb.isCleanBuild {
		slice = append(slice, "clean")
	}

	slice = append(slice, "archive")

	if xb.archivePath != "" {
		slice = append(slice, "-archivePath", xb.archivePath)
	}

	if xb.forceDevelopmentTeam != "" {
		slice = append(slice, "DEVELOPMENT_TEAM", xb.forceDevelopmentTeam)
	}
	if xb.forceProvisioningProfileSpecifier != "" {
		slice = append(slice, "DEVELOPMENT_TEAM", xb.forceProvisioningProfileSpecifier)
	}
	if xb.forceProvisioningProfile != "" {
		slice = append(slice, "PROVISIONING_PROFILE", xb.forceProvisioningProfile)
	}
	if xb.forceCodeSignIdentity != "" {
		slice = append(slice, "CODE_SIGN_IDENTITY", xb.forceCodeSignIdentity)
	}

	slice = append(slice, xb.customOptions...)

	return slice
}

// PrintableArchiveCmd ...
func (xb Model) PrintableArchiveCmd() string {
	cmdSlice := xb.archiveCmdSlice()
	return cmdex.PrintableCommandArgs(false, cmdSlice)
}

// ArchiveCmd ...
func (xb Model) ArchiveCmd() (*exec.Cmd, error) {
	cmdSlice := xb.archiveCmdSlice()
	cmd, err := cmdex.NewCommandFromSlice(cmdSlice)
	if err != nil {
		return nil, err
	}
	return cmd.GetCmd(), nil
}

// Archive ...
func (xb Model) Archive() error {
	cmdSlice := xb.archiveCmdSlice()
	cmd, err := cmdex.NewCommandFromSlice(cmdSlice)
	if err != nil {
		return err
	}

	cmd.SetStdout(os.Stdout)
	cmd.SetStderr(os.Stderr)

	return cmd.Run()
}
