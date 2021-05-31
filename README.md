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


Unobtrusive, fast and informative: `archimede` is the project directory information fetcher. Whether it is a new git project or old legacy code, visualise a quick summary of project structure, files composition, disk space and status.


## Installation
Go get it!
```
go get github.com/gennaro-tedesco/archimede
```

## Usage
Many flags allow to customise the output: see `archimede --help` for full details or examples below

| flag      | description                               | default
|:--------- |:----------------------------------------- |:-------
| -s (bool) | display output in short format?           | false
| -g (bool) | include `./.git` files in files counts?   | false
| -c (str)  | choose display text color (see `--help`)  | cyan
| -d (str)  | choose delimiter character                | empty


## Examples
```
archimede -s -d":"
```
<details>
  <summary>Show output</summary>

  <img alt="solarized" src="https://user-images.githubusercontent.com/15387611/114108427-5cf4ca80-98d3-11eb-8b39-99600dd42807.png">
</details>

```
archimede -c blue -g
```
<details>
  <summary>Show output</summary>

  <img alt="solarized" src="https://user-images.githubusercontent.com/15387611/114108427-5cf4ca80-98d3-11eb-8b39-99600dd42807.png">
</details>


## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

