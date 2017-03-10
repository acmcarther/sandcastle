with (import <nixpkgs> {});
let
  pkgs = import <nixpkgs> {};
in pkgs.stdenv.mkDerivation rec {
  name = "raze";
  propagatedBuildInputs = [ pkgs.pkgconfig pkgs.openssl pkgs.curl pkgs.scala ];
  shellHook = ''
    # Allow my shell to add custom snippet
    export IS_NIX_SHELL=1
    export GOPATH="$(pwd)"
    export BAZEL_SH=/run/current-system/sw/bin/bash
  '';
}
