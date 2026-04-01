{
  description = "Soltty - Solidtime CLI time tracking tool";

  inputs.nixpkgs.url = "nixpkgs/nixos-24.05";

  outputs = { self, nixpkgs }:
    let
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
          version = builtins.readFile ./VERSION;
        in
        {
          soltty = pkgs.buildGoModule {
            pname = "soltty";
            version = pkgs.lib.strings.trim version;
            src = ./.;

            # vendorHash is automatically updated by release.sh when Go dependencies change.
            # To manually update: Run `nix build .#soltty 2>&1 | grep "got:"` and use the hash shown.
            vendorHash = "sha256-Wr+GuQ7AbD0RmHppiLU+6GxRl8FRqp7yht8mnZak/2A=";

            ldflags = [
              "-s"
              "-w"
              "-X github.com/torreirow/soltty/cmd.version=${pkgs.lib.strings.trim version}"
            ];

            meta = with pkgs.lib; {
              description = "Command-line time tracking tool for Solidtime";
              homepage = "https://github.com/torreirow/soltty";
              license = licenses.mit;
              maintainers = [ ];
            };
          };
        });

      defaultPackage = forAllSystems (system: self.packages.${system}.soltty);

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
              gotools
              go-tools
            ];

            shellHook = ''
              echo "Soltty development environment"
              echo "Go version: $(go version)"
              echo ""
              echo "Available commands:"
              echo "  go build -o soltty    - Build the binary"
              echo "  go test ./...         - Run tests"
              echo "  go run main.go        - Run without building"
            '';
          };
        });
    };
}
