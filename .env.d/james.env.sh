#!/bin/bash

VERSION="0.2.1"

activate-venv() {
    # To create venv: `python3 -m venv <dir>`
    local dir
    for dir in venv .venv; do
	[ -d "$dir" ] || continue
        [ -f "$dir/bin/activate" ] || (python3 -m venv "$dir" && "$dir/bin/python" -m pip install packaging)
	source "$dir/bin/activate"
	echo "Activated $dir"
	unset PS1
	command -v ensure-py-deps > /dev/null && ensure-py-deps
	return
    done

    # Fall back to a shared general venv
    [ -f ~/Python/venv/general/bin/activate ] && source ~/Python/venv/general/bin/activate
}

rename-tmux-window() {
    if [ -n "$TMUX" ]; then
        tmux rename-window "$(basename "$PWD")"
    fi
}

set-go-cache() {
    repo_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)
    if [ -d "$repo_root/.go-cache" ]; then
	export GOCACHE="$repo_root/.go-cache"
    fi
}

set-patchwise-context-dir() {
    local context_dir
    if [ -n "${TMUX_RECORD_DIR:-}" ]; then
        context_dir="$TMUX_RECORD_DIR"
    else
        local dir_specifier="${PWD#/}"
        dir_specifier="${dir_specifier//\//-}"
        context_dir="$HOME/var/tmux-record/$dir_specifier"
    fi
    export PATCHWISE_CONTEXT_DIR="$context_dir"
}

sync-envrc() {
    local shared_envrc="$HOME/src/envrc/.envrc"
    local repo_envrc="$PWD/.envrc"
    local shared_envd="$HOME/src/envrc/.env.d"
    local repo_envd="$PWD/.env.d"

    if [ -f "$shared_envrc" ] && [ ! -L "$shared_envrc" ] \
        && [ -f "$repo_envrc" ] && [ ! -L "$repo_envrc" ] \
        && [ "$shared_envrc" -nt "$repo_envrc" ]; then
        cp -p -- "$shared_envrc" "$repo_envrc"
    fi

    if [ -d "$shared_envd" ] && [ ! -L "$shared_envd" ] \
        && [ -d "$repo_envd" ] && [ ! -L "$repo_envd" ]; then
        local shared_script repo_script script_name
        for shared_script in "$shared_envd"/*; do
            [ -f "$shared_script" ] || continue
            [ -L "$shared_script" ] && continue
            script_name=$(basename "$shared_script")
            repo_script="$repo_envd/$script_name"
            if [ ! -e "$repo_script" ] || [ "$shared_script" -nt "$repo_script" ]; then
                cp -p -- "$shared_script" "$repo_script"
            fi
        done
    fi
}

activate-venv
rename-tmux-window
set-go-cache
set-patchwise-context-dir
sync-envrc

# Local Variables:
# mode: shell-script
# End:
# -----BEGIN PGP SIGNATURE-----
# 
# iHUEABYKAB0WIQTE1CImRUv0z3tyc0sZxqbCbxAX5wUCaohSiAAKCRAZxqbCbxAX
# 5wMCAP4jzueB3hefYHZ1rpch5cTjEqPqMY3gBNblaEtJzJwwQwD/UoItd7pdNb90
# ovMl9fFPbtI9gQHZChm0tT8RoKRLiQU=
# =CqHR
# -----END PGP SIGNATURE-----
