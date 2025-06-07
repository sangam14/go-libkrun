/*
 * This is an example implementing chroot-like functionality with libkrun.
 *
 * It executes the requested command (relative to NEWROOT) inside a fresh
 * Virtual Machine created and managed by libkrun.
 */

package main

import (
	"fmt"
	"os"

	"go-libkrun/pkg/krun"
)

var errno int32

func perror(message string) {
	fmt.Fprintf(os.Stderr, "%s: %d\n", message, errno)
}

func listenShutdownRequest(fd int) {
}

func bootEfi(args []string) int {
	diskImage := args[1]    // raw format

	// Set the log level to "off".
	e := krun.SetLogLevel(0)
	if e != 0 {
		errno = -e
		perror("Error configuring log level")
		return -1
	}

	// Create the configuration context.
	ctx := krun.CreateCtx()
	if ctx < 0 {
		errno = -ctx
		perror("Error creating configuration context")
		return -1
	}
	ctxId := uint32(ctx)

	// Configure the number of vCPUs (2) and the amount of RAM (1024 MiB).
	if e := krun.SetVmConfig(ctxId, 2, 1024); e != 0 {
		errno = -e
		perror("Error configuring the number of vCPUs and/or the amount of RAM")
		return -1
	}

	if e := krun.SetRootDisk(ctxId, diskImage); e != 0 {
		errno = -e
        perror("Error configuring disk image");
		return -1
	}

	efd := krun.GetShutdownEventfd(ctxId);
	if (efd < 0) {
		perror("Can't get shutdown eventfd");
		return -1;
	}

	// Spawn a thread to listen on "/tmp/krun_shutdown.sock" for a request to send
	// a shutdown signal to the guest.
	go listenShutdownRequest(efd);

	// Start and enter the microVM. Unless there is some error while creating the microVM
	// this function never returns.
	if e := krun.StartEnter(ctxId); e != 0 {
		errno = -e
		perror("Error creating the microVM")
		return -1
	}

	// Not reached.
	return 0
}

func main() {
	os.Exit(bootEfi(os.Args))
}
