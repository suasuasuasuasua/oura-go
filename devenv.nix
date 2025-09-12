{ pkgs, ... }:
{
  # https://devenv.sh/packages/
  packages = [ pkgs.git ];

  # https://devenv.sh/languages/
  languages.go = {
    enable = true;
    package = pkgs.go.overrideAttrs (prevAttrs: rec {
      pname = "go";
      version = "1.25.0";

      src = pkgs.fetchurl {
        url = "https://go.dev/dl/go${version}.src.tar.gz";
        hash = "sha256-S9AekSlyB7+kUOpA1NWpOxtTGl5DhHOyoG4Y4HciciU=";
      };

      patches = [
        # NOTE: this was the new patch that we need to go from 1.24.5 to 1.25.0
        # the old patch was iana-etc-1.17, whatever that is
        # https://github.com/NixOS/nixpkgs/blob/ab0f3607a6c7486ea22229b92ed2d355f1482ee0/pkgs/development/compilers/go/iana-etc-1.17.patch
        #
        # all of the other patches are the same and just copied from web links
        (pkgs.replaceVars
          (pkgs.fetchurl {
            url = "https://raw.githubusercontent.com/NixOS/nixpkgs/ab0f3607a6c7486ea22229b92ed2d355f1482ee0/pkgs/development/compilers/go/iana-etc-1.25.patch";
            hash = "sha256-cybgmNkGgxEENpbryS9EWzgIPMjD6svRpVXC/iM0J0A=";
          })
          {
            iana = pkgs.iana-etc;
          }
        )
        # Patch the mimetype database location which is missing on NixOS.
        # but also allow static binaries built with NixOS to run outside nix
        (pkgs.replaceVars
          (pkgs.fetchurl {
            url = "https://raw.githubusercontent.com/NixOS/nixpkgs/ab0f3607a6c7486ea22229b92ed2d355f1482ee0/pkgs/development/compilers/go/mailcap-1.17.patch";
            hash = "sha256-NVx/TYPtlLykcgNaJXu5dNyKZxiHNV5d9DNpsf8Gmdc=";
          })
          {
            inherit (pkgs) mailcap;
          }
        )
        # prepend the nix path to the zoneinfo files but also leave the original value for static binaries
        # that run outside a nix server
        (pkgs.replaceVars
          (pkgs.fetchurl {
            url = "https://raw.githubusercontent.com/NixOS/nixpkgs/ab0f3607a6c7486ea22229b92ed2d355f1482ee0/pkgs/development/compilers/go/tzdata-1.19.patch";
            hash = "sha256-cfil5FBE9yL+e3APTIRHt8/Dn05jYO4S1jnS0KYLO9Y=";
          })
          {
            inherit (pkgs) tzdata;
          }
        )
        (pkgs.fetchurl {
          url = "https://raw.githubusercontent.com/NixOS/nixpkgs/ab0f3607a6c7486ea22229b92ed2d355f1482ee0/pkgs/development/compilers/go/remove-tools-1.11.patch";
          hash = "sha256-KEKm9Hv+gIJoq1c7FWC7f9IIYcwJhaXHQOdhMlXe6A4=";
        })
        (pkgs.fetchurl {
          url = "https://raw.githubusercontent.com/NixOS/nixpkgs/ab0f3607a6c7486ea22229b92ed2d355f1482ee0/pkgs/development/compilers/go/go_no_vendor_checks-1.23.patch";
          hash = "sha256-uFoDPTCXEJ5sG4A0fpFwNJHnLdeoasYIHMGWyugmZZc=";
        })
      ];
    });
  };

  # https://devenv.sh/pre-commit-hooks/
  git-hooks.hooks = {
    # https://github.com/cachix/git-hooks.nix?tab=readme-ov-file#golang
    gofmt.enable = true;
    golangci-lint.enable = true;
    golines.enable = true;
    govet.enable = true;
    staticcheck.enable = true;
    # https://github.com/cachix/git-hooks.nix?tab=readme-ov-file#nix-1
    nixfmt-rfc-style.enable = true;
    prettier.enable = true;
  };

  cachix.enable = true;
  cachix.pull = [ "pre-commit-hooks" ];

  # See full reference at https://devenv.sh/reference/options/
}
