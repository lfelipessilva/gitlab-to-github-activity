
# Go Commiter

This will show your gitlab activity from the last year in your github account without needing to share sensitive data.



## How to use

1. Clone this repo
```bash
  git@github.com:lfelipessilva/gitlab-to-github-activity.git
```
2. Enter the repo
```bash
  cd gitlab-to-github-activity
```
3. Create necessary folders and files
```bash
  touch commits.json && mkdir ./repo
```
4. Start a git repository on `./repo`
```bash
  cd ./repo && git init
```
5. Find your activity in https://gitlab.com/users/{YOUR_USER}/calendar.json (you have to be logged in gitlab)
6. Copy everything and paste inside `./commits.json`
7. Run the project
```bash
  go run main.go
```
8. Then push the commits to a private repository in your github (after setting up the remote)
```bash
  cd ./repo && git push
```


You can now periodically just get your data and past it all inside `commits.json`, it will only generate commits from the date of the last commit and on.


## Authors

- [@luisfelipess](https://www.github.com/luisfelipess)
