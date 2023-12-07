# nix-build
# Autogenerated from man page /home/runner/.nix-profile/share/man/man1/nix-build.1.gz
complete -c nix-build -l no-out-link -d 'Do not create a symlink to the output path'
complete -c nix-build -l dry-run -d 'Show what store paths would be built or downloaded'
complete -c nix-build -l out-link -s o -d 'Change the name of the symlink to the output path created from result to outl…'
complete -c nix-build -l help -d 'Prints out a summary of the command syntax and exits'
complete -c nix-build -l version -d 'Prints out the Nix version number on standard output and exits'
complete -c nix-build -l quiet -d 'Decreases the level of verbosity of diagnostic messages printed on standard e…'
complete -c nix-build -l no-build-output -s Q -d 'By default, output written by builders to standard output and standard error …'
complete -c nix-build -l max-jobs -s j -d 'Sets the maximum number of build jobs that Nix will perform in parallel to th…'
complete -c nix-build -l cores -d 'Sets the value of the NIX_BUILD_CORES environment variable in the invocation …'
complete -c nix-build -l max-silent-time -d 'Sets the maximum number of seconds that a builder can go without producing an…'
complete -c nix-build -l timeout -d 'Sets the maximum number of seconds that a builder can run'
complete -c nix-build -l keep-going -s k -d 'Keep going in case of failed builds, to the greatest extent possible'
complete -c nix-build -l keep-failed -s K -d 'Specifies that in case of a build failure, the temporary directory (usually i…'
complete -c nix-build -l fallback -d 'Whenever Nix attempts to build a derivation for which substitutes are known f…'
complete -c nix-build -l no-build-hook -d 'Disables the build hook mechanism'
complete -c nix-build -l readonly-mode -d 'When this option is used, no attempt is made to open the Nix database'
complete -c nix-build -l arg -d 'This option is accepted by nix-env, nix-instantiate and nix-build'
complete -c nix-build -l argstr -d 'This option is like --arg, only the value is not a Nix expression but a string'
complete -c nix-build -l attr -s A -d 'Select an attribute from the top-level Nix expression being evaluated'
complete -c nix-build -l expr -s E -d 'Interpret the command line arguments as a list of Nix expressions to be parse…'
complete -c nix-build -s I -d 'Add a path to the Nix expression search path'
complete -c nix-build -l option -d 'Set the Nix configuration option name to value'
complete -c nix-build -l repair -d 'Fix corrupted or missing store paths by redownloading or rebuilding them'
complete -c nix-build -l verbose -s v
complete -c nix-build -o jN -d 'flag to GNU Make'

