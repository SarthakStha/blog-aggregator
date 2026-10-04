# gator
gator in a CLI based blog aggregator. It provides CLI commands that allows multiple user to login and add feeds they want to aggregate. It also allows user to follow other's feed and browse the content of their followed feeds.

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
gator uses Postgres as the database to store all it's information. You can find the installation specific to your OS on their [official website](https://www.postgresql.org/download/).

### 3. Setting up the Database
Once PostgreSQL is installed and the PSQL server is running. Create a database for this project.
```
CREATE DATABASE <database name>
\c <database name>
```

Once the database is created, we need to store the database connection string in the config file. The database connection string will be in the format: `protocol://username:password@host:port/database`.

The config file is expected to have the name `.gatorconfig.json` and be present in the home directory `~/`. Here is a sample structure of the file:
```
{
    "db_url": "database connection string"
}
```
(Note: This file will also be used to store the currently logged in user.)

### 4. Migrate the Database Schemas
All the migration queries have been provided in sql/schema. It is highly recommended to use [goose](https://github.com/pressly/goose) for the migration as the annotation for it have already been provided. If using goose you can simply cd into the schema directory and run:
```
goose <database connection string> up
```
In case you need to revert a migration, you can run:
```
goose <database connection string> down
```

### 5. Install gator
Run the following command. This should compile and install the `gator` command into your bin folder.
```
go install github.com/SarthakStha/gator@latest
```

## List of Commands
|Name       | Description           |
|-----------|-----------------------|
|login      | Takes 1 argument, name of a currently registered user.|
|register   | Takes 1 argument, name of a new user, also logins as that user.|
|reset      | Completely resets all the content of the database.|
|users      | Lists all the users currently registered.|
|addfeed    | Takes 2 arguments, the title of the feed and the URL of the feed, adds it to the feeds table.|
|removefeed | Take 1 argument, the URL of the feed and removes it from the feeds table.|
|feeds      | Lists all feeds currently available.|
|follow     | Takes 1 argument, the URL of the feed to follow. Since, we don't allow multiple feeds with the same URL. This still allows the user to still view the contents of other's feed.|
|unfollow   | Takes 1 argument, the URL of the feed to unfollow. Removes the relation between the user and the feed.|
|following  | Lists all the feed followed by the user.|
|agg        | Takes 1 argument, time string('5s', '30s') which is the duration between each refresh. It will start refreshing all feeds in a loop starting with the oldest first.|
|browse     | Optionally takes 1 argument, number of posts to show. If no argument is provided it shows 2 posts by default.|

**Usage**
```
gator <command name> <argument 1> <argument 2>
```
