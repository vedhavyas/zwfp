# zwfp
[![Build Status](https://travis-ci.org/vedhavyas/zwfp.svg?branch=master)](https://travis-ci.org/vedhavyas/zwfp)
[![GitHub tag](https://img.shields.io/github/tag/vedhavyas/zwfp.svg)](https://github.com/vedhavyas/zwfp/tags)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub issues](https://img.shields.io/github/issues/vedhavyas/zwfp.svg)](https://github.com/vedhavyas/zwfp/issues)
![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/vedhavyas/zwfp.svg)](https://github.com/vedhavyas/zwfp/pulls)
[![Website](https://img.shields.io/website-up-down-green-red/http/vedhavyas.com.svg?label=my-website)](https://vedhavyas.com)

ZWFP - Zero-width fingerprinting uses [Zero Width](https://en.wikipedia.org/wiki/Zero_width) characters to embed text into a given plain text.
Inspiration for the project from [this article](https://medium.com/@umpox/be-careful-what-you-copy-invisibly-inserting-usernames-into-text-with-zero-width-characters-18b4e6f17b66) by [Tom](https://medium.com/@umpox)

## Demo
![ZWFP Demo](https://j.gifs.com/yrnwx7.gif)

## Getting Started

Installation assumes that you have Go environment configured.
 
### Installing

Go get the project with following command

```
go get -u github.com/vedhavyas/zwfp/cmd/zwfp/...
```

## Running the tests

Once inside project' folder, simply run `go test ./...` to run the tests.

### Usage

```bash
Usage:
         ./zwfp PlainText Key
                 Embeds Key into PlainText

         ./zwfp EmbedText
                 Extracts Key from EmbedText

```

## Built With

* [Go](https://golang.org/)
* [Glide](https://glide.sh/) - Dependency Management
* [HomeDir](https://github.com/mitchellh/go-homedir) - Used to get User's home directory
* [Crypto](https://golang.org/x/crypto) - Used to get password from terminal securely

## Contributing

PRs, Issues, and Feedback are very welcome and appreciated.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/vedhavyas/twothy/tags). 

## Authors

* **Vedhavyas Singareddi** - [Vedhavyas](https://github.com/vedhavyas)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
