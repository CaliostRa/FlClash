#!/usr/bin/env bash
set -euo pipefail

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
core_dir="${root_dir}/core/Clash.Meta"
patch_file="${root_dir}/tool/patches/clash-meta-reality-firefox148.patch"

if git -C "${core_dir}" apply --unidiff-zero --reverse --check "${patch_file}" >/dev/null 2>&1; then
  echo "Core patch already applied"
  exit 0
fi

git -C "${core_dir}" apply --unidiff-zero --check "${patch_file}"
git -C "${core_dir}" apply --unidiff-zero "${patch_file}"
echo "Applied Firefox 148 REALITY core patch"
