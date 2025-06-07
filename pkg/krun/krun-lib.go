//go:build !sev && !efi
package krun

/*
#include <libkrun.h>
#include <stdlib.h>
#cgo pkg-config: libkrun
*/
import "C"
