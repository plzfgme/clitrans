{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, flake-utils, nixpkgs, gomod2nix }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        overlays = [ gomod2nix.overlays.default ];

        pkgs = (import nixpkgs) {
          inherit system overlays;
        };
      in
      {
        packages = {
          default = pkgs.buildGoApplication {
            pname = "clitrans";
            version = "0.1.0";
            src = ./.;
            modules = ./gomod2nix.toml;
          };
        };

        devShells = {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              gopls
              gotools
              golangci-lint
              pkgs.gomod2nix
            ];
          };
        };
      }
    );
}
