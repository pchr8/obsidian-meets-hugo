# This fork 
In this fork I fix things I need for my own website, <https://serhii.net>, posts for which are also written in Obsidian. 
Specifically, this fixes:
- Use Obsidian file front-matter title if present (instead of the filename)
- Allow writing directly to site root with empty `--sub-path= ` by fixing double-leading-slash assets-linking bug 
- Cyrillc now allowed in filenames 
- Leave `_index.md` as-is without renaming them to `-index.md` — that way these files can be created directly in Obsidian instead of manually in Hugo

This is literally my first 2h coding something in Go with which I'm _completely_ unfamiliar, so YMMV and be careful :) 

- - -

[![codecov](https://codecov.io/gh/ukautz/obsidian-meets-hugo/branch/main/graph/badge.svg?token=89QLLNR10S)](https://codecov.io/gh/ukautz/obsidian-meets-hugo)

# `omh` - Obsidian Meets Hugo

Command line tool to marry [Obsidian](https://obsidian.md/) vaults to [Hugo](https://gohugo.io/) published websites.

```sh
# convert and copy all obsidian notes in directory (and sub-directories)
#  into hugo path in `/path/to/hugo/content/obsidian` and respective static
#  files into `/path/to/hugo/static/obsidian`
$ omh \
    --obsidian-root /path/to/obsidian \
    --hugo-root /path/to/hugo
```

See `omh -h` for extended options.

_Note: on Mac you can find your iCloud synced notes in `~/Library/Mobile\ Documents/iCloud\~md\~obsidian/Documents/`_

## Install

```sh
$ go install github.com/ukautz/obsidian-meets-hugo/cmds/omh
```

## Use-Case

This command line tool allows you to easily export an Obsidian vault, or a sub-set thereof, into a Hugo published website.

I am using this tool to publish my own notes - that follow a standard in between [Zettelkasten and Wikipedia](https://en.wikipedia.org/wiki/Zettelkasten) - to my [Blog](https://ulrichkautz.com), as you can see here: <https://ulrichkautz.com/zettel/>. This way I have them easily available and can reference them in Blog entries.

## License

MIT

## Alternatives

Things I found that do not exactly fit my needs, but maybe yours:

- <https://github.com/khalednassar/obyde>
- <https://github.com/jackyzha0/hugo-obsidian>
