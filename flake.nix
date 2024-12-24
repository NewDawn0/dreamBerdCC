{
  description = "A dreamBerd compiler";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    nix-systems.url = "github:nix-systems/default";
  };

  outputs = { self, nixpkgs, ... }@inputs:
    let eachSystem = nixpkgs.lib.genAttrs (import inputs.nix-systems);
    in {
      overlays.default =
        (final: prev: { note = self.packages.${prev.system}.default; });
      packages = eachSystem (system:
        let pkgs = nixpkgs.legacyPackages.${system};
        in {
          default = pkgs.buildGoModule {
            pname = "dreamcc";
            version = "1.0.0";
            src = ./.;
            vendorHash = null;
            meta = {
              description = "A dreamBerd compiler";
              longDescription = ''
                Compiles dreamBerd code to a binary
              '';
              homepage = "https://github.com/NewDawn0/dreamBerdCC";
              license = pkgs.lib.licenses.mit;
              maintainers = with pkgs.lib.maintainers; [ NewDawn0 ];
              platforms = pkgs.lib.platforms.all;
            };
          };
        });
    };
}
