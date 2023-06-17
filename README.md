## RepoWiki

RepoWiki is a wiki server that pulls the newest content from the given git repo,
which can be hosted anywhere as long as it supports the git protocol.  
After getting the newest content, it is displayed in a raneto instance on port 3000.

### Env vars

#### Mandatory
- GITURL: The URL for the git repo to clone

#### Optional
- FETCH_INTERVAL: The interval in seconds between updates for the articles (default: 300)
- GIT_REPO_LOCATION: The location of where the cloned repo is stored (default: ./repo)
- LOG_FILE_LOCATION: The location of the log file (default: ./RepoWiki.log)
- LOG_LEVEL: The level to be logged, the higher, the more will be logged (default: 1)

### Road map

- Support for multiple repositories at once
- Support for private repositories