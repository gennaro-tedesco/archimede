<h1 align="center">
  archimede
</h1>

<h2 align="center">
  <a href="#" onclick="return false;">
    <img alt="PR" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat"/>
  </a>
  <a href="https://golang.org/">
    <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=flat&logo=go&logoColor=white"/>
  </a>
  <a href="https://github.com/gennaro-tedesco/archimede/releases">
    <img alt="releases" src="https://img.shields.io/github/release/gennaro-tedesco/archimede"/>
  </a>
</h2>

<h4 align="center">Unobtrusive project information fetcher</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a>
</h3>


The unobtrusive, fast and informative project information fetcher. Whether it is a new git project or old legacy code, visualise a quick summary of directory structure, files composition, disk space and status.


## Installation
Go get it!
```
go get github.com/gennaro-tedesco/archimede
```

## Usage
Various flags allow to customise the output: see `archimede --help` for full details or examples below

| flag             | type   | description                               | default
|:---------------- |:------ |:----------------------------------------- |:-------
| -s/--short       | bool   | display output in short format?           | false
| -g/--git         | bool   | include `./.git` folder in files stats?   | false
| -e/--exclude-dir | string | directory to exclude from stats/counts    | none
| -v/--exclude-file| string | file type to exclude from stats/counts    | none
| -c/--colour      | string | choose display text color (see `--help`)  | cyan
| -d/--delimiter   | string | choose delimiter character                | empty string
| -h/--help        |        | open the help page                        |

Notice that the default display option in long format (no `-s` flag) makes use of unicode
terminal characters that may not render perfectly with all terminals/fonts: if so use `-s`.

## Examples
```
# short format, blue text
archimede -s -c blue
```
<details>
  <summary>Show output</summary>

  <img alt="" src="https://user-images.githubusercontent.com/15387611/120498035-34fb8200-c3bf-11eb-83c7-a67ed06f3b42.png">
</details>

```
# change delimiter, exclude test folder
archimede -d":" -e test
```
<details>
  <summary>Show output</summary>

  <img alt="" src="https://user-images.githubusercontent.com/15387611/120498908-e6021c80-c3bf-11eb-9d03-bbc502fde473.png">
</details>

```
# short format, exclude .vim files
archimede -s -v .vim
```
<details>
  <summary>Show output</summary>

  <img alt="" src="https://user-images.githubusercontent.com/15387611/120499354-42fdd280-c3c0-11eb-9a38-28a627681068.png">
</details>

```
archimede --help
```
<details>
  <summary>Show output</summary>

  <img alt="help" src="https://user-images.githubusercontent.com/15387611/120497389-a4bd3d00-c3be-11eb-8134-3e3fd10f34e6.png">
</details>

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

