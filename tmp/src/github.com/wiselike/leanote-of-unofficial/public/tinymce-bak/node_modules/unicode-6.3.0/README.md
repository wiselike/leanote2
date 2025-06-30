# Unicode v6.3.0 data

JavaScript-compatible Unicode data for use in Node.js. Included: arrays of code points, arrays of symbols, and regular expressions for Unicode v6.3.0’s categories, scripts, script extensions, blocks, and properties, as well as bidi mirroring and case folding data.

The data files in this module are generated as part of the [node-unicode-data](https://mths.be/node-unicode-data) project. **Please report any bugs or requests [in the appropriate issue tracker](https://github.com/mathiasbynens/node-unicode-data/issues).**

## Installation

```bash
npm install unicode-6.3.0 --save-dev
```

**Note:** _unicode-6.3.0_ is supposed to be used in build scripts (i.e. as a `devDependency`), and not at runtime (i.e. as a regular `dependency`).

## Regular expressions

The Unicode data modules ship with pre-compiled regular expressions for categories, scripts, script extensions, blocks, and properties. But maybe you want to create a single regular expression that combines several categories, scripts, etc. In that case, [***you should use Regenerate***](https://mths.be/regenerate). For example, to construct a regex that matches all symbols in the Arabic and Greek scripts as per Unicode v6.3.0:

```js
const regenerate = require('regenerate');
const set = regenerate()
  .add(require('unicode-6.3.0/scripts/Arabic/code-points')) // or `…/symbols`, doesn’t matter
  .add(require('unicode-6.3.0/scripts/Greek/code-points')); // or `…/symbols`, doesn’t matter
console.log(set.toString());
// Then you might want to use a template like this to write the result to a file, along with any regex flags you might need:
// const regex = /<%= set.toString() %>/gim;
```

## Usage

```js
// Get an array of code points in a given Unicode category:
const codePoints = require('unicode-6.3.0/categories/Lu/code-points');
// Get an array of symbols (strings) in a given Unicode category:
const symbols = require('unicode-6.3.0/categories/Lu/symbols');
// Get a regular expression that matches any symbol in a given Unicode category:
const regex = require('unicode-6.3.0/categories/Lu/regex');
// Get the canonical category a given code point belongs to:
// (Note: U+0041 is LATIN CAPITAL LETTER A)
const category = require('unicode-6.3.0/categories')[ 0x41 ];
// Get an array of all code points with the `Bidi_ON` bidi property:
const on = require('unicode-6.3.0/bidi/ON/code-points');
// Get the directionality of a given code point:
const directionality = require('unicode-6.3.0/bidi')[ 0x41 ];

// What glyph is the mirror image of `«` (U+00AB)?
const mirrored = require('unicode-6.3.0/bidi-mirroring')[ 0xAB ];

// Get a regular expression that matches all opening brackets:
const openingBrackets = require('unicode-6.3.0/bidi-brackets/Open/regex');

// …you get the idea.
```

Other than categories, data on Unicode properties, blocks, scripts, and script extensions is available too (for recent versions of the Unicode standard). Here’s the full list of the available data for v6.3.0:

```js
// properties:

require('unicode-6.3.0/properties/Any/code-points');
require('unicode-6.3.0/properties/Any/symbols');
require('unicode-6.3.0/properties/Any/regex');

require('unicode-6.3.0/properties/Assigned/code-points');
require('unicode-6.3.0/properties/Assigned/symbols');
require('unicode-6.3.0/properties/Assigned/regex');

require('unicode-6.3.0/properties/ASCII/code-points');
require('unicode-6.3.0/properties/ASCII/symbols');
require('unicode-6.3.0/properties/ASCII/regex');

require('unicode-6.3.0/properties/White_Space/code-points');
require('unicode-6.3.0/properties/White_Space/symbols');
require('unicode-6.3.0/properties/White_Space/regex');

require('unicode-6.3.0/properties/Bidi_Control/code-points');
require('unicode-6.3.0/properties/Bidi_Control/symbols');
require('unicode-6.3.0/properties/Bidi_Control/regex');

require('unicode-6.3.0/properties/Join_Control/code-points');
require('unicode-6.3.0/properties/Join_Control/symbols');
require('unicode-6.3.0/properties/Join_Control/regex');

require('unicode-6.3.0/properties/Dash/code-points');
require('unicode-6.3.0/properties/Dash/symbols');
require('unicode-6.3.0/properties/Dash/regex');

require('unicode-6.3.0/properties/Hyphen/code-points');
require('unicode-6.3.0/properties/Hyphen/symbols');
require('unicode-6.3.0/properties/Hyphen/regex');

require('unicode-6.3.0/properties/Quotation_Mark/code-points');
require('unicode-6.3.0/properties/Quotation_Mark/symbols');
require('unicode-6.3.0/properties/Quotation_Mark/regex');

require('unicode-6.3.0/properties/Terminal_Punctuation/code-points');
require('unicode-6.3.0/properties/Terminal_Punctuation/symbols');
require('unicode-6.3.0/properties/Terminal_Punctuation/regex');

require('unicode-6.3.0/properties/Other_Math/code-points');
require('unicode-6.3.0/properties/Other_Math/symbols');
require('unicode-6.3.0/properties/Other_Math/regex');

require('unicode-6.3.0/properties/Hex_Digit/code-points');
require('unicode-6.3.0/properties/Hex_Digit/symbols');
require('unicode-6.3.0/properties/Hex_Digit/regex');

require('unicode-6.3.0/properties/ASCII_Hex_Digit/code-points');
require('unicode-6.3.0/properties/ASCII_Hex_Digit/symbols');
require('unicode-6.3.0/properties/ASCII_Hex_Digit/regex');

require('unicode-6.3.0/properties/Other_Alphabetic/code-points');
require('unicode-6.3.0/properties/Other_Alphabetic/symbols');
require('unicode-6.3.0/properties/Other_Alphabetic/regex');

require('unicode-6.3.0/properties/Ideographic/code-points');
require('unicode-6.3.0/properties/Ideographic/symbols');
require('unicode-6.3.0/properties/Ideographic/regex');

require('unicode-6.3.0/properties/Diacritic/code-points');
require('unicode-6.3.0/properties/Diacritic/symbols');
require('unicode-6.3.0/properties/Diacritic/regex');

require('unicode-6.3.0/properties/Extender/code-points');
require('unicode-6.3.0/properties/Extender/symbols');
require('unicode-6.3.0/properties/Extender/regex');

require('unicode-6.3.0/properties/Other_Lowercase/code-points');
require('unicode-6.3.0/properties/Other_Lowercase/symbols');
require('unicode-6.3.0/properties/Other_Lowercase/regex');

require('unicode-6.3.0/properties/Other_Uppercase/code-points');
require('unicode-6.3.0/properties/Other_Uppercase/symbols');
require('unicode-6.3.0/properties/Other_Uppercase/regex');

require('unicode-6.3.0/properties/Noncharacter_Code_Point/code-points');
require('unicode-6.3.0/properties/Noncharacter_Code_Point/symbols');
require('unicode-6.3.0/properties/Noncharacter_Code_Point/regex');

require('unicode-6.3.0/properties/Other_Grapheme_Extend/code-points');
require('unicode-6.3.0/properties/Other_Grapheme_Extend/symbols');
require('unicode-6.3.0/properties/Other_Grapheme_Extend/regex');

require('unicode-6.3.0/properties/IDS_Binary_Operator/code-points');
require('unicode-6.3.0/properties/IDS_Binary_Operator/symbols');
require('unicode-6.3.0/properties/IDS_Binary_Operator/regex');

require('unicode-6.3.0/properties/IDS_Trinary_Operator/code-points');
require('unicode-6.3.0/properties/IDS_Trinary_Operator/symbols');
require('unicode-6.3.0/properties/IDS_Trinary_Operator/regex');

require('unicode-6.3.0/properties/Radical/code-points');
require('unicode-6.3.0/properties/Radical/symbols');
require('unicode-6.3.0/properties/Radical/regex');

require('unicode-6.3.0/properties/Unified_Ideograph/code-points');
require('unicode-6.3.0/properties/Unified_Ideograph/symbols');
require('unicode-6.3.0/properties/Unified_Ideograph/regex');

require('unicode-6.3.0/properties/Other_Default_Ignorable_Code_Point/code-points');
require('unicode-6.3.0/properties/Other_Default_Ignorable_Code_Point/symbols');
require('unicode-6.3.0/properties/Other_Default_Ignorable_Code_Point/regex');

require('unicode-6.3.0/properties/Deprecated/code-points');
require('unicode-6.3.0/properties/Deprecated/symbols');
require('unicode-6.3.0/properties/Deprecated/regex');

require('unicode-6.3.0/properties/Soft_Dotted/code-points');
require('unicode-6.3.0/properties/Soft_Dotted/symbols');
require('unicode-6.3.0/properties/Soft_Dotted/regex');

require('unicode-6.3.0/properties/Logical_Order_Exception/code-points');
require('unicode-6.3.0/properties/Logical_Order_Exception/symbols');
require('unicode-6.3.0/properties/Logical_Order_Exception/regex');

require('unicode-6.3.0/properties/Other_ID_Start/code-points');
require('unicode-6.3.0/properties/Other_ID_Start/symbols');
require('unicode-6.3.0/properties/Other_ID_Start/regex');

require('unicode-6.3.0/properties/Other_ID_Continue/code-points');
require('unicode-6.3.0/properties/Other_ID_Continue/symbols');
require('unicode-6.3.0/properties/Other_ID_Continue/regex');

require('unicode-6.3.0/properties/STerm/code-points');
require('unicode-6.3.0/properties/STerm/symbols');
require('unicode-6.3.0/properties/STerm/regex');

require('unicode-6.3.0/properties/Variation_Selector/code-points');
require('unicode-6.3.0/properties/Variation_Selector/symbols');
require('unicode-6.3.0/properties/Variation_Selector/regex');

require('unicode-6.3.0/properties/Pattern_White_Space/code-points');
require('unicode-6.3.0/properties/Pattern_White_Space/symbols');
require('unicode-6.3.0/properties/Pattern_White_Space/regex');

require('unicode-6.3.0/properties/Pattern_Syntax/code-points');
require('unicode-6.3.0/properties/Pattern_Syntax/symbols');
require('unicode-6.3.0/properties/Pattern_Syntax/regex');

require('unicode-6.3.0/properties/Math/code-points');
require('unicode-6.3.0/properties/Math/symbols');
require('unicode-6.3.0/properties/Math/regex');

require('unicode-6.3.0/properties/Alphabetic/code-points');
require('unicode-6.3.0/properties/Alphabetic/symbols');
require('unicode-6.3.0/properties/Alphabetic/regex');

require('unicode-6.3.0/properties/Lowercase/code-points');
require('unicode-6.3.0/properties/Lowercase/symbols');
require('unicode-6.3.0/properties/Lowercase/regex');

require('unicode-6.3.0/properties/Uppercase/code-points');
require('unicode-6.3.0/properties/Uppercase/symbols');
require('unicode-6.3.0/properties/Uppercase/regex');

require('unicode-6.3.0/properties/Cased/code-points');
require('unicode-6.3.0/properties/Cased/symbols');
require('unicode-6.3.0/properties/Cased/regex');

require('unicode-6.3.0/properties/Case_Ignorable/code-points');
require('unicode-6.3.0/properties/Case_Ignorable/symbols');
require('unicode-6.3.0/properties/Case_Ignorable/regex');

require('unicode-6.3.0/properties/Changes_When_Lowercased/code-points');
require('unicode-6.3.0/properties/Changes_When_Lowercased/symbols');
require('unicode-6.3.0/properties/Changes_When_Lowercased/regex');

require('unicode-6.3.0/properties/Changes_When_Uppercased/code-points');
require('unicode-6.3.0/properties/Changes_When_Uppercased/symbols');
require('unicode-6.3.0/properties/Changes_When_Uppercased/regex');

require('unicode-6.3.0/properties/Changes_When_Titlecased/code-points');
require('unicode-6.3.0/properties/Changes_When_Titlecased/symbols');
require('unicode-6.3.0/properties/Changes_When_Titlecased/regex');

require('unicode-6.3.0/properties/Changes_When_Casefolded/code-points');
require('unicode-6.3.0/properties/Changes_When_Casefolded/symbols');
require('unicode-6.3.0/properties/Changes_When_Casefolded/regex');

require('unicode-6.3.0/properties/Changes_When_Casemapped/code-points');
require('unicode-6.3.0/properties/Changes_When_Casemapped/symbols');
require('unicode-6.3.0/properties/Changes_When_Casemapped/regex');

require('unicode-6.3.0/properties/ID_Start/code-points');
require('unicode-6.3.0/properties/ID_Start/symbols');
require('unicode-6.3.0/properties/ID_Start/regex');

require('unicode-6.3.0/properties/ID_Continue/code-points');
require('unicode-6.3.0/properties/ID_Continue/symbols');
require('unicode-6.3.0/properties/ID_Continue/regex');

require('unicode-6.3.0/properties/XID_Start/code-points');
require('unicode-6.3.0/properties/XID_Start/symbols');
require('unicode-6.3.0/properties/XID_Start/regex');

require('unicode-6.3.0/properties/XID_Continue/code-points');
require('unicode-6.3.0/properties/XID_Continue/symbols');
require('unicode-6.3.0/properties/XID_Continue/regex');

require('unicode-6.3.0/properties/Default_Ignorable_Code_Point/code-points');
require('unicode-6.3.0/properties/Default_Ignorable_Code_Point/symbols');
require('unicode-6.3.0/properties/Default_Ignorable_Code_Point/regex');

require('unicode-6.3.0/properties/Grapheme_Extend/code-points');
require('unicode-6.3.0/properties/Grapheme_Extend/symbols');
require('unicode-6.3.0/properties/Grapheme_Extend/regex');

require('unicode-6.3.0/properties/Grapheme_Base/code-points');
require('unicode-6.3.0/properties/Grapheme_Base/symbols');
require('unicode-6.3.0/properties/Grapheme_Base/regex');

require('unicode-6.3.0/properties/Grapheme_Link/code-points');
require('unicode-6.3.0/properties/Grapheme_Link/symbols');
require('unicode-6.3.0/properties/Grapheme_Link/regex');

// categories:

require('unicode-6.3.0/categories')[ codePoint ]; // lookup array

require('unicode-6.3.0/categories/Cc/code-points');
require('unicode-6.3.0/categories/Cc/symbols');
require('unicode-6.3.0/categories/Cc/regex');

require('unicode-6.3.0/categories/C/code-points');
require('unicode-6.3.0/categories/C/symbols');
require('unicode-6.3.0/categories/C/regex');

require('unicode-6.3.0/categories/Zs/code-points');
require('unicode-6.3.0/categories/Zs/symbols');
require('unicode-6.3.0/categories/Zs/regex');

require('unicode-6.3.0/categories/Z/code-points');
require('unicode-6.3.0/categories/Z/symbols');
require('unicode-6.3.0/categories/Z/regex');

require('unicode-6.3.0/categories/Po/code-points');
require('unicode-6.3.0/categories/Po/symbols');
require('unicode-6.3.0/categories/Po/regex');

require('unicode-6.3.0/categories/P/code-points');
require('unicode-6.3.0/categories/P/symbols');
require('unicode-6.3.0/categories/P/regex');

require('unicode-6.3.0/categories/Sc/code-points');
require('unicode-6.3.0/categories/Sc/symbols');
require('unicode-6.3.0/categories/Sc/regex');

require('unicode-6.3.0/categories/S/code-points');
require('unicode-6.3.0/categories/S/symbols');
require('unicode-6.3.0/categories/S/regex');

require('unicode-6.3.0/categories/Ps/code-points');
require('unicode-6.3.0/categories/Ps/symbols');
require('unicode-6.3.0/categories/Ps/regex');

require('unicode-6.3.0/categories/Pe/code-points');
require('unicode-6.3.0/categories/Pe/symbols');
require('unicode-6.3.0/categories/Pe/regex');

require('unicode-6.3.0/categories/Sm/code-points');
require('unicode-6.3.0/categories/Sm/symbols');
require('unicode-6.3.0/categories/Sm/regex');

require('unicode-6.3.0/categories/Pd/code-points');
require('unicode-6.3.0/categories/Pd/symbols');
require('unicode-6.3.0/categories/Pd/regex');

require('unicode-6.3.0/categories/Nd/code-points');
require('unicode-6.3.0/categories/Nd/symbols');
require('unicode-6.3.0/categories/Nd/regex');

require('unicode-6.3.0/categories/N/code-points');
require('unicode-6.3.0/categories/N/symbols');
require('unicode-6.3.0/categories/N/regex');

require('unicode-6.3.0/categories/Lu/code-points');
require('unicode-6.3.0/categories/Lu/symbols');
require('unicode-6.3.0/categories/Lu/regex');

require('unicode-6.3.0/categories/L/code-points');
require('unicode-6.3.0/categories/L/symbols');
require('unicode-6.3.0/categories/L/regex');

require('unicode-6.3.0/categories/LC/code-points');
require('unicode-6.3.0/categories/LC/symbols');
require('unicode-6.3.0/categories/LC/regex');

require('unicode-6.3.0/categories/Sk/code-points');
require('unicode-6.3.0/categories/Sk/symbols');
require('unicode-6.3.0/categories/Sk/regex');

require('unicode-6.3.0/categories/Pc/code-points');
require('unicode-6.3.0/categories/Pc/symbols');
require('unicode-6.3.0/categories/Pc/regex');

require('unicode-6.3.0/categories/Ll/code-points');
require('unicode-6.3.0/categories/Ll/symbols');
require('unicode-6.3.0/categories/Ll/regex');

require('unicode-6.3.0/categories/So/code-points');
require('unicode-6.3.0/categories/So/symbols');
require('unicode-6.3.0/categories/So/regex');

require('unicode-6.3.0/categories/Lo/code-points');
require('unicode-6.3.0/categories/Lo/symbols');
require('unicode-6.3.0/categories/Lo/regex');

require('unicode-6.3.0/categories/Pi/code-points');
require('unicode-6.3.0/categories/Pi/symbols');
require('unicode-6.3.0/categories/Pi/regex');

require('unicode-6.3.0/categories/Cf/code-points');
require('unicode-6.3.0/categories/Cf/symbols');
require('unicode-6.3.0/categories/Cf/regex');

require('unicode-6.3.0/categories/No/code-points');
require('unicode-6.3.0/categories/No/symbols');
require('unicode-6.3.0/categories/No/regex');

require('unicode-6.3.0/categories/Pf/code-points');
require('unicode-6.3.0/categories/Pf/symbols');
require('unicode-6.3.0/categories/Pf/regex');

require('unicode-6.3.0/categories/Lt/code-points');
require('unicode-6.3.0/categories/Lt/symbols');
require('unicode-6.3.0/categories/Lt/regex');

require('unicode-6.3.0/categories/Lm/code-points');
require('unicode-6.3.0/categories/Lm/symbols');
require('unicode-6.3.0/categories/Lm/regex');

require('unicode-6.3.0/categories/Mn/code-points');
require('unicode-6.3.0/categories/Mn/symbols');
require('unicode-6.3.0/categories/Mn/regex');

require('unicode-6.3.0/categories/M/code-points');
require('unicode-6.3.0/categories/M/symbols');
require('unicode-6.3.0/categories/M/regex');

require('unicode-6.3.0/categories/Cn/code-points');
require('unicode-6.3.0/categories/Cn/symbols');
require('unicode-6.3.0/categories/Cn/regex');

require('unicode-6.3.0/categories/Me/code-points');
require('unicode-6.3.0/categories/Me/symbols');
require('unicode-6.3.0/categories/Me/regex');

require('unicode-6.3.0/categories/Mc/code-points');
require('unicode-6.3.0/categories/Mc/symbols');
require('unicode-6.3.0/categories/Mc/regex');

require('unicode-6.3.0/categories/Nl/code-points');
require('unicode-6.3.0/categories/Nl/symbols');
require('unicode-6.3.0/categories/Nl/regex');

require('unicode-6.3.0/categories/Zl/code-points');
require('unicode-6.3.0/categories/Zl/symbols');
require('unicode-6.3.0/categories/Zl/regex');

require('unicode-6.3.0/categories/Zp/code-points');
require('unicode-6.3.0/categories/Zp/symbols');
require('unicode-6.3.0/categories/Zp/regex');

require('unicode-6.3.0/categories/Cs/code-points');
require('unicode-6.3.0/categories/Cs/symbols');
require('unicode-6.3.0/categories/Cs/regex');

require('unicode-6.3.0/categories/Co/code-points');
require('unicode-6.3.0/categories/Co/symbols');
require('unicode-6.3.0/categories/Co/regex');

// bidi:

require('unicode-6.3.0/bidi')[ codePoint ]; // lookup array

require('unicode-6.3.0/bidi/BN/code-points');
require('unicode-6.3.0/bidi/BN/symbols');
require('unicode-6.3.0/bidi/BN/regex');

require('unicode-6.3.0/bidi/S/code-points');
require('unicode-6.3.0/bidi/S/symbols');
require('unicode-6.3.0/bidi/S/regex');

require('unicode-6.3.0/bidi/B/code-points');
require('unicode-6.3.0/bidi/B/symbols');
require('unicode-6.3.0/bidi/B/regex');

require('unicode-6.3.0/bidi/WS/code-points');
require('unicode-6.3.0/bidi/WS/symbols');
require('unicode-6.3.0/bidi/WS/regex');

require('unicode-6.3.0/bidi/ON/code-points');
require('unicode-6.3.0/bidi/ON/symbols');
require('unicode-6.3.0/bidi/ON/regex');

require('unicode-6.3.0/bidi/ET/code-points');
require('unicode-6.3.0/bidi/ET/symbols');
require('unicode-6.3.0/bidi/ET/regex');

require('unicode-6.3.0/bidi/ES/code-points');
require('unicode-6.3.0/bidi/ES/symbols');
require('unicode-6.3.0/bidi/ES/regex');

require('unicode-6.3.0/bidi/CS/code-points');
require('unicode-6.3.0/bidi/CS/symbols');
require('unicode-6.3.0/bidi/CS/regex');

require('unicode-6.3.0/bidi/EN/code-points');
require('unicode-6.3.0/bidi/EN/symbols');
require('unicode-6.3.0/bidi/EN/regex');

require('unicode-6.3.0/bidi/L/code-points');
require('unicode-6.3.0/bidi/L/symbols');
require('unicode-6.3.0/bidi/L/regex');

require('unicode-6.3.0/bidi/NSM/code-points');
require('unicode-6.3.0/bidi/NSM/symbols');
require('unicode-6.3.0/bidi/NSM/regex');

require('unicode-6.3.0/bidi/R/code-points');
require('unicode-6.3.0/bidi/R/symbols');
require('unicode-6.3.0/bidi/R/regex');

require('unicode-6.3.0/bidi/AN/code-points');
require('unicode-6.3.0/bidi/AN/symbols');
require('unicode-6.3.0/bidi/AN/regex');

require('unicode-6.3.0/bidi/AL/code-points');
require('unicode-6.3.0/bidi/AL/symbols');
require('unicode-6.3.0/bidi/AL/regex');

require('unicode-6.3.0/bidi/LRE/code-points');
require('unicode-6.3.0/bidi/LRE/symbols');
require('unicode-6.3.0/bidi/LRE/regex');

require('unicode-6.3.0/bidi/RLE/code-points');
require('unicode-6.3.0/bidi/RLE/symbols');
require('unicode-6.3.0/bidi/RLE/regex');

require('unicode-6.3.0/bidi/PDF/code-points');
require('unicode-6.3.0/bidi/PDF/symbols');
require('unicode-6.3.0/bidi/PDF/regex');

require('unicode-6.3.0/bidi/LRO/code-points');
require('unicode-6.3.0/bidi/LRO/symbols');
require('unicode-6.3.0/bidi/LRO/regex');

require('unicode-6.3.0/bidi/RLO/code-points');
require('unicode-6.3.0/bidi/RLO/symbols');
require('unicode-6.3.0/bidi/RLO/regex');

require('unicode-6.3.0/bidi/LRI/code-points');
require('unicode-6.3.0/bidi/LRI/symbols');
require('unicode-6.3.0/bidi/LRI/regex');

require('unicode-6.3.0/bidi/RLI/code-points');
require('unicode-6.3.0/bidi/RLI/symbols');
require('unicode-6.3.0/bidi/RLI/regex');

require('unicode-6.3.0/bidi/FSI/code-points');
require('unicode-6.3.0/bidi/FSI/symbols');
require('unicode-6.3.0/bidi/FSI/regex');

require('unicode-6.3.0/bidi/PDI/code-points');
require('unicode-6.3.0/bidi/PDI/symbols');
require('unicode-6.3.0/bidi/PDI/regex');

// scripts:

require('unicode-6.3.0/scripts/Common/code-points');
require('unicode-6.3.0/scripts/Common/symbols');
require('unicode-6.3.0/scripts/Common/regex');

require('unicode-6.3.0/scripts/Latin/code-points');
require('unicode-6.3.0/scripts/Latin/symbols');
require('unicode-6.3.0/scripts/Latin/regex');

require('unicode-6.3.0/scripts/Greek/code-points');
require('unicode-6.3.0/scripts/Greek/symbols');
require('unicode-6.3.0/scripts/Greek/regex');

require('unicode-6.3.0/scripts/Cyrillic/code-points');
require('unicode-6.3.0/scripts/Cyrillic/symbols');
require('unicode-6.3.0/scripts/Cyrillic/regex');

require('unicode-6.3.0/scripts/Armenian/code-points');
require('unicode-6.3.0/scripts/Armenian/symbols');
require('unicode-6.3.0/scripts/Armenian/regex');

require('unicode-6.3.0/scripts/Hebrew/code-points');
require('unicode-6.3.0/scripts/Hebrew/symbols');
require('unicode-6.3.0/scripts/Hebrew/regex');

require('unicode-6.3.0/scripts/Arabic/code-points');
require('unicode-6.3.0/scripts/Arabic/symbols');
require('unicode-6.3.0/scripts/Arabic/regex');

require('unicode-6.3.0/scripts/Syriac/code-points');
require('unicode-6.3.0/scripts/Syriac/symbols');
require('unicode-6.3.0/scripts/Syriac/regex');

require('unicode-6.3.0/scripts/Thaana/code-points');
require('unicode-6.3.0/scripts/Thaana/symbols');
require('unicode-6.3.0/scripts/Thaana/regex');

require('unicode-6.3.0/scripts/Devanagari/code-points');
require('unicode-6.3.0/scripts/Devanagari/symbols');
require('unicode-6.3.0/scripts/Devanagari/regex');

require('unicode-6.3.0/scripts/Bengali/code-points');
require('unicode-6.3.0/scripts/Bengali/symbols');
require('unicode-6.3.0/scripts/Bengali/regex');

require('unicode-6.3.0/scripts/Gurmukhi/code-points');
require('unicode-6.3.0/scripts/Gurmukhi/symbols');
require('unicode-6.3.0/scripts/Gurmukhi/regex');

require('unicode-6.3.0/scripts/Gujarati/code-points');
require('unicode-6.3.0/scripts/Gujarati/symbols');
require('unicode-6.3.0/scripts/Gujarati/regex');

require('unicode-6.3.0/scripts/Oriya/code-points');
require('unicode-6.3.0/scripts/Oriya/symbols');
require('unicode-6.3.0/scripts/Oriya/regex');

require('unicode-6.3.0/scripts/Tamil/code-points');
require('unicode-6.3.0/scripts/Tamil/symbols');
require('unicode-6.3.0/scripts/Tamil/regex');

require('unicode-6.3.0/scripts/Telugu/code-points');
require('unicode-6.3.0/scripts/Telugu/symbols');
require('unicode-6.3.0/scripts/Telugu/regex');

require('unicode-6.3.0/scripts/Kannada/code-points');
require('unicode-6.3.0/scripts/Kannada/symbols');
require('unicode-6.3.0/scripts/Kannada/regex');

require('unicode-6.3.0/scripts/Malayalam/code-points');
require('unicode-6.3.0/scripts/Malayalam/symbols');
require('unicode-6.3.0/scripts/Malayalam/regex');

require('unicode-6.3.0/scripts/Sinhala/code-points');
require('unicode-6.3.0/scripts/Sinhala/symbols');
require('unicode-6.3.0/scripts/Sinhala/regex');

require('unicode-6.3.0/scripts/Thai/code-points');
require('unicode-6.3.0/scripts/Thai/symbols');
require('unicode-6.3.0/scripts/Thai/regex');

require('unicode-6.3.0/scripts/Lao/code-points');
require('unicode-6.3.0/scripts/Lao/symbols');
require('unicode-6.3.0/scripts/Lao/regex');

require('unicode-6.3.0/scripts/Tibetan/code-points');
require('unicode-6.3.0/scripts/Tibetan/symbols');
require('unicode-6.3.0/scripts/Tibetan/regex');

require('unicode-6.3.0/scripts/Myanmar/code-points');
require('unicode-6.3.0/scripts/Myanmar/symbols');
require('unicode-6.3.0/scripts/Myanmar/regex');

require('unicode-6.3.0/scripts/Georgian/code-points');
require('unicode-6.3.0/scripts/Georgian/symbols');
require('unicode-6.3.0/scripts/Georgian/regex');

require('unicode-6.3.0/scripts/Hangul/code-points');
require('unicode-6.3.0/scripts/Hangul/symbols');
require('unicode-6.3.0/scripts/Hangul/regex');

require('unicode-6.3.0/scripts/Ethiopic/code-points');
require('unicode-6.3.0/scripts/Ethiopic/symbols');
require('unicode-6.3.0/scripts/Ethiopic/regex');

require('unicode-6.3.0/scripts/Cherokee/code-points');
require('unicode-6.3.0/scripts/Cherokee/symbols');
require('unicode-6.3.0/scripts/Cherokee/regex');

require('unicode-6.3.0/scripts/Canadian_Aboriginal/code-points');
require('unicode-6.3.0/scripts/Canadian_Aboriginal/symbols');
require('unicode-6.3.0/scripts/Canadian_Aboriginal/regex');

require('unicode-6.3.0/scripts/Ogham/code-points');
require('unicode-6.3.0/scripts/Ogham/symbols');
require('unicode-6.3.0/scripts/Ogham/regex');

require('unicode-6.3.0/scripts/Runic/code-points');
require('unicode-6.3.0/scripts/Runic/symbols');
require('unicode-6.3.0/scripts/Runic/regex');

require('unicode-6.3.0/scripts/Khmer/code-points');
require('unicode-6.3.0/scripts/Khmer/symbols');
require('unicode-6.3.0/scripts/Khmer/regex');

require('unicode-6.3.0/scripts/Mongolian/code-points');
require('unicode-6.3.0/scripts/Mongolian/symbols');
require('unicode-6.3.0/scripts/Mongolian/regex');

require('unicode-6.3.0/scripts/Hiragana/code-points');
require('unicode-6.3.0/scripts/Hiragana/symbols');
require('unicode-6.3.0/scripts/Hiragana/regex');

require('unicode-6.3.0/scripts/Katakana/code-points');
require('unicode-6.3.0/scripts/Katakana/symbols');
require('unicode-6.3.0/scripts/Katakana/regex');

require('unicode-6.3.0/scripts/Bopomofo/code-points');
require('unicode-6.3.0/scripts/Bopomofo/symbols');
require('unicode-6.3.0/scripts/Bopomofo/regex');

require('unicode-6.3.0/scripts/Han/code-points');
require('unicode-6.3.0/scripts/Han/symbols');
require('unicode-6.3.0/scripts/Han/regex');

require('unicode-6.3.0/scripts/Yi/code-points');
require('unicode-6.3.0/scripts/Yi/symbols');
require('unicode-6.3.0/scripts/Yi/regex');

require('unicode-6.3.0/scripts/Old_Italic/code-points');
require('unicode-6.3.0/scripts/Old_Italic/symbols');
require('unicode-6.3.0/scripts/Old_Italic/regex');

require('unicode-6.3.0/scripts/Gothic/code-points');
require('unicode-6.3.0/scripts/Gothic/symbols');
require('unicode-6.3.0/scripts/Gothic/regex');

require('unicode-6.3.0/scripts/Deseret/code-points');
require('unicode-6.3.0/scripts/Deseret/symbols');
require('unicode-6.3.0/scripts/Deseret/regex');

require('unicode-6.3.0/scripts/Inherited/code-points');
require('unicode-6.3.0/scripts/Inherited/symbols');
require('unicode-6.3.0/scripts/Inherited/regex');

require('unicode-6.3.0/scripts/Tagalog/code-points');
require('unicode-6.3.0/scripts/Tagalog/symbols');
require('unicode-6.3.0/scripts/Tagalog/regex');

require('unicode-6.3.0/scripts/Hanunoo/code-points');
require('unicode-6.3.0/scripts/Hanunoo/symbols');
require('unicode-6.3.0/scripts/Hanunoo/regex');

require('unicode-6.3.0/scripts/Buhid/code-points');
require('unicode-6.3.0/scripts/Buhid/symbols');
require('unicode-6.3.0/scripts/Buhid/regex');

require('unicode-6.3.0/scripts/Tagbanwa/code-points');
require('unicode-6.3.0/scripts/Tagbanwa/symbols');
require('unicode-6.3.0/scripts/Tagbanwa/regex');

require('unicode-6.3.0/scripts/Limbu/code-points');
require('unicode-6.3.0/scripts/Limbu/symbols');
require('unicode-6.3.0/scripts/Limbu/regex');

require('unicode-6.3.0/scripts/Tai_Le/code-points');
require('unicode-6.3.0/scripts/Tai_Le/symbols');
require('unicode-6.3.0/scripts/Tai_Le/regex');

require('unicode-6.3.0/scripts/Linear_B/code-points');
require('unicode-6.3.0/scripts/Linear_B/symbols');
require('unicode-6.3.0/scripts/Linear_B/regex');

require('unicode-6.3.0/scripts/Ugaritic/code-points');
require('unicode-6.3.0/scripts/Ugaritic/symbols');
require('unicode-6.3.0/scripts/Ugaritic/regex');

require('unicode-6.3.0/scripts/Shavian/code-points');
require('unicode-6.3.0/scripts/Shavian/symbols');
require('unicode-6.3.0/scripts/Shavian/regex');

require('unicode-6.3.0/scripts/Osmanya/code-points');
require('unicode-6.3.0/scripts/Osmanya/symbols');
require('unicode-6.3.0/scripts/Osmanya/regex');

require('unicode-6.3.0/scripts/Cypriot/code-points');
require('unicode-6.3.0/scripts/Cypriot/symbols');
require('unicode-6.3.0/scripts/Cypriot/regex');

require('unicode-6.3.0/scripts/Braille/code-points');
require('unicode-6.3.0/scripts/Braille/symbols');
require('unicode-6.3.0/scripts/Braille/regex');

require('unicode-6.3.0/scripts/Buginese/code-points');
require('unicode-6.3.0/scripts/Buginese/symbols');
require('unicode-6.3.0/scripts/Buginese/regex');

require('unicode-6.3.0/scripts/Coptic/code-points');
require('unicode-6.3.0/scripts/Coptic/symbols');
require('unicode-6.3.0/scripts/Coptic/regex');

require('unicode-6.3.0/scripts/New_Tai_Lue/code-points');
require('unicode-6.3.0/scripts/New_Tai_Lue/symbols');
require('unicode-6.3.0/scripts/New_Tai_Lue/regex');

require('unicode-6.3.0/scripts/Glagolitic/code-points');
require('unicode-6.3.0/scripts/Glagolitic/symbols');
require('unicode-6.3.0/scripts/Glagolitic/regex');

require('unicode-6.3.0/scripts/Tifinagh/code-points');
require('unicode-6.3.0/scripts/Tifinagh/symbols');
require('unicode-6.3.0/scripts/Tifinagh/regex');

require('unicode-6.3.0/scripts/Syloti_Nagri/code-points');
require('unicode-6.3.0/scripts/Syloti_Nagri/symbols');
require('unicode-6.3.0/scripts/Syloti_Nagri/regex');

require('unicode-6.3.0/scripts/Old_Persian/code-points');
require('unicode-6.3.0/scripts/Old_Persian/symbols');
require('unicode-6.3.0/scripts/Old_Persian/regex');

require('unicode-6.3.0/scripts/Kharoshthi/code-points');
require('unicode-6.3.0/scripts/Kharoshthi/symbols');
require('unicode-6.3.0/scripts/Kharoshthi/regex');

require('unicode-6.3.0/scripts/Balinese/code-points');
require('unicode-6.3.0/scripts/Balinese/symbols');
require('unicode-6.3.0/scripts/Balinese/regex');

require('unicode-6.3.0/scripts/Cuneiform/code-points');
require('unicode-6.3.0/scripts/Cuneiform/symbols');
require('unicode-6.3.0/scripts/Cuneiform/regex');

require('unicode-6.3.0/scripts/Phoenician/code-points');
require('unicode-6.3.0/scripts/Phoenician/symbols');
require('unicode-6.3.0/scripts/Phoenician/regex');

require('unicode-6.3.0/scripts/Phags_Pa/code-points');
require('unicode-6.3.0/scripts/Phags_Pa/symbols');
require('unicode-6.3.0/scripts/Phags_Pa/regex');

require('unicode-6.3.0/scripts/Nko/code-points');
require('unicode-6.3.0/scripts/Nko/symbols');
require('unicode-6.3.0/scripts/Nko/regex');

require('unicode-6.3.0/scripts/Sundanese/code-points');
require('unicode-6.3.0/scripts/Sundanese/symbols');
require('unicode-6.3.0/scripts/Sundanese/regex');

require('unicode-6.3.0/scripts/Lepcha/code-points');
require('unicode-6.3.0/scripts/Lepcha/symbols');
require('unicode-6.3.0/scripts/Lepcha/regex');

require('unicode-6.3.0/scripts/Ol_Chiki/code-points');
require('unicode-6.3.0/scripts/Ol_Chiki/symbols');
require('unicode-6.3.0/scripts/Ol_Chiki/regex');

require('unicode-6.3.0/scripts/Vai/code-points');
require('unicode-6.3.0/scripts/Vai/symbols');
require('unicode-6.3.0/scripts/Vai/regex');

require('unicode-6.3.0/scripts/Saurashtra/code-points');
require('unicode-6.3.0/scripts/Saurashtra/symbols');
require('unicode-6.3.0/scripts/Saurashtra/regex');

require('unicode-6.3.0/scripts/Kayah_Li/code-points');
require('unicode-6.3.0/scripts/Kayah_Li/symbols');
require('unicode-6.3.0/scripts/Kayah_Li/regex');

require('unicode-6.3.0/scripts/Rejang/code-points');
require('unicode-6.3.0/scripts/Rejang/symbols');
require('unicode-6.3.0/scripts/Rejang/regex');

require('unicode-6.3.0/scripts/Lycian/code-points');
require('unicode-6.3.0/scripts/Lycian/symbols');
require('unicode-6.3.0/scripts/Lycian/regex');

require('unicode-6.3.0/scripts/Carian/code-points');
require('unicode-6.3.0/scripts/Carian/symbols');
require('unicode-6.3.0/scripts/Carian/regex');

require('unicode-6.3.0/scripts/Lydian/code-points');
require('unicode-6.3.0/scripts/Lydian/symbols');
require('unicode-6.3.0/scripts/Lydian/regex');

require('unicode-6.3.0/scripts/Cham/code-points');
require('unicode-6.3.0/scripts/Cham/symbols');
require('unicode-6.3.0/scripts/Cham/regex');

require('unicode-6.3.0/scripts/Tai_Tham/code-points');
require('unicode-6.3.0/scripts/Tai_Tham/symbols');
require('unicode-6.3.0/scripts/Tai_Tham/regex');

require('unicode-6.3.0/scripts/Tai_Viet/code-points');
require('unicode-6.3.0/scripts/Tai_Viet/symbols');
require('unicode-6.3.0/scripts/Tai_Viet/regex');

require('unicode-6.3.0/scripts/Avestan/code-points');
require('unicode-6.3.0/scripts/Avestan/symbols');
require('unicode-6.3.0/scripts/Avestan/regex');

require('unicode-6.3.0/scripts/Egyptian_Hieroglyphs/code-points');
require('unicode-6.3.0/scripts/Egyptian_Hieroglyphs/symbols');
require('unicode-6.3.0/scripts/Egyptian_Hieroglyphs/regex');

require('unicode-6.3.0/scripts/Samaritan/code-points');
require('unicode-6.3.0/scripts/Samaritan/symbols');
require('unicode-6.3.0/scripts/Samaritan/regex');

require('unicode-6.3.0/scripts/Lisu/code-points');
require('unicode-6.3.0/scripts/Lisu/symbols');
require('unicode-6.3.0/scripts/Lisu/regex');

require('unicode-6.3.0/scripts/Bamum/code-points');
require('unicode-6.3.0/scripts/Bamum/symbols');
require('unicode-6.3.0/scripts/Bamum/regex');

require('unicode-6.3.0/scripts/Javanese/code-points');
require('unicode-6.3.0/scripts/Javanese/symbols');
require('unicode-6.3.0/scripts/Javanese/regex');

require('unicode-6.3.0/scripts/Meetei_Mayek/code-points');
require('unicode-6.3.0/scripts/Meetei_Mayek/symbols');
require('unicode-6.3.0/scripts/Meetei_Mayek/regex');

require('unicode-6.3.0/scripts/Imperial_Aramaic/code-points');
require('unicode-6.3.0/scripts/Imperial_Aramaic/symbols');
require('unicode-6.3.0/scripts/Imperial_Aramaic/regex');

require('unicode-6.3.0/scripts/Old_South_Arabian/code-points');
require('unicode-6.3.0/scripts/Old_South_Arabian/symbols');
require('unicode-6.3.0/scripts/Old_South_Arabian/regex');

require('unicode-6.3.0/scripts/Inscriptional_Parthian/code-points');
require('unicode-6.3.0/scripts/Inscriptional_Parthian/symbols');
require('unicode-6.3.0/scripts/Inscriptional_Parthian/regex');

require('unicode-6.3.0/scripts/Inscriptional_Pahlavi/code-points');
require('unicode-6.3.0/scripts/Inscriptional_Pahlavi/symbols');
require('unicode-6.3.0/scripts/Inscriptional_Pahlavi/regex');

require('unicode-6.3.0/scripts/Old_Turkic/code-points');
require('unicode-6.3.0/scripts/Old_Turkic/symbols');
require('unicode-6.3.0/scripts/Old_Turkic/regex');

require('unicode-6.3.0/scripts/Kaithi/code-points');
require('unicode-6.3.0/scripts/Kaithi/symbols');
require('unicode-6.3.0/scripts/Kaithi/regex');

require('unicode-6.3.0/scripts/Batak/code-points');
require('unicode-6.3.0/scripts/Batak/symbols');
require('unicode-6.3.0/scripts/Batak/regex');

require('unicode-6.3.0/scripts/Brahmi/code-points');
require('unicode-6.3.0/scripts/Brahmi/symbols');
require('unicode-6.3.0/scripts/Brahmi/regex');

require('unicode-6.3.0/scripts/Mandaic/code-points');
require('unicode-6.3.0/scripts/Mandaic/symbols');
require('unicode-6.3.0/scripts/Mandaic/regex');

require('unicode-6.3.0/scripts/Chakma/code-points');
require('unicode-6.3.0/scripts/Chakma/symbols');
require('unicode-6.3.0/scripts/Chakma/regex');

require('unicode-6.3.0/scripts/Meroitic_Cursive/code-points');
require('unicode-6.3.0/scripts/Meroitic_Cursive/symbols');
require('unicode-6.3.0/scripts/Meroitic_Cursive/regex');

require('unicode-6.3.0/scripts/Meroitic_Hieroglyphs/code-points');
require('unicode-6.3.0/scripts/Meroitic_Hieroglyphs/symbols');
require('unicode-6.3.0/scripts/Meroitic_Hieroglyphs/regex');

require('unicode-6.3.0/scripts/Miao/code-points');
require('unicode-6.3.0/scripts/Miao/symbols');
require('unicode-6.3.0/scripts/Miao/regex');

require('unicode-6.3.0/scripts/Sharada/code-points');
require('unicode-6.3.0/scripts/Sharada/symbols');
require('unicode-6.3.0/scripts/Sharada/regex');

require('unicode-6.3.0/scripts/Sora_Sompeng/code-points');
require('unicode-6.3.0/scripts/Sora_Sompeng/symbols');
require('unicode-6.3.0/scripts/Sora_Sompeng/regex');

require('unicode-6.3.0/scripts/Takri/code-points');
require('unicode-6.3.0/scripts/Takri/symbols');
require('unicode-6.3.0/scripts/Takri/regex');

// script extensions:

require('unicode-6.3.0/script-extensions/Common/code-points');
require('unicode-6.3.0/script-extensions/Common/symbols');
require('unicode-6.3.0/script-extensions/Common/regex');

require('unicode-6.3.0/script-extensions/Latin/code-points');
require('unicode-6.3.0/script-extensions/Latin/symbols');
require('unicode-6.3.0/script-extensions/Latin/regex');

require('unicode-6.3.0/script-extensions/Greek/code-points');
require('unicode-6.3.0/script-extensions/Greek/symbols');
require('unicode-6.3.0/script-extensions/Greek/regex');

require('unicode-6.3.0/script-extensions/Cyrillic/code-points');
require('unicode-6.3.0/script-extensions/Cyrillic/symbols');
require('unicode-6.3.0/script-extensions/Cyrillic/regex');

require('unicode-6.3.0/script-extensions/Armenian/code-points');
require('unicode-6.3.0/script-extensions/Armenian/symbols');
require('unicode-6.3.0/script-extensions/Armenian/regex');

require('unicode-6.3.0/script-extensions/Hebrew/code-points');
require('unicode-6.3.0/script-extensions/Hebrew/symbols');
require('unicode-6.3.0/script-extensions/Hebrew/regex');

require('unicode-6.3.0/script-extensions/Arabic/code-points');
require('unicode-6.3.0/script-extensions/Arabic/symbols');
require('unicode-6.3.0/script-extensions/Arabic/regex');

require('unicode-6.3.0/script-extensions/Syriac/code-points');
require('unicode-6.3.0/script-extensions/Syriac/symbols');
require('unicode-6.3.0/script-extensions/Syriac/regex');

require('unicode-6.3.0/script-extensions/Thaana/code-points');
require('unicode-6.3.0/script-extensions/Thaana/symbols');
require('unicode-6.3.0/script-extensions/Thaana/regex');

require('unicode-6.3.0/script-extensions/Devanagari/code-points');
require('unicode-6.3.0/script-extensions/Devanagari/symbols');
require('unicode-6.3.0/script-extensions/Devanagari/regex');

require('unicode-6.3.0/script-extensions/Bengali/code-points');
require('unicode-6.3.0/script-extensions/Bengali/symbols');
require('unicode-6.3.0/script-extensions/Bengali/regex');

require('unicode-6.3.0/script-extensions/Gurmukhi/code-points');
require('unicode-6.3.0/script-extensions/Gurmukhi/symbols');
require('unicode-6.3.0/script-extensions/Gurmukhi/regex');

require('unicode-6.3.0/script-extensions/Gujarati/code-points');
require('unicode-6.3.0/script-extensions/Gujarati/symbols');
require('unicode-6.3.0/script-extensions/Gujarati/regex');

require('unicode-6.3.0/script-extensions/Oriya/code-points');
require('unicode-6.3.0/script-extensions/Oriya/symbols');
require('unicode-6.3.0/script-extensions/Oriya/regex');

require('unicode-6.3.0/script-extensions/Tamil/code-points');
require('unicode-6.3.0/script-extensions/Tamil/symbols');
require('unicode-6.3.0/script-extensions/Tamil/regex');

require('unicode-6.3.0/script-extensions/Telugu/code-points');
require('unicode-6.3.0/script-extensions/Telugu/symbols');
require('unicode-6.3.0/script-extensions/Telugu/regex');

require('unicode-6.3.0/script-extensions/Kannada/code-points');
require('unicode-6.3.0/script-extensions/Kannada/symbols');
require('unicode-6.3.0/script-extensions/Kannada/regex');

require('unicode-6.3.0/script-extensions/Malayalam/code-points');
require('unicode-6.3.0/script-extensions/Malayalam/symbols');
require('unicode-6.3.0/script-extensions/Malayalam/regex');

require('unicode-6.3.0/script-extensions/Sinhala/code-points');
require('unicode-6.3.0/script-extensions/Sinhala/symbols');
require('unicode-6.3.0/script-extensions/Sinhala/regex');

require('unicode-6.3.0/script-extensions/Thai/code-points');
require('unicode-6.3.0/script-extensions/Thai/symbols');
require('unicode-6.3.0/script-extensions/Thai/regex');

require('unicode-6.3.0/script-extensions/Lao/code-points');
require('unicode-6.3.0/script-extensions/Lao/symbols');
require('unicode-6.3.0/script-extensions/Lao/regex');

require('unicode-6.3.0/script-extensions/Tibetan/code-points');
require('unicode-6.3.0/script-extensions/Tibetan/symbols');
require('unicode-6.3.0/script-extensions/Tibetan/regex');

require('unicode-6.3.0/script-extensions/Myanmar/code-points');
require('unicode-6.3.0/script-extensions/Myanmar/symbols');
require('unicode-6.3.0/script-extensions/Myanmar/regex');

require('unicode-6.3.0/script-extensions/Georgian/code-points');
require('unicode-6.3.0/script-extensions/Georgian/symbols');
require('unicode-6.3.0/script-extensions/Georgian/regex');

require('unicode-6.3.0/script-extensions/Hangul/code-points');
require('unicode-6.3.0/script-extensions/Hangul/symbols');
require('unicode-6.3.0/script-extensions/Hangul/regex');

require('unicode-6.3.0/script-extensions/Ethiopic/code-points');
require('unicode-6.3.0/script-extensions/Ethiopic/symbols');
require('unicode-6.3.0/script-extensions/Ethiopic/regex');

require('unicode-6.3.0/script-extensions/Cherokee/code-points');
require('unicode-6.3.0/script-extensions/Cherokee/symbols');
require('unicode-6.3.0/script-extensions/Cherokee/regex');

require('unicode-6.3.0/script-extensions/Canadian_Aboriginal/code-points');
require('unicode-6.3.0/script-extensions/Canadian_Aboriginal/symbols');
require('unicode-6.3.0/script-extensions/Canadian_Aboriginal/regex');

require('unicode-6.3.0/script-extensions/Ogham/code-points');
require('unicode-6.3.0/script-extensions/Ogham/symbols');
require('unicode-6.3.0/script-extensions/Ogham/regex');

require('unicode-6.3.0/script-extensions/Runic/code-points');
require('unicode-6.3.0/script-extensions/Runic/symbols');
require('unicode-6.3.0/script-extensions/Runic/regex');

require('unicode-6.3.0/script-extensions/Khmer/code-points');
require('unicode-6.3.0/script-extensions/Khmer/symbols');
require('unicode-6.3.0/script-extensions/Khmer/regex');

require('unicode-6.3.0/script-extensions/Mongolian/code-points');
require('unicode-6.3.0/script-extensions/Mongolian/symbols');
require('unicode-6.3.0/script-extensions/Mongolian/regex');

require('unicode-6.3.0/script-extensions/Hiragana/code-points');
require('unicode-6.3.0/script-extensions/Hiragana/symbols');
require('unicode-6.3.0/script-extensions/Hiragana/regex');

require('unicode-6.3.0/script-extensions/Katakana/code-points');
require('unicode-6.3.0/script-extensions/Katakana/symbols');
require('unicode-6.3.0/script-extensions/Katakana/regex');

require('unicode-6.3.0/script-extensions/Bopomofo/code-points');
require('unicode-6.3.0/script-extensions/Bopomofo/symbols');
require('unicode-6.3.0/script-extensions/Bopomofo/regex');

require('unicode-6.3.0/script-extensions/Han/code-points');
require('unicode-6.3.0/script-extensions/Han/symbols');
require('unicode-6.3.0/script-extensions/Han/regex');

require('unicode-6.3.0/script-extensions/Yi/code-points');
require('unicode-6.3.0/script-extensions/Yi/symbols');
require('unicode-6.3.0/script-extensions/Yi/regex');

require('unicode-6.3.0/script-extensions/Old_Italic/code-points');
require('unicode-6.3.0/script-extensions/Old_Italic/symbols');
require('unicode-6.3.0/script-extensions/Old_Italic/regex');

require('unicode-6.3.0/script-extensions/Gothic/code-points');
require('unicode-6.3.0/script-extensions/Gothic/symbols');
require('unicode-6.3.0/script-extensions/Gothic/regex');

require('unicode-6.3.0/script-extensions/Deseret/code-points');
require('unicode-6.3.0/script-extensions/Deseret/symbols');
require('unicode-6.3.0/script-extensions/Deseret/regex');

require('unicode-6.3.0/script-extensions/Inherited/code-points');
require('unicode-6.3.0/script-extensions/Inherited/symbols');
require('unicode-6.3.0/script-extensions/Inherited/regex');

require('unicode-6.3.0/script-extensions/Tagalog/code-points');
require('unicode-6.3.0/script-extensions/Tagalog/symbols');
require('unicode-6.3.0/script-extensions/Tagalog/regex');

require('unicode-6.3.0/script-extensions/Hanunoo/code-points');
require('unicode-6.3.0/script-extensions/Hanunoo/symbols');
require('unicode-6.3.0/script-extensions/Hanunoo/regex');

require('unicode-6.3.0/script-extensions/Buhid/code-points');
require('unicode-6.3.0/script-extensions/Buhid/symbols');
require('unicode-6.3.0/script-extensions/Buhid/regex');

require('unicode-6.3.0/script-extensions/Tagbanwa/code-points');
require('unicode-6.3.0/script-extensions/Tagbanwa/symbols');
require('unicode-6.3.0/script-extensions/Tagbanwa/regex');

require('unicode-6.3.0/script-extensions/Limbu/code-points');
require('unicode-6.3.0/script-extensions/Limbu/symbols');
require('unicode-6.3.0/script-extensions/Limbu/regex');

require('unicode-6.3.0/script-extensions/Tai_Le/code-points');
require('unicode-6.3.0/script-extensions/Tai_Le/symbols');
require('unicode-6.3.0/script-extensions/Tai_Le/regex');

require('unicode-6.3.0/script-extensions/Linear_B/code-points');
require('unicode-6.3.0/script-extensions/Linear_B/symbols');
require('unicode-6.3.0/script-extensions/Linear_B/regex');

require('unicode-6.3.0/script-extensions/Ugaritic/code-points');
require('unicode-6.3.0/script-extensions/Ugaritic/symbols');
require('unicode-6.3.0/script-extensions/Ugaritic/regex');

require('unicode-6.3.0/script-extensions/Shavian/code-points');
require('unicode-6.3.0/script-extensions/Shavian/symbols');
require('unicode-6.3.0/script-extensions/Shavian/regex');

require('unicode-6.3.0/script-extensions/Osmanya/code-points');
require('unicode-6.3.0/script-extensions/Osmanya/symbols');
require('unicode-6.3.0/script-extensions/Osmanya/regex');

require('unicode-6.3.0/script-extensions/Cypriot/code-points');
require('unicode-6.3.0/script-extensions/Cypriot/symbols');
require('unicode-6.3.0/script-extensions/Cypriot/regex');

require('unicode-6.3.0/script-extensions/Braille/code-points');
require('unicode-6.3.0/script-extensions/Braille/symbols');
require('unicode-6.3.0/script-extensions/Braille/regex');

require('unicode-6.3.0/script-extensions/Buginese/code-points');
require('unicode-6.3.0/script-extensions/Buginese/symbols');
require('unicode-6.3.0/script-extensions/Buginese/regex');

require('unicode-6.3.0/script-extensions/Coptic/code-points');
require('unicode-6.3.0/script-extensions/Coptic/symbols');
require('unicode-6.3.0/script-extensions/Coptic/regex');

require('unicode-6.3.0/script-extensions/New_Tai_Lue/code-points');
require('unicode-6.3.0/script-extensions/New_Tai_Lue/symbols');
require('unicode-6.3.0/script-extensions/New_Tai_Lue/regex');

require('unicode-6.3.0/script-extensions/Glagolitic/code-points');
require('unicode-6.3.0/script-extensions/Glagolitic/symbols');
require('unicode-6.3.0/script-extensions/Glagolitic/regex');

require('unicode-6.3.0/script-extensions/Tifinagh/code-points');
require('unicode-6.3.0/script-extensions/Tifinagh/symbols');
require('unicode-6.3.0/script-extensions/Tifinagh/regex');

require('unicode-6.3.0/script-extensions/Syloti_Nagri/code-points');
require('unicode-6.3.0/script-extensions/Syloti_Nagri/symbols');
require('unicode-6.3.0/script-extensions/Syloti_Nagri/regex');

require('unicode-6.3.0/script-extensions/Old_Persian/code-points');
require('unicode-6.3.0/script-extensions/Old_Persian/symbols');
require('unicode-6.3.0/script-extensions/Old_Persian/regex');

require('unicode-6.3.0/script-extensions/Kharoshthi/code-points');
require('unicode-6.3.0/script-extensions/Kharoshthi/symbols');
require('unicode-6.3.0/script-extensions/Kharoshthi/regex');

require('unicode-6.3.0/script-extensions/Balinese/code-points');
require('unicode-6.3.0/script-extensions/Balinese/symbols');
require('unicode-6.3.0/script-extensions/Balinese/regex');

require('unicode-6.3.0/script-extensions/Cuneiform/code-points');
require('unicode-6.3.0/script-extensions/Cuneiform/symbols');
require('unicode-6.3.0/script-extensions/Cuneiform/regex');

require('unicode-6.3.0/script-extensions/Phoenician/code-points');
require('unicode-6.3.0/script-extensions/Phoenician/symbols');
require('unicode-6.3.0/script-extensions/Phoenician/regex');

require('unicode-6.3.0/script-extensions/Phags_Pa/code-points');
require('unicode-6.3.0/script-extensions/Phags_Pa/symbols');
require('unicode-6.3.0/script-extensions/Phags_Pa/regex');

require('unicode-6.3.0/script-extensions/Nko/code-points');
require('unicode-6.3.0/script-extensions/Nko/symbols');
require('unicode-6.3.0/script-extensions/Nko/regex');

require('unicode-6.3.0/script-extensions/Sundanese/code-points');
require('unicode-6.3.0/script-extensions/Sundanese/symbols');
require('unicode-6.3.0/script-extensions/Sundanese/regex');

require('unicode-6.3.0/script-extensions/Lepcha/code-points');
require('unicode-6.3.0/script-extensions/Lepcha/symbols');
require('unicode-6.3.0/script-extensions/Lepcha/regex');

require('unicode-6.3.0/script-extensions/Ol_Chiki/code-points');
require('unicode-6.3.0/script-extensions/Ol_Chiki/symbols');
require('unicode-6.3.0/script-extensions/Ol_Chiki/regex');

require('unicode-6.3.0/script-extensions/Vai/code-points');
require('unicode-6.3.0/script-extensions/Vai/symbols');
require('unicode-6.3.0/script-extensions/Vai/regex');

require('unicode-6.3.0/script-extensions/Saurashtra/code-points');
require('unicode-6.3.0/script-extensions/Saurashtra/symbols');
require('unicode-6.3.0/script-extensions/Saurashtra/regex');

require('unicode-6.3.0/script-extensions/Kayah_Li/code-points');
require('unicode-6.3.0/script-extensions/Kayah_Li/symbols');
require('unicode-6.3.0/script-extensions/Kayah_Li/regex');

require('unicode-6.3.0/script-extensions/Rejang/code-points');
require('unicode-6.3.0/script-extensions/Rejang/symbols');
require('unicode-6.3.0/script-extensions/Rejang/regex');

require('unicode-6.3.0/script-extensions/Lycian/code-points');
require('unicode-6.3.0/script-extensions/Lycian/symbols');
require('unicode-6.3.0/script-extensions/Lycian/regex');

require('unicode-6.3.0/script-extensions/Carian/code-points');
require('unicode-6.3.0/script-extensions/Carian/symbols');
require('unicode-6.3.0/script-extensions/Carian/regex');

require('unicode-6.3.0/script-extensions/Lydian/code-points');
require('unicode-6.3.0/script-extensions/Lydian/symbols');
require('unicode-6.3.0/script-extensions/Lydian/regex');

require('unicode-6.3.0/script-extensions/Cham/code-points');
require('unicode-6.3.0/script-extensions/Cham/symbols');
require('unicode-6.3.0/script-extensions/Cham/regex');

require('unicode-6.3.0/script-extensions/Tai_Tham/code-points');
require('unicode-6.3.0/script-extensions/Tai_Tham/symbols');
require('unicode-6.3.0/script-extensions/Tai_Tham/regex');

require('unicode-6.3.0/script-extensions/Tai_Viet/code-points');
require('unicode-6.3.0/script-extensions/Tai_Viet/symbols');
require('unicode-6.3.0/script-extensions/Tai_Viet/regex');

require('unicode-6.3.0/script-extensions/Avestan/code-points');
require('unicode-6.3.0/script-extensions/Avestan/symbols');
require('unicode-6.3.0/script-extensions/Avestan/regex');

require('unicode-6.3.0/script-extensions/Egyptian_Hieroglyphs/code-points');
require('unicode-6.3.0/script-extensions/Egyptian_Hieroglyphs/symbols');
require('unicode-6.3.0/script-extensions/Egyptian_Hieroglyphs/regex');

require('unicode-6.3.0/script-extensions/Samaritan/code-points');
require('unicode-6.3.0/script-extensions/Samaritan/symbols');
require('unicode-6.3.0/script-extensions/Samaritan/regex');

require('unicode-6.3.0/script-extensions/Lisu/code-points');
require('unicode-6.3.0/script-extensions/Lisu/symbols');
require('unicode-6.3.0/script-extensions/Lisu/regex');

require('unicode-6.3.0/script-extensions/Bamum/code-points');
require('unicode-6.3.0/script-extensions/Bamum/symbols');
require('unicode-6.3.0/script-extensions/Bamum/regex');

require('unicode-6.3.0/script-extensions/Javanese/code-points');
require('unicode-6.3.0/script-extensions/Javanese/symbols');
require('unicode-6.3.0/script-extensions/Javanese/regex');

require('unicode-6.3.0/script-extensions/Meetei_Mayek/code-points');
require('unicode-6.3.0/script-extensions/Meetei_Mayek/symbols');
require('unicode-6.3.0/script-extensions/Meetei_Mayek/regex');

require('unicode-6.3.0/script-extensions/Imperial_Aramaic/code-points');
require('unicode-6.3.0/script-extensions/Imperial_Aramaic/symbols');
require('unicode-6.3.0/script-extensions/Imperial_Aramaic/regex');

require('unicode-6.3.0/script-extensions/Old_South_Arabian/code-points');
require('unicode-6.3.0/script-extensions/Old_South_Arabian/symbols');
require('unicode-6.3.0/script-extensions/Old_South_Arabian/regex');

require('unicode-6.3.0/script-extensions/Inscriptional_Parthian/code-points');
require('unicode-6.3.0/script-extensions/Inscriptional_Parthian/symbols');
require('unicode-6.3.0/script-extensions/Inscriptional_Parthian/regex');

require('unicode-6.3.0/script-extensions/Inscriptional_Pahlavi/code-points');
require('unicode-6.3.0/script-extensions/Inscriptional_Pahlavi/symbols');
require('unicode-6.3.0/script-extensions/Inscriptional_Pahlavi/regex');

require('unicode-6.3.0/script-extensions/Old_Turkic/code-points');
require('unicode-6.3.0/script-extensions/Old_Turkic/symbols');
require('unicode-6.3.0/script-extensions/Old_Turkic/regex');

require('unicode-6.3.0/script-extensions/Kaithi/code-points');
require('unicode-6.3.0/script-extensions/Kaithi/symbols');
require('unicode-6.3.0/script-extensions/Kaithi/regex');

require('unicode-6.3.0/script-extensions/Batak/code-points');
require('unicode-6.3.0/script-extensions/Batak/symbols');
require('unicode-6.3.0/script-extensions/Batak/regex');

require('unicode-6.3.0/script-extensions/Brahmi/code-points');
require('unicode-6.3.0/script-extensions/Brahmi/symbols');
require('unicode-6.3.0/script-extensions/Brahmi/regex');

require('unicode-6.3.0/script-extensions/Mandaic/code-points');
require('unicode-6.3.0/script-extensions/Mandaic/symbols');
require('unicode-6.3.0/script-extensions/Mandaic/regex');

require('unicode-6.3.0/script-extensions/Chakma/code-points');
require('unicode-6.3.0/script-extensions/Chakma/symbols');
require('unicode-6.3.0/script-extensions/Chakma/regex');

require('unicode-6.3.0/script-extensions/Meroitic_Cursive/code-points');
require('unicode-6.3.0/script-extensions/Meroitic_Cursive/symbols');
require('unicode-6.3.0/script-extensions/Meroitic_Cursive/regex');

require('unicode-6.3.0/script-extensions/Meroitic_Hieroglyphs/code-points');
require('unicode-6.3.0/script-extensions/Meroitic_Hieroglyphs/symbols');
require('unicode-6.3.0/script-extensions/Meroitic_Hieroglyphs/regex');

require('unicode-6.3.0/script-extensions/Miao/code-points');
require('unicode-6.3.0/script-extensions/Miao/symbols');
require('unicode-6.3.0/script-extensions/Miao/regex');

require('unicode-6.3.0/script-extensions/Sharada/code-points');
require('unicode-6.3.0/script-extensions/Sharada/symbols');
require('unicode-6.3.0/script-extensions/Sharada/regex');

require('unicode-6.3.0/script-extensions/Sora_Sompeng/code-points');
require('unicode-6.3.0/script-extensions/Sora_Sompeng/symbols');
require('unicode-6.3.0/script-extensions/Sora_Sompeng/regex');

require('unicode-6.3.0/script-extensions/Takri/code-points');
require('unicode-6.3.0/script-extensions/Takri/symbols');
require('unicode-6.3.0/script-extensions/Takri/regex');

// case folding:

require('unicode-6.3.0/case-folding/C/code-points'); // lookup table with code point to code point mappings
require('unicode-6.3.0/case-folding/C/code-points')[ codePoint ]; // lookup table
require('unicode-6.3.0/case-folding/C/symbols'); // lookup table with symbol to symbol mappings
require('unicode-6.3.0/case-folding/C/symbols')[ symbol ]; // lookup table

require('unicode-6.3.0/case-folding/T/code-points'); // lookup table with code point to code point mappings
require('unicode-6.3.0/case-folding/T/code-points')[ codePoint ]; // lookup table
require('unicode-6.3.0/case-folding/T/symbols'); // lookup table with symbol to symbol mappings
require('unicode-6.3.0/case-folding/T/symbols')[ symbol ]; // lookup table

require('unicode-6.3.0/case-folding/F/code-points'); // lookup table with code point to code point mappings
require('unicode-6.3.0/case-folding/F/code-points')[ codePoint ]; // lookup table
require('unicode-6.3.0/case-folding/F/symbols'); // lookup table with symbol to symbol mappings
require('unicode-6.3.0/case-folding/F/symbols')[ symbol ]; // lookup table

require('unicode-6.3.0/case-folding/S/code-points'); // lookup table with code point to code point mappings
require('unicode-6.3.0/case-folding/S/code-points')[ codePoint ]; // lookup table
require('unicode-6.3.0/case-folding/S/symbols'); // lookup table with symbol to symbol mappings
require('unicode-6.3.0/case-folding/S/symbols')[ symbol ]; // lookup table

// blocks:

require('unicode-6.3.0/blocks/Basic Latin/code-points');
require('unicode-6.3.0/blocks/Basic Latin/symbols');
require('unicode-6.3.0/blocks/Basic Latin/regex');

require('unicode-6.3.0/blocks/Latin-1 Supplement/code-points');
require('unicode-6.3.0/blocks/Latin-1 Supplement/symbols');
require('unicode-6.3.0/blocks/Latin-1 Supplement/regex');

require('unicode-6.3.0/blocks/Latin Extended-A/code-points');
require('unicode-6.3.0/blocks/Latin Extended-A/symbols');
require('unicode-6.3.0/blocks/Latin Extended-A/regex');

require('unicode-6.3.0/blocks/Latin Extended-B/code-points');
require('unicode-6.3.0/blocks/Latin Extended-B/symbols');
require('unicode-6.3.0/blocks/Latin Extended-B/regex');

require('unicode-6.3.0/blocks/IPA Extensions/code-points');
require('unicode-6.3.0/blocks/IPA Extensions/symbols');
require('unicode-6.3.0/blocks/IPA Extensions/regex');

require('unicode-6.3.0/blocks/Spacing Modifier Letters/code-points');
require('unicode-6.3.0/blocks/Spacing Modifier Letters/symbols');
require('unicode-6.3.0/blocks/Spacing Modifier Letters/regex');

require('unicode-6.3.0/blocks/Combining Diacritical Marks/code-points');
require('unicode-6.3.0/blocks/Combining Diacritical Marks/symbols');
require('unicode-6.3.0/blocks/Combining Diacritical Marks/regex');

require('unicode-6.3.0/blocks/Greek and Coptic/code-points');
require('unicode-6.3.0/blocks/Greek and Coptic/symbols');
require('unicode-6.3.0/blocks/Greek and Coptic/regex');

require('unicode-6.3.0/blocks/Cyrillic/code-points');
require('unicode-6.3.0/blocks/Cyrillic/symbols');
require('unicode-6.3.0/blocks/Cyrillic/regex');

require('unicode-6.3.0/blocks/Cyrillic Supplement/code-points');
require('unicode-6.3.0/blocks/Cyrillic Supplement/symbols');
require('unicode-6.3.0/blocks/Cyrillic Supplement/regex');

require('unicode-6.3.0/blocks/Armenian/code-points');
require('unicode-6.3.0/blocks/Armenian/symbols');
require('unicode-6.3.0/blocks/Armenian/regex');

require('unicode-6.3.0/blocks/Hebrew/code-points');
require('unicode-6.3.0/blocks/Hebrew/symbols');
require('unicode-6.3.0/blocks/Hebrew/regex');

require('unicode-6.3.0/blocks/Arabic/code-points');
require('unicode-6.3.0/blocks/Arabic/symbols');
require('unicode-6.3.0/blocks/Arabic/regex');

require('unicode-6.3.0/blocks/Syriac/code-points');
require('unicode-6.3.0/blocks/Syriac/symbols');
require('unicode-6.3.0/blocks/Syriac/regex');

require('unicode-6.3.0/blocks/Arabic Supplement/code-points');
require('unicode-6.3.0/blocks/Arabic Supplement/symbols');
require('unicode-6.3.0/blocks/Arabic Supplement/regex');

require('unicode-6.3.0/blocks/Thaana/code-points');
require('unicode-6.3.0/blocks/Thaana/symbols');
require('unicode-6.3.0/blocks/Thaana/regex');

require('unicode-6.3.0/blocks/NKo/code-points');
require('unicode-6.3.0/blocks/NKo/symbols');
require('unicode-6.3.0/blocks/NKo/regex');

require('unicode-6.3.0/blocks/Samaritan/code-points');
require('unicode-6.3.0/blocks/Samaritan/symbols');
require('unicode-6.3.0/blocks/Samaritan/regex');

require('unicode-6.3.0/blocks/Mandaic/code-points');
require('unicode-6.3.0/blocks/Mandaic/symbols');
require('unicode-6.3.0/blocks/Mandaic/regex');

require('unicode-6.3.0/blocks/Arabic Extended-A/code-points');
require('unicode-6.3.0/blocks/Arabic Extended-A/symbols');
require('unicode-6.3.0/blocks/Arabic Extended-A/regex');

require('unicode-6.3.0/blocks/Devanagari/code-points');
require('unicode-6.3.0/blocks/Devanagari/symbols');
require('unicode-6.3.0/blocks/Devanagari/regex');

require('unicode-6.3.0/blocks/Bengali/code-points');
require('unicode-6.3.0/blocks/Bengali/symbols');
require('unicode-6.3.0/blocks/Bengali/regex');

require('unicode-6.3.0/blocks/Gurmukhi/code-points');
require('unicode-6.3.0/blocks/Gurmukhi/symbols');
require('unicode-6.3.0/blocks/Gurmukhi/regex');

require('unicode-6.3.0/blocks/Gujarati/code-points');
require('unicode-6.3.0/blocks/Gujarati/symbols');
require('unicode-6.3.0/blocks/Gujarati/regex');

require('unicode-6.3.0/blocks/Oriya/code-points');
require('unicode-6.3.0/blocks/Oriya/symbols');
require('unicode-6.3.0/blocks/Oriya/regex');

require('unicode-6.3.0/blocks/Tamil/code-points');
require('unicode-6.3.0/blocks/Tamil/symbols');
require('unicode-6.3.0/blocks/Tamil/regex');

require('unicode-6.3.0/blocks/Telugu/code-points');
require('unicode-6.3.0/blocks/Telugu/symbols');
require('unicode-6.3.0/blocks/Telugu/regex');

require('unicode-6.3.0/blocks/Kannada/code-points');
require('unicode-6.3.0/blocks/Kannada/symbols');
require('unicode-6.3.0/blocks/Kannada/regex');

require('unicode-6.3.0/blocks/Malayalam/code-points');
require('unicode-6.3.0/blocks/Malayalam/symbols');
require('unicode-6.3.0/blocks/Malayalam/regex');

require('unicode-6.3.0/blocks/Sinhala/code-points');
require('unicode-6.3.0/blocks/Sinhala/symbols');
require('unicode-6.3.0/blocks/Sinhala/regex');

require('unicode-6.3.0/blocks/Thai/code-points');
require('unicode-6.3.0/blocks/Thai/symbols');
require('unicode-6.3.0/blocks/Thai/regex');

require('unicode-6.3.0/blocks/Lao/code-points');
require('unicode-6.3.0/blocks/Lao/symbols');
require('unicode-6.3.0/blocks/Lao/regex');

require('unicode-6.3.0/blocks/Tibetan/code-points');
require('unicode-6.3.0/blocks/Tibetan/symbols');
require('unicode-6.3.0/blocks/Tibetan/regex');

require('unicode-6.3.0/blocks/Myanmar/code-points');
require('unicode-6.3.0/blocks/Myanmar/symbols');
require('unicode-6.3.0/blocks/Myanmar/regex');

require('unicode-6.3.0/blocks/Georgian/code-points');
require('unicode-6.3.0/blocks/Georgian/symbols');
require('unicode-6.3.0/blocks/Georgian/regex');

require('unicode-6.3.0/blocks/Hangul Jamo/code-points');
require('unicode-6.3.0/blocks/Hangul Jamo/symbols');
require('unicode-6.3.0/blocks/Hangul Jamo/regex');

require('unicode-6.3.0/blocks/Ethiopic/code-points');
require('unicode-6.3.0/blocks/Ethiopic/symbols');
require('unicode-6.3.0/blocks/Ethiopic/regex');

require('unicode-6.3.0/blocks/Ethiopic Supplement/code-points');
require('unicode-6.3.0/blocks/Ethiopic Supplement/symbols');
require('unicode-6.3.0/blocks/Ethiopic Supplement/regex');

require('unicode-6.3.0/blocks/Cherokee/code-points');
require('unicode-6.3.0/blocks/Cherokee/symbols');
require('unicode-6.3.0/blocks/Cherokee/regex');

require('unicode-6.3.0/blocks/Unified Canadian Aboriginal Syllabics/code-points');
require('unicode-6.3.0/blocks/Unified Canadian Aboriginal Syllabics/symbols');
require('unicode-6.3.0/blocks/Unified Canadian Aboriginal Syllabics/regex');

require('unicode-6.3.0/blocks/Ogham/code-points');
require('unicode-6.3.0/blocks/Ogham/symbols');
require('unicode-6.3.0/blocks/Ogham/regex');

require('unicode-6.3.0/blocks/Runic/code-points');
require('unicode-6.3.0/blocks/Runic/symbols');
require('unicode-6.3.0/blocks/Runic/regex');

require('unicode-6.3.0/blocks/Tagalog/code-points');
require('unicode-6.3.0/blocks/Tagalog/symbols');
require('unicode-6.3.0/blocks/Tagalog/regex');

require('unicode-6.3.0/blocks/Hanunoo/code-points');
require('unicode-6.3.0/blocks/Hanunoo/symbols');
require('unicode-6.3.0/blocks/Hanunoo/regex');

require('unicode-6.3.0/blocks/Buhid/code-points');
require('unicode-6.3.0/blocks/Buhid/symbols');
require('unicode-6.3.0/blocks/Buhid/regex');

require('unicode-6.3.0/blocks/Tagbanwa/code-points');
require('unicode-6.3.0/blocks/Tagbanwa/symbols');
require('unicode-6.3.0/blocks/Tagbanwa/regex');

require('unicode-6.3.0/blocks/Khmer/code-points');
require('unicode-6.3.0/blocks/Khmer/symbols');
require('unicode-6.3.0/blocks/Khmer/regex');

require('unicode-6.3.0/blocks/Mongolian/code-points');
require('unicode-6.3.0/blocks/Mongolian/symbols');
require('unicode-6.3.0/blocks/Mongolian/regex');

require('unicode-6.3.0/blocks/Unified Canadian Aboriginal Syllabics Extended/code-points');
require('unicode-6.3.0/blocks/Unified Canadian Aboriginal Syllabics Extended/symbols');
require('unicode-6.3.0/blocks/Unified Canadian Aboriginal Syllabics Extended/regex');

require('unicode-6.3.0/blocks/Limbu/code-points');
require('unicode-6.3.0/blocks/Limbu/symbols');
require('unicode-6.3.0/blocks/Limbu/regex');

require('unicode-6.3.0/blocks/Tai Le/code-points');
require('unicode-6.3.0/blocks/Tai Le/symbols');
require('unicode-6.3.0/blocks/Tai Le/regex');

require('unicode-6.3.0/blocks/New Tai Lue/code-points');
require('unicode-6.3.0/blocks/New Tai Lue/symbols');
require('unicode-6.3.0/blocks/New Tai Lue/regex');

require('unicode-6.3.0/blocks/Khmer Symbols/code-points');
require('unicode-6.3.0/blocks/Khmer Symbols/symbols');
require('unicode-6.3.0/blocks/Khmer Symbols/regex');

require('unicode-6.3.0/blocks/Buginese/code-points');
require('unicode-6.3.0/blocks/Buginese/symbols');
require('unicode-6.3.0/blocks/Buginese/regex');

require('unicode-6.3.0/blocks/Tai Tham/code-points');
require('unicode-6.3.0/blocks/Tai Tham/symbols');
require('unicode-6.3.0/blocks/Tai Tham/regex');

require('unicode-6.3.0/blocks/Balinese/code-points');
require('unicode-6.3.0/blocks/Balinese/symbols');
require('unicode-6.3.0/blocks/Balinese/regex');

require('unicode-6.3.0/blocks/Sundanese/code-points');
require('unicode-6.3.0/blocks/Sundanese/symbols');
require('unicode-6.3.0/blocks/Sundanese/regex');

require('unicode-6.3.0/blocks/Batak/code-points');
require('unicode-6.3.0/blocks/Batak/symbols');
require('unicode-6.3.0/blocks/Batak/regex');

require('unicode-6.3.0/blocks/Lepcha/code-points');
require('unicode-6.3.0/blocks/Lepcha/symbols');
require('unicode-6.3.0/blocks/Lepcha/regex');

require('unicode-6.3.0/blocks/Ol Chiki/code-points');
require('unicode-6.3.0/blocks/Ol Chiki/symbols');
require('unicode-6.3.0/blocks/Ol Chiki/regex');

require('unicode-6.3.0/blocks/Sundanese Supplement/code-points');
require('unicode-6.3.0/blocks/Sundanese Supplement/symbols');
require('unicode-6.3.0/blocks/Sundanese Supplement/regex');

require('unicode-6.3.0/blocks/Vedic Extensions/code-points');
require('unicode-6.3.0/blocks/Vedic Extensions/symbols');
require('unicode-6.3.0/blocks/Vedic Extensions/regex');

require('unicode-6.3.0/blocks/Phonetic Extensions/code-points');
require('unicode-6.3.0/blocks/Phonetic Extensions/symbols');
require('unicode-6.3.0/blocks/Phonetic Extensions/regex');

require('unicode-6.3.0/blocks/Phonetic Extensions Supplement/code-points');
require('unicode-6.3.0/blocks/Phonetic Extensions Supplement/symbols');
require('unicode-6.3.0/blocks/Phonetic Extensions Supplement/regex');

require('unicode-6.3.0/blocks/Combining Diacritical Marks Supplement/code-points');
require('unicode-6.3.0/blocks/Combining Diacritical Marks Supplement/symbols');
require('unicode-6.3.0/blocks/Combining Diacritical Marks Supplement/regex');

require('unicode-6.3.0/blocks/Latin Extended Additional/code-points');
require('unicode-6.3.0/blocks/Latin Extended Additional/symbols');
require('unicode-6.3.0/blocks/Latin Extended Additional/regex');

require('unicode-6.3.0/blocks/Greek Extended/code-points');
require('unicode-6.3.0/blocks/Greek Extended/symbols');
require('unicode-6.3.0/blocks/Greek Extended/regex');

require('unicode-6.3.0/blocks/General Punctuation/code-points');
require('unicode-6.3.0/blocks/General Punctuation/symbols');
require('unicode-6.3.0/blocks/General Punctuation/regex');

require('unicode-6.3.0/blocks/Superscripts and Subscripts/code-points');
require('unicode-6.3.0/blocks/Superscripts and Subscripts/symbols');
require('unicode-6.3.0/blocks/Superscripts and Subscripts/regex');

require('unicode-6.3.0/blocks/Currency Symbols/code-points');
require('unicode-6.3.0/blocks/Currency Symbols/symbols');
require('unicode-6.3.0/blocks/Currency Symbols/regex');

require('unicode-6.3.0/blocks/Combining Diacritical Marks for Symbols/code-points');
require('unicode-6.3.0/blocks/Combining Diacritical Marks for Symbols/symbols');
require('unicode-6.3.0/blocks/Combining Diacritical Marks for Symbols/regex');

require('unicode-6.3.0/blocks/Letterlike Symbols/code-points');
require('unicode-6.3.0/blocks/Letterlike Symbols/symbols');
require('unicode-6.3.0/blocks/Letterlike Symbols/regex');

require('unicode-6.3.0/blocks/Number Forms/code-points');
require('unicode-6.3.0/blocks/Number Forms/symbols');
require('unicode-6.3.0/blocks/Number Forms/regex');

require('unicode-6.3.0/blocks/Arrows/code-points');
require('unicode-6.3.0/blocks/Arrows/symbols');
require('unicode-6.3.0/blocks/Arrows/regex');

require('unicode-6.3.0/blocks/Mathematical Operators/code-points');
require('unicode-6.3.0/blocks/Mathematical Operators/symbols');
require('unicode-6.3.0/blocks/Mathematical Operators/regex');

require('unicode-6.3.0/blocks/Miscellaneous Technical/code-points');
require('unicode-6.3.0/blocks/Miscellaneous Technical/symbols');
require('unicode-6.3.0/blocks/Miscellaneous Technical/regex');

require('unicode-6.3.0/blocks/Control Pictures/code-points');
require('unicode-6.3.0/blocks/Control Pictures/symbols');
require('unicode-6.3.0/blocks/Control Pictures/regex');

require('unicode-6.3.0/blocks/Optical Character Recognition/code-points');
require('unicode-6.3.0/blocks/Optical Character Recognition/symbols');
require('unicode-6.3.0/blocks/Optical Character Recognition/regex');

require('unicode-6.3.0/blocks/Enclosed Alphanumerics/code-points');
require('unicode-6.3.0/blocks/Enclosed Alphanumerics/symbols');
require('unicode-6.3.0/blocks/Enclosed Alphanumerics/regex');

require('unicode-6.3.0/blocks/Box Drawing/code-points');
require('unicode-6.3.0/blocks/Box Drawing/symbols');
require('unicode-6.3.0/blocks/Box Drawing/regex');

require('unicode-6.3.0/blocks/Block Elements/code-points');
require('unicode-6.3.0/blocks/Block Elements/symbols');
require('unicode-6.3.0/blocks/Block Elements/regex');

require('unicode-6.3.0/blocks/Geometric Shapes/code-points');
require('unicode-6.3.0/blocks/Geometric Shapes/symbols');
require('unicode-6.3.0/blocks/Geometric Shapes/regex');

require('unicode-6.3.0/blocks/Miscellaneous Symbols/code-points');
require('unicode-6.3.0/blocks/Miscellaneous Symbols/symbols');
require('unicode-6.3.0/blocks/Miscellaneous Symbols/regex');

require('unicode-6.3.0/blocks/Dingbats/code-points');
require('unicode-6.3.0/blocks/Dingbats/symbols');
require('unicode-6.3.0/blocks/Dingbats/regex');

require('unicode-6.3.0/blocks/Miscellaneous Mathematical Symbols-A/code-points');
require('unicode-6.3.0/blocks/Miscellaneous Mathematical Symbols-A/symbols');
require('unicode-6.3.0/blocks/Miscellaneous Mathematical Symbols-A/regex');

require('unicode-6.3.0/blocks/Supplemental Arrows-A/code-points');
require('unicode-6.3.0/blocks/Supplemental Arrows-A/symbols');
require('unicode-6.3.0/blocks/Supplemental Arrows-A/regex');

require('unicode-6.3.0/blocks/Braille Patterns/code-points');
require('unicode-6.3.0/blocks/Braille Patterns/symbols');
require('unicode-6.3.0/blocks/Braille Patterns/regex');

require('unicode-6.3.0/blocks/Supplemental Arrows-B/code-points');
require('unicode-6.3.0/blocks/Supplemental Arrows-B/symbols');
require('unicode-6.3.0/blocks/Supplemental Arrows-B/regex');

require('unicode-6.3.0/blocks/Miscellaneous Mathematical Symbols-B/code-points');
require('unicode-6.3.0/blocks/Miscellaneous Mathematical Symbols-B/symbols');
require('unicode-6.3.0/blocks/Miscellaneous Mathematical Symbols-B/regex');

require('unicode-6.3.0/blocks/Supplemental Mathematical Operators/code-points');
require('unicode-6.3.0/blocks/Supplemental Mathematical Operators/symbols');
require('unicode-6.3.0/blocks/Supplemental Mathematical Operators/regex');

require('unicode-6.3.0/blocks/Miscellaneous Symbols and Arrows/code-points');
require('unicode-6.3.0/blocks/Miscellaneous Symbols and Arrows/symbols');
require('unicode-6.3.0/blocks/Miscellaneous Symbols and Arrows/regex');

require('unicode-6.3.0/blocks/Glagolitic/code-points');
require('unicode-6.3.0/blocks/Glagolitic/symbols');
require('unicode-6.3.0/blocks/Glagolitic/regex');

require('unicode-6.3.0/blocks/Latin Extended-C/code-points');
require('unicode-6.3.0/blocks/Latin Extended-C/symbols');
require('unicode-6.3.0/blocks/Latin Extended-C/regex');

require('unicode-6.3.0/blocks/Coptic/code-points');
require('unicode-6.3.0/blocks/Coptic/symbols');
require('unicode-6.3.0/blocks/Coptic/regex');

require('unicode-6.3.0/blocks/Georgian Supplement/code-points');
require('unicode-6.3.0/blocks/Georgian Supplement/symbols');
require('unicode-6.3.0/blocks/Georgian Supplement/regex');

require('unicode-6.3.0/blocks/Tifinagh/code-points');
require('unicode-6.3.0/blocks/Tifinagh/symbols');
require('unicode-6.3.0/blocks/Tifinagh/regex');

require('unicode-6.3.0/blocks/Ethiopic Extended/code-points');
require('unicode-6.3.0/blocks/Ethiopic Extended/symbols');
require('unicode-6.3.0/blocks/Ethiopic Extended/regex');

require('unicode-6.3.0/blocks/Cyrillic Extended-A/code-points');
require('unicode-6.3.0/blocks/Cyrillic Extended-A/symbols');
require('unicode-6.3.0/blocks/Cyrillic Extended-A/regex');

require('unicode-6.3.0/blocks/Supplemental Punctuation/code-points');
require('unicode-6.3.0/blocks/Supplemental Punctuation/symbols');
require('unicode-6.3.0/blocks/Supplemental Punctuation/regex');

require('unicode-6.3.0/blocks/CJK Radicals Supplement/code-points');
require('unicode-6.3.0/blocks/CJK Radicals Supplement/symbols');
require('unicode-6.3.0/blocks/CJK Radicals Supplement/regex');

require('unicode-6.3.0/blocks/Kangxi Radicals/code-points');
require('unicode-6.3.0/blocks/Kangxi Radicals/symbols');
require('unicode-6.3.0/blocks/Kangxi Radicals/regex');

require('unicode-6.3.0/blocks/Ideographic Description Characters/code-points');
require('unicode-6.3.0/blocks/Ideographic Description Characters/symbols');
require('unicode-6.3.0/blocks/Ideographic Description Characters/regex');

require('unicode-6.3.0/blocks/CJK Symbols and Punctuation/code-points');
require('unicode-6.3.0/blocks/CJK Symbols and Punctuation/symbols');
require('unicode-6.3.0/blocks/CJK Symbols and Punctuation/regex');

require('unicode-6.3.0/blocks/Hiragana/code-points');
require('unicode-6.3.0/blocks/Hiragana/symbols');
require('unicode-6.3.0/blocks/Hiragana/regex');

require('unicode-6.3.0/blocks/Katakana/code-points');
require('unicode-6.3.0/blocks/Katakana/symbols');
require('unicode-6.3.0/blocks/Katakana/regex');

require('unicode-6.3.0/blocks/Bopomofo/code-points');
require('unicode-6.3.0/blocks/Bopomofo/symbols');
require('unicode-6.3.0/blocks/Bopomofo/regex');

require('unicode-6.3.0/blocks/Hangul Compatibility Jamo/code-points');
require('unicode-6.3.0/blocks/Hangul Compatibility Jamo/symbols');
require('unicode-6.3.0/blocks/Hangul Compatibility Jamo/regex');

require('unicode-6.3.0/blocks/Kanbun/code-points');
require('unicode-6.3.0/blocks/Kanbun/symbols');
require('unicode-6.3.0/blocks/Kanbun/regex');

require('unicode-6.3.0/blocks/Bopomofo Extended/code-points');
require('unicode-6.3.0/blocks/Bopomofo Extended/symbols');
require('unicode-6.3.0/blocks/Bopomofo Extended/regex');

require('unicode-6.3.0/blocks/CJK Strokes/code-points');
require('unicode-6.3.0/blocks/CJK Strokes/symbols');
require('unicode-6.3.0/blocks/CJK Strokes/regex');

require('unicode-6.3.0/blocks/Katakana Phonetic Extensions/code-points');
require('unicode-6.3.0/blocks/Katakana Phonetic Extensions/symbols');
require('unicode-6.3.0/blocks/Katakana Phonetic Extensions/regex');

require('unicode-6.3.0/blocks/Enclosed CJK Letters and Months/code-points');
require('unicode-6.3.0/blocks/Enclosed CJK Letters and Months/symbols');
require('unicode-6.3.0/blocks/Enclosed CJK Letters and Months/regex');

require('unicode-6.3.0/blocks/CJK Compatibility/code-points');
require('unicode-6.3.0/blocks/CJK Compatibility/symbols');
require('unicode-6.3.0/blocks/CJK Compatibility/regex');

require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension A/code-points');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension A/symbols');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension A/regex');

require('unicode-6.3.0/blocks/Yijing Hexagram Symbols/code-points');
require('unicode-6.3.0/blocks/Yijing Hexagram Symbols/symbols');
require('unicode-6.3.0/blocks/Yijing Hexagram Symbols/regex');

require('unicode-6.3.0/blocks/CJK Unified Ideographs/code-points');
require('unicode-6.3.0/blocks/CJK Unified Ideographs/symbols');
require('unicode-6.3.0/blocks/CJK Unified Ideographs/regex');

require('unicode-6.3.0/blocks/Yi Syllables/code-points');
require('unicode-6.3.0/blocks/Yi Syllables/symbols');
require('unicode-6.3.0/blocks/Yi Syllables/regex');

require('unicode-6.3.0/blocks/Yi Radicals/code-points');
require('unicode-6.3.0/blocks/Yi Radicals/symbols');
require('unicode-6.3.0/blocks/Yi Radicals/regex');

require('unicode-6.3.0/blocks/Lisu/code-points');
require('unicode-6.3.0/blocks/Lisu/symbols');
require('unicode-6.3.0/blocks/Lisu/regex');

require('unicode-6.3.0/blocks/Vai/code-points');
require('unicode-6.3.0/blocks/Vai/symbols');
require('unicode-6.3.0/blocks/Vai/regex');

require('unicode-6.3.0/blocks/Cyrillic Extended-B/code-points');
require('unicode-6.3.0/blocks/Cyrillic Extended-B/symbols');
require('unicode-6.3.0/blocks/Cyrillic Extended-B/regex');

require('unicode-6.3.0/blocks/Bamum/code-points');
require('unicode-6.3.0/blocks/Bamum/symbols');
require('unicode-6.3.0/blocks/Bamum/regex');

require('unicode-6.3.0/blocks/Modifier Tone Letters/code-points');
require('unicode-6.3.0/blocks/Modifier Tone Letters/symbols');
require('unicode-6.3.0/blocks/Modifier Tone Letters/regex');

require('unicode-6.3.0/blocks/Latin Extended-D/code-points');
require('unicode-6.3.0/blocks/Latin Extended-D/symbols');
require('unicode-6.3.0/blocks/Latin Extended-D/regex');

require('unicode-6.3.0/blocks/Syloti Nagri/code-points');
require('unicode-6.3.0/blocks/Syloti Nagri/symbols');
require('unicode-6.3.0/blocks/Syloti Nagri/regex');

require('unicode-6.3.0/blocks/Common Indic Number Forms/code-points');
require('unicode-6.3.0/blocks/Common Indic Number Forms/symbols');
require('unicode-6.3.0/blocks/Common Indic Number Forms/regex');

require('unicode-6.3.0/blocks/Phags-pa/code-points');
require('unicode-6.3.0/blocks/Phags-pa/symbols');
require('unicode-6.3.0/blocks/Phags-pa/regex');

require('unicode-6.3.0/blocks/Saurashtra/code-points');
require('unicode-6.3.0/blocks/Saurashtra/symbols');
require('unicode-6.3.0/blocks/Saurashtra/regex');

require('unicode-6.3.0/blocks/Devanagari Extended/code-points');
require('unicode-6.3.0/blocks/Devanagari Extended/symbols');
require('unicode-6.3.0/blocks/Devanagari Extended/regex');

require('unicode-6.3.0/blocks/Kayah Li/code-points');
require('unicode-6.3.0/blocks/Kayah Li/symbols');
require('unicode-6.3.0/blocks/Kayah Li/regex');

require('unicode-6.3.0/blocks/Rejang/code-points');
require('unicode-6.3.0/blocks/Rejang/symbols');
require('unicode-6.3.0/blocks/Rejang/regex');

require('unicode-6.3.0/blocks/Hangul Jamo Extended-A/code-points');
require('unicode-6.3.0/blocks/Hangul Jamo Extended-A/symbols');
require('unicode-6.3.0/blocks/Hangul Jamo Extended-A/regex');

require('unicode-6.3.0/blocks/Javanese/code-points');
require('unicode-6.3.0/blocks/Javanese/symbols');
require('unicode-6.3.0/blocks/Javanese/regex');

require('unicode-6.3.0/blocks/Cham/code-points');
require('unicode-6.3.0/blocks/Cham/symbols');
require('unicode-6.3.0/blocks/Cham/regex');

require('unicode-6.3.0/blocks/Myanmar Extended-A/code-points');
require('unicode-6.3.0/blocks/Myanmar Extended-A/symbols');
require('unicode-6.3.0/blocks/Myanmar Extended-A/regex');

require('unicode-6.3.0/blocks/Tai Viet/code-points');
require('unicode-6.3.0/blocks/Tai Viet/symbols');
require('unicode-6.3.0/blocks/Tai Viet/regex');

require('unicode-6.3.0/blocks/Meetei Mayek Extensions/code-points');
require('unicode-6.3.0/blocks/Meetei Mayek Extensions/symbols');
require('unicode-6.3.0/blocks/Meetei Mayek Extensions/regex');

require('unicode-6.3.0/blocks/Ethiopic Extended-A/code-points');
require('unicode-6.3.0/blocks/Ethiopic Extended-A/symbols');
require('unicode-6.3.0/blocks/Ethiopic Extended-A/regex');

require('unicode-6.3.0/blocks/Meetei Mayek/code-points');
require('unicode-6.3.0/blocks/Meetei Mayek/symbols');
require('unicode-6.3.0/blocks/Meetei Mayek/regex');

require('unicode-6.3.0/blocks/Hangul Syllables/code-points');
require('unicode-6.3.0/blocks/Hangul Syllables/symbols');
require('unicode-6.3.0/blocks/Hangul Syllables/regex');

require('unicode-6.3.0/blocks/Hangul Jamo Extended-B/code-points');
require('unicode-6.3.0/blocks/Hangul Jamo Extended-B/symbols');
require('unicode-6.3.0/blocks/Hangul Jamo Extended-B/regex');

require('unicode-6.3.0/blocks/High Surrogates/code-points');
require('unicode-6.3.0/blocks/High Surrogates/symbols');
require('unicode-6.3.0/blocks/High Surrogates/regex');

require('unicode-6.3.0/blocks/High Private Use Surrogates/code-points');
require('unicode-6.3.0/blocks/High Private Use Surrogates/symbols');
require('unicode-6.3.0/blocks/High Private Use Surrogates/regex');

require('unicode-6.3.0/blocks/Low Surrogates/code-points');
require('unicode-6.3.0/blocks/Low Surrogates/symbols');
require('unicode-6.3.0/blocks/Low Surrogates/regex');

require('unicode-6.3.0/blocks/Private Use Area/code-points');
require('unicode-6.3.0/blocks/Private Use Area/symbols');
require('unicode-6.3.0/blocks/Private Use Area/regex');

require('unicode-6.3.0/blocks/CJK Compatibility Ideographs/code-points');
require('unicode-6.3.0/blocks/CJK Compatibility Ideographs/symbols');
require('unicode-6.3.0/blocks/CJK Compatibility Ideographs/regex');

require('unicode-6.3.0/blocks/Alphabetic Presentation Forms/code-points');
require('unicode-6.3.0/blocks/Alphabetic Presentation Forms/symbols');
require('unicode-6.3.0/blocks/Alphabetic Presentation Forms/regex');

require('unicode-6.3.0/blocks/Arabic Presentation Forms-A/code-points');
require('unicode-6.3.0/blocks/Arabic Presentation Forms-A/symbols');
require('unicode-6.3.0/blocks/Arabic Presentation Forms-A/regex');

require('unicode-6.3.0/blocks/Variation Selectors/code-points');
require('unicode-6.3.0/blocks/Variation Selectors/symbols');
require('unicode-6.3.0/blocks/Variation Selectors/regex');

require('unicode-6.3.0/blocks/Vertical Forms/code-points');
require('unicode-6.3.0/blocks/Vertical Forms/symbols');
require('unicode-6.3.0/blocks/Vertical Forms/regex');

require('unicode-6.3.0/blocks/Combining Half Marks/code-points');
require('unicode-6.3.0/blocks/Combining Half Marks/symbols');
require('unicode-6.3.0/blocks/Combining Half Marks/regex');

require('unicode-6.3.0/blocks/CJK Compatibility Forms/code-points');
require('unicode-6.3.0/blocks/CJK Compatibility Forms/symbols');
require('unicode-6.3.0/blocks/CJK Compatibility Forms/regex');

require('unicode-6.3.0/blocks/Small Form Variants/code-points');
require('unicode-6.3.0/blocks/Small Form Variants/symbols');
require('unicode-6.3.0/blocks/Small Form Variants/regex');

require('unicode-6.3.0/blocks/Arabic Presentation Forms-B/code-points');
require('unicode-6.3.0/blocks/Arabic Presentation Forms-B/symbols');
require('unicode-6.3.0/blocks/Arabic Presentation Forms-B/regex');

require('unicode-6.3.0/blocks/Halfwidth and Fullwidth Forms/code-points');
require('unicode-6.3.0/blocks/Halfwidth and Fullwidth Forms/symbols');
require('unicode-6.3.0/blocks/Halfwidth and Fullwidth Forms/regex');

require('unicode-6.3.0/blocks/Specials/code-points');
require('unicode-6.3.0/blocks/Specials/symbols');
require('unicode-6.3.0/blocks/Specials/regex');

require('unicode-6.3.0/blocks/Linear B Syllabary/code-points');
require('unicode-6.3.0/blocks/Linear B Syllabary/symbols');
require('unicode-6.3.0/blocks/Linear B Syllabary/regex');

require('unicode-6.3.0/blocks/Linear B Ideograms/code-points');
require('unicode-6.3.0/blocks/Linear B Ideograms/symbols');
require('unicode-6.3.0/blocks/Linear B Ideograms/regex');

require('unicode-6.3.0/blocks/Aegean Numbers/code-points');
require('unicode-6.3.0/blocks/Aegean Numbers/symbols');
require('unicode-6.3.0/blocks/Aegean Numbers/regex');

require('unicode-6.3.0/blocks/Ancient Greek Numbers/code-points');
require('unicode-6.3.0/blocks/Ancient Greek Numbers/symbols');
require('unicode-6.3.0/blocks/Ancient Greek Numbers/regex');

require('unicode-6.3.0/blocks/Ancient Symbols/code-points');
require('unicode-6.3.0/blocks/Ancient Symbols/symbols');
require('unicode-6.3.0/blocks/Ancient Symbols/regex');

require('unicode-6.3.0/blocks/Phaistos Disc/code-points');
require('unicode-6.3.0/blocks/Phaistos Disc/symbols');
require('unicode-6.3.0/blocks/Phaistos Disc/regex');

require('unicode-6.3.0/blocks/Lycian/code-points');
require('unicode-6.3.0/blocks/Lycian/symbols');
require('unicode-6.3.0/blocks/Lycian/regex');

require('unicode-6.3.0/blocks/Carian/code-points');
require('unicode-6.3.0/blocks/Carian/symbols');
require('unicode-6.3.0/blocks/Carian/regex');

require('unicode-6.3.0/blocks/Old Italic/code-points');
require('unicode-6.3.0/blocks/Old Italic/symbols');
require('unicode-6.3.0/blocks/Old Italic/regex');

require('unicode-6.3.0/blocks/Gothic/code-points');
require('unicode-6.3.0/blocks/Gothic/symbols');
require('unicode-6.3.0/blocks/Gothic/regex');

require('unicode-6.3.0/blocks/Ugaritic/code-points');
require('unicode-6.3.0/blocks/Ugaritic/symbols');
require('unicode-6.3.0/blocks/Ugaritic/regex');

require('unicode-6.3.0/blocks/Old Persian/code-points');
require('unicode-6.3.0/blocks/Old Persian/symbols');
require('unicode-6.3.0/blocks/Old Persian/regex');

require('unicode-6.3.0/blocks/Deseret/code-points');
require('unicode-6.3.0/blocks/Deseret/symbols');
require('unicode-6.3.0/blocks/Deseret/regex');

require('unicode-6.3.0/blocks/Shavian/code-points');
require('unicode-6.3.0/blocks/Shavian/symbols');
require('unicode-6.3.0/blocks/Shavian/regex');

require('unicode-6.3.0/blocks/Osmanya/code-points');
require('unicode-6.3.0/blocks/Osmanya/symbols');
require('unicode-6.3.0/blocks/Osmanya/regex');

require('unicode-6.3.0/blocks/Cypriot Syllabary/code-points');
require('unicode-6.3.0/blocks/Cypriot Syllabary/symbols');
require('unicode-6.3.0/blocks/Cypriot Syllabary/regex');

require('unicode-6.3.0/blocks/Imperial Aramaic/code-points');
require('unicode-6.3.0/blocks/Imperial Aramaic/symbols');
require('unicode-6.3.0/blocks/Imperial Aramaic/regex');

require('unicode-6.3.0/blocks/Phoenician/code-points');
require('unicode-6.3.0/blocks/Phoenician/symbols');
require('unicode-6.3.0/blocks/Phoenician/regex');

require('unicode-6.3.0/blocks/Lydian/code-points');
require('unicode-6.3.0/blocks/Lydian/symbols');
require('unicode-6.3.0/blocks/Lydian/regex');

require('unicode-6.3.0/blocks/Meroitic Hieroglyphs/code-points');
require('unicode-6.3.0/blocks/Meroitic Hieroglyphs/symbols');
require('unicode-6.3.0/blocks/Meroitic Hieroglyphs/regex');

require('unicode-6.3.0/blocks/Meroitic Cursive/code-points');
require('unicode-6.3.0/blocks/Meroitic Cursive/symbols');
require('unicode-6.3.0/blocks/Meroitic Cursive/regex');

require('unicode-6.3.0/blocks/Kharoshthi/code-points');
require('unicode-6.3.0/blocks/Kharoshthi/symbols');
require('unicode-6.3.0/blocks/Kharoshthi/regex');

require('unicode-6.3.0/blocks/Old South Arabian/code-points');
require('unicode-6.3.0/blocks/Old South Arabian/symbols');
require('unicode-6.3.0/blocks/Old South Arabian/regex');

require('unicode-6.3.0/blocks/Avestan/code-points');
require('unicode-6.3.0/blocks/Avestan/symbols');
require('unicode-6.3.0/blocks/Avestan/regex');

require('unicode-6.3.0/blocks/Inscriptional Parthian/code-points');
require('unicode-6.3.0/blocks/Inscriptional Parthian/symbols');
require('unicode-6.3.0/blocks/Inscriptional Parthian/regex');

require('unicode-6.3.0/blocks/Inscriptional Pahlavi/code-points');
require('unicode-6.3.0/blocks/Inscriptional Pahlavi/symbols');
require('unicode-6.3.0/blocks/Inscriptional Pahlavi/regex');

require('unicode-6.3.0/blocks/Old Turkic/code-points');
require('unicode-6.3.0/blocks/Old Turkic/symbols');
require('unicode-6.3.0/blocks/Old Turkic/regex');

require('unicode-6.3.0/blocks/Rumi Numeral Symbols/code-points');
require('unicode-6.3.0/blocks/Rumi Numeral Symbols/symbols');
require('unicode-6.3.0/blocks/Rumi Numeral Symbols/regex');

require('unicode-6.3.0/blocks/Brahmi/code-points');
require('unicode-6.3.0/blocks/Brahmi/symbols');
require('unicode-6.3.0/blocks/Brahmi/regex');

require('unicode-6.3.0/blocks/Kaithi/code-points');
require('unicode-6.3.0/blocks/Kaithi/symbols');
require('unicode-6.3.0/blocks/Kaithi/regex');

require('unicode-6.3.0/blocks/Sora Sompeng/code-points');
require('unicode-6.3.0/blocks/Sora Sompeng/symbols');
require('unicode-6.3.0/blocks/Sora Sompeng/regex');

require('unicode-6.3.0/blocks/Chakma/code-points');
require('unicode-6.3.0/blocks/Chakma/symbols');
require('unicode-6.3.0/blocks/Chakma/regex');

require('unicode-6.3.0/blocks/Sharada/code-points');
require('unicode-6.3.0/blocks/Sharada/symbols');
require('unicode-6.3.0/blocks/Sharada/regex');

require('unicode-6.3.0/blocks/Takri/code-points');
require('unicode-6.3.0/blocks/Takri/symbols');
require('unicode-6.3.0/blocks/Takri/regex');

require('unicode-6.3.0/blocks/Cuneiform/code-points');
require('unicode-6.3.0/blocks/Cuneiform/symbols');
require('unicode-6.3.0/blocks/Cuneiform/regex');

require('unicode-6.3.0/blocks/Cuneiform Numbers and Punctuation/code-points');
require('unicode-6.3.0/blocks/Cuneiform Numbers and Punctuation/symbols');
require('unicode-6.3.0/blocks/Cuneiform Numbers and Punctuation/regex');

require('unicode-6.3.0/blocks/Egyptian Hieroglyphs/code-points');
require('unicode-6.3.0/blocks/Egyptian Hieroglyphs/symbols');
require('unicode-6.3.0/blocks/Egyptian Hieroglyphs/regex');

require('unicode-6.3.0/blocks/Bamum Supplement/code-points');
require('unicode-6.3.0/blocks/Bamum Supplement/symbols');
require('unicode-6.3.0/blocks/Bamum Supplement/regex');

require('unicode-6.3.0/blocks/Miao/code-points');
require('unicode-6.3.0/blocks/Miao/symbols');
require('unicode-6.3.0/blocks/Miao/regex');

require('unicode-6.3.0/blocks/Kana Supplement/code-points');
require('unicode-6.3.0/blocks/Kana Supplement/symbols');
require('unicode-6.3.0/blocks/Kana Supplement/regex');

require('unicode-6.3.0/blocks/Byzantine Musical Symbols/code-points');
require('unicode-6.3.0/blocks/Byzantine Musical Symbols/symbols');
require('unicode-6.3.0/blocks/Byzantine Musical Symbols/regex');

require('unicode-6.3.0/blocks/Musical Symbols/code-points');
require('unicode-6.3.0/blocks/Musical Symbols/symbols');
require('unicode-6.3.0/blocks/Musical Symbols/regex');

require('unicode-6.3.0/blocks/Ancient Greek Musical Notation/code-points');
require('unicode-6.3.0/blocks/Ancient Greek Musical Notation/symbols');
require('unicode-6.3.0/blocks/Ancient Greek Musical Notation/regex');

require('unicode-6.3.0/blocks/Tai Xuan Jing Symbols/code-points');
require('unicode-6.3.0/blocks/Tai Xuan Jing Symbols/symbols');
require('unicode-6.3.0/blocks/Tai Xuan Jing Symbols/regex');

require('unicode-6.3.0/blocks/Counting Rod Numerals/code-points');
require('unicode-6.3.0/blocks/Counting Rod Numerals/symbols');
require('unicode-6.3.0/blocks/Counting Rod Numerals/regex');

require('unicode-6.3.0/blocks/Mathematical Alphanumeric Symbols/code-points');
require('unicode-6.3.0/blocks/Mathematical Alphanumeric Symbols/symbols');
require('unicode-6.3.0/blocks/Mathematical Alphanumeric Symbols/regex');

require('unicode-6.3.0/blocks/Arabic Mathematical Alphabetic Symbols/code-points');
require('unicode-6.3.0/blocks/Arabic Mathematical Alphabetic Symbols/symbols');
require('unicode-6.3.0/blocks/Arabic Mathematical Alphabetic Symbols/regex');

require('unicode-6.3.0/blocks/Mahjong Tiles/code-points');
require('unicode-6.3.0/blocks/Mahjong Tiles/symbols');
require('unicode-6.3.0/blocks/Mahjong Tiles/regex');

require('unicode-6.3.0/blocks/Domino Tiles/code-points');
require('unicode-6.3.0/blocks/Domino Tiles/symbols');
require('unicode-6.3.0/blocks/Domino Tiles/regex');

require('unicode-6.3.0/blocks/Playing Cards/code-points');
require('unicode-6.3.0/blocks/Playing Cards/symbols');
require('unicode-6.3.0/blocks/Playing Cards/regex');

require('unicode-6.3.0/blocks/Enclosed Alphanumeric Supplement/code-points');
require('unicode-6.3.0/blocks/Enclosed Alphanumeric Supplement/symbols');
require('unicode-6.3.0/blocks/Enclosed Alphanumeric Supplement/regex');

require('unicode-6.3.0/blocks/Enclosed Ideographic Supplement/code-points');
require('unicode-6.3.0/blocks/Enclosed Ideographic Supplement/symbols');
require('unicode-6.3.0/blocks/Enclosed Ideographic Supplement/regex');

require('unicode-6.3.0/blocks/Miscellaneous Symbols And Pictographs/code-points');
require('unicode-6.3.0/blocks/Miscellaneous Symbols And Pictographs/symbols');
require('unicode-6.3.0/blocks/Miscellaneous Symbols And Pictographs/regex');

require('unicode-6.3.0/blocks/Emoticons/code-points');
require('unicode-6.3.0/blocks/Emoticons/symbols');
require('unicode-6.3.0/blocks/Emoticons/regex');

require('unicode-6.3.0/blocks/Transport And Map Symbols/code-points');
require('unicode-6.3.0/blocks/Transport And Map Symbols/symbols');
require('unicode-6.3.0/blocks/Transport And Map Symbols/regex');

require('unicode-6.3.0/blocks/Alchemical Symbols/code-points');
require('unicode-6.3.0/blocks/Alchemical Symbols/symbols');
require('unicode-6.3.0/blocks/Alchemical Symbols/regex');

require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension B/code-points');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension B/symbols');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension B/regex');

require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension C/code-points');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension C/symbols');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension C/regex');

require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension D/code-points');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension D/symbols');
require('unicode-6.3.0/blocks/CJK Unified Ideographs Extension D/regex');

require('unicode-6.3.0/blocks/CJK Compatibility Ideographs Supplement/code-points');
require('unicode-6.3.0/blocks/CJK Compatibility Ideographs Supplement/symbols');
require('unicode-6.3.0/blocks/CJK Compatibility Ideographs Supplement/regex');

require('unicode-6.3.0/blocks/Tags/code-points');
require('unicode-6.3.0/blocks/Tags/symbols');
require('unicode-6.3.0/blocks/Tags/regex');

require('unicode-6.3.0/blocks/Variation Selectors Supplement/code-points');
require('unicode-6.3.0/blocks/Variation Selectors Supplement/symbols');
require('unicode-6.3.0/blocks/Variation Selectors Supplement/regex');

require('unicode-6.3.0/blocks/Supplementary Private Use Area-A/code-points');
require('unicode-6.3.0/blocks/Supplementary Private Use Area-A/symbols');
require('unicode-6.3.0/blocks/Supplementary Private Use Area-A/regex');

require('unicode-6.3.0/blocks/Supplementary Private Use Area-B/code-points');
require('unicode-6.3.0/blocks/Supplementary Private Use Area-B/symbols');
require('unicode-6.3.0/blocks/Supplementary Private Use Area-B/regex');

// bidi mirroring:

require('unicode-6.3.0/bidi-mirroring')[ codePoint ]; // lookup array

// bidi brackets:

require('unicode-6.3.0/bidi-brackets')[ codePoint ]; // lookup array

require('unicode-6.3.0/bidi-brackets/Open/code-points');
require('unicode-6.3.0/bidi-brackets/Open/symbols');
require('unicode-6.3.0/bidi-brackets/Open/regex');

require('unicode-6.3.0/bidi-brackets/Close/code-points');
require('unicode-6.3.0/bidi-brackets/Close/symbols');
require('unicode-6.3.0/bidi-brackets/Close/regex');
```

## Author

| [![twitter/mathias](https://gravatar.com/avatar/24e08a9ea84deb17ae121074d0f17125?s=70)](https://twitter.com/mathias "Follow @mathias on Twitter") |
|---|
| [Mathias Bynens](https://mathiasbynens.be/) |

## License

This module is available under the [MIT](https://mths.be/mit) license.
