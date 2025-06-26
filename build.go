// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2021 Authors of Cilium */

//go:generate sh -c "echo Generating for $TARGET_GOARCH"
//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target $TARGET_GOARCH -cc clang -no-strip KProbePWRU ./bpf/kprobe_pwru.c -- -I./bpf/headers -Wno-address-of-packed-member

// Manual build with asm output
// env BPF2GO_CFLAGS="-save-temps" ARCHS=arm64 CC=aarch64-linux-gnu-gcc DOLLAR=$'$' GOARCH=arm64 GOFILE=build.go GOLINE=5 GOOS=linux GOPACKAGE=main GOROOT=/usr/lib/go LIBPCAP_ARCH=aarch64-unknown-linux-gnu MAKEFLAGS=$' -- CC=aarch64-linux-gnu-gcc LIBPCAP_ARCH=aarch64-unknown-linux-gnu TARGET_GOARCH=arm64' MAKELEVEL=2 MAKE_TERMERR=/dev/pts/29 MAKE_TERMOUT=/dev/pts/29 MFLAGS='' TARGET_GOARCH=arm64 PATH=$'/usr/lib/go/bin:/usr/lib/go/bin:/home/kxxt/mambaforge/bin:/home/kxxt/mambaforge/bin:/home/kxxt/.cargo/bin:/home/kxxt/mambaforge/bin:/home/kxxt/.nix-profile/bin:/nix/var/nix/profiles/default/bin:/usr/local/sbin:/usr/local/bin:/usr/bin:/usr/lib/jvm/default/bin:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl:/usr/lib/rustup/bin' SHLVL=5 _=/usr/bin/go /home/kxxt/.cache/go-build/60/60edac95b7e037bdce49a27e176af31c6507bbcc81a45030dfcc8854c8488a04-d/bpf2go -target arm64 -cc clang -no-strip KProbePWRU ./bpf/kprobe_pwru.c -- -I./bpf/headers -Wno-address-of-packed-member

package main
