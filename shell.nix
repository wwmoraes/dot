{
  pkgs,
  ...
}:
rec {
  default = pkgs.mkShell {
    nativeBuildInputs = [
      # keep-sorted start
      pkgs.graphviz
      pkgs.remake
      pkgs.unstable.go
      (pkgs.python3.withPackages (
        pyPkgs: with pyPkgs; [
          mkdocs
          mkdocs-material
        ]
      ))
      # keep-sorted end
    ];
  };

  ci = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = [
        # keep-sorted start
        # keep-sorted end
      ]
      ++ prev.nativeBuildInputs;
    }
  );

  terminal = default.overrideAttrs (
    final: prev: {
      nativeBuildInputs = [
        # keep-sorted start
        pkgs.nix-update
        pkgs.unstable.golangci-lint
        # keep-sorted end
      ]
      ++ prev.nativeBuildInputs;
    }
  );
}
