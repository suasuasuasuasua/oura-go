{ pkgs, lib, config, inputs, ... }:

{
  # https://devenv.sh/packages/
  packages = [ pkgs.git ];

  # https://devenv.sh/languages/
  languages.go.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  pre-commit.hooks = {
    # https://github.com/cachix/git-hooks.nix?tab=readme-ov-file#golang
    gofmt.enable = true;
    golangci-lint.enable = true;
    golines.enable = true;
    govet.enable = true;
    staticcheck.enable = true;
    # https://github.com/cachix/git-hooks.nix?tab=readme-ov-file#nix-1
    nixfmt-rfc-style.enable = true;
  };

  # See full reference at https://devenv.sh/reference/options/
}