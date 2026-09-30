# Licenses for the embedded fonts

The four `.ttf` files in this directory are **MesloLGS Nerd Font Mono**. They are
redistributed inside the `termshot` binary through `//go:embed`, so their licenses
travel with every build. This file records them.

The rest of this repository is under termshot's own MIT license, in `LICENSE` at the
repository root. That license does not cover these font files.

## What the font is made of

MesloLGS Nerd Font is Meslo LG with the Nerd Fonts icon set patched in, so two licenses
apply at once:

| Layer | License | Holder |
|---|---|---|
| Meslo LG (base font) | Apache License 2.0 | Copyright 2009, 2010, 2013 André Berg |
| Nerd Fonts patch (icon glyphs) | SIL Open Font License 1.1 | Copyright (c) 2014 Ryan L McIntyre |

Meslo LG is itself derived from Apple's Menlo, which derives from Bitstream Vera. Those
credits are recorded in the font's own metadata: Copyright 2009 Apple Inc., Copyright
2006 Tavmjong Bah, Copyright 2003 Bitstream Inc.

## Meslo LG: Apache License 2.0

Meslo LG is distributed under the Apache License, Version 2.0. The full text is at
http://www.apache.org/licenses/LICENSE-2.0

The license requires this notice, quoted here word for word:

> Unless required by applicable law or agreed to in writing, software distributed under
> the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF
> ANY KIND, either express or implied. See the License for the specific language
> governing permissions and limitations under the License.

Upstream: https://github.com/andreberg/Meslo-Font

## Nerd Fonts patch: SIL Open Font License 1.1

The patched fonts Nerd Fonts produces are licensed under the SIL Open Font License,
Version 1.1. The full text is at https://openfontlicense.org and in the Nerd Fonts
repository at https://github.com/ryanoasis/nerd-fonts/blob/master/LICENSE

The OFL permits redistribution, with or without modification, under three conditions.
The fonts are not sold on their own. The copyright and license notice travel with them.
No derivative uses the Reserved Font Name without permission.

Upstream: https://github.com/ryanoasis/nerd-fonts

## Why these fonts are embedded

termshot renders text with `golang/freetype`, which performs no font fallback. A code
point the single bundled face lacks is drawn as a tofu box. The upstream face (Hack)
lacks the dingbat, braille, and several symbol ranges. Command-line tools use those
ranges for status markers, so those screens rendered as boxes. MesloLGS Nerd Font Mono covers
those ranges, and the "Mono" variant keeps every icon one cell wide so tables stay
aligned. See `NOTES.md` at the repository root.
