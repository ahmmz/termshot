# termshot-meslo

Fork of [homeport/termshot](https://github.com/homeport/termshot) that renders
with MesloLGS Nerd Font Mono instead of Hack. Hack has no glyph for several
dingbats and symbol characters. There is no font flag and no fallback, so those
characters came out as tofu boxes.

Two new files, `internal/img/font.go` and `internal/img/fonts/`, embed the four
Meslo font faces behind the old Hack interface. `internal/img/output.go` uses
them.

The files in `test/data/*.png` are golden images compared byte-for-byte. This
fork regenerated them for Meslo, so they conflict on a merge from upstream. To
redo them, run a throwaway `main` inside `internal/img/`. That program rebuilds
the scaffolds from `output_test.go`.
