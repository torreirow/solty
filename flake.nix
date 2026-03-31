{
  description = "Solty - Solidtime CLI time tracking tool";

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
          solty = pkgs.buildGoModule {
            pname = "solty";
            version = pkgs.lib.strings.trim version;
            src = ./.;

            vendorHash = "sha256-yGF6iVMU1e+rjwo77Q/2LGX+5pnWuv5AXOQ9rHG1TxQ=";

            ldflags = [
              "-s"
              "-w"
              "-X main.version=${pkgs.lib.strings.trim version}"
            ];

            meta = with pkgs.lib; {
              description = "Command-line time tracking tool for Solidtime";
              homepage = "https://github.com/torreirow/solty";
              license = licenses.unfree;
              maintainers = [ ];
            };
          };
        });

      defaultPackage = forAllSystems (system: self.packages.${system}.solty);

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
              echo "Solty development environment"
              echo "Go version: $(go version)"
              echo ""
              echo "Available commands:"
              echo "  go build -o solty     - Build the binary"
              echo "  go test ./...         - Run tests"
              echo "  go run main.go        - Run without building"
            '';
          };
        });
    };
}
