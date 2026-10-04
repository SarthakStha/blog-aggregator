# gator
gator in a CLI based blog aggregator. It provides CLI commands that allow the multiple user to login and add feeds they want to aggregate. It also allows user to follow other's feed and browse the content of their followed feeds.

## Installation 
### 1. Install Go
gator is build in golang. So, to install it ensure you have the latest version of Go.
The official Go installation guide can we found here: [Go Download and install](https://go.dev/doc/install).

Alternatively, you can also use the go webi install [link](https://webinstall.dev/golang/).

**For mac & linux:**
```
curl -sS https://webi.sh/golang | sh
source ~/.config/envman/PATH.env
```

**For windows:**
```
curl.exe https://webi.ms/golang | powershell
```

### 2. Install PostgreSQL
gator uses Postgres as the database to store all it's information. You can find the installation specific to your OS in their [official website](https://www.postgresql.org/download/).
