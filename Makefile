ARCH = $(shell uname -m)
OS = $(shell uname -s)
ROOTFS_DISTRO := fedora
ROOTFS_DIR = rootfs_$(ROOTFS_DISTRO)

.PHONY: clean rootfs

EXAMPLES := chroot_vm
ifeq ($(EFI),1)
    EXAMPLES := boot_efi
endif

all: $(EXAMPLES)

chroot_vm:
	go build -o $@ ./cmd/chroot_vm
ifeq ($(OS),Darwin)
	codesign --entitlements chroot_vm.entitlements --force -s - $@
endif

boot_efi:
	go build -tags efi -o $@ ./cmd/boot_efi
ifeq ($(OS),Darwin)
	codesign --entitlements chroot_vm.entitlements --force -s - $@
endif

# Build the rootfs to be used with chroot_vm.
rootfs:
	mkdir -p $(ROOTFS_DIR)
	podman create --name libkrun_chroot_vm $(ROOTFS_DISTRO)
	podman export libkrun_chroot_vm | tar xpf - -C $(ROOTFS_DIR)
	podman rm libkrun_chroot_vm

clean:
	rm -rf chroot_vm $(ROOTFS_DIR)
