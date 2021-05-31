<h1 align="center">
  archimede
</h1>

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

| flag      | description             | default
|:--------- |:----------------------- |:-------
| -s (bool) | short output?           | false
| -g (bool) | include `./.git` files? | false
| -c (str)  | output text color       | cyan
| -d (str)  | delimiter character     | empty


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

