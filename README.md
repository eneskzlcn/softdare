### About Project

This is a social media project that users can post their
writings, follow other users, comment on other users posts,
like other users posts or comments.


### Local Setup To Run

You need the docker installed on your machine to run the
application.

- `docker compose up -d` to start rabbitmq and postgres database services.
- `go mod tidy` to install all needed dependencies and a cleanup for not needed ones...
- `make migrate-tables` to create needed tables on postgres database
- `make build` to build the application and `make run` to start.
or you can directly use `make start` to build first then run.

### Additional Information About Local Setup
- You can find all the database schema under postgres/migration/schema.sql. If you
want to contribute or use for your own purposes, you can simply add the tables you
want to create to schema.sql and the command `make migrate-tables` will rebuild 
all the tables for you after you clean up the database with the command `make drop-tables`
(the thing you need to remember, you should add your new table to postgres/migration/drop.sql
too to clean up correctly for next `make drop-tables` commands.).

- If you want to clean up all the tables in database and restart again you can
  call `make drop-tables` to clean up the database , and `make migrate-tables` to create
  the tables again.

- If you want to consume one more queue from rabbitmq, you need to add the queue name to
the config files which is under the .dev/ folder. After you add that queue name to the config
file, the program will create the queue automatically when it restarts.

- If you want to add one more consumer for a queue, you can find the consumers under
internal/client/queue. You can add a consumer function like the others I write and provide
a consume operation for your purposes.

### Local Setup For Tests
- `go mod tidy` to install all needed dependencies and a cleanup for not needed ones...
- `make clean` to clear all additional files.
- `make generate-mocks` to generate all needed mock to run tests.
- `make test` to run all written tests.

### Architectural Decisions
This project teach me a lot of on how you need to choose
an architecture for your application. I was generally using DDD approach
to create backend applications which recommends that you need to
keep inside the things that domain consist of. An example
if you are creating an order domain, then all structures, functionalities,
tests, business and database access operations should be in the package of that domain.

So first, I tried to make all the domains like post, login, comment, user etc.
seperated and independent of each other. After I structured the project like 
that, I enjoyed and thought to gain a victory about that. Then when time comes
to the creating new features on that architecture things started to get more complicated.
Every need were causing a dependency between domains/subdomains. The most great
example was the entities. I thought any way to protect that structure
so I see that it will be costly. So I decided to move all the entities out of
the domains and all the domains can reach them without creating a dependency
between domains. The folder structure were like;

- entities/
- login/
- post/

After this change, I was comfortable to reach the entities from anywhere
and it was easy to add or remove an entity for new feature purposes.

Working on new things obligate me to repeat the implementation of same functions
just to keep the domains independent. An example of that is a data access or business
function one domain has is also used by another domain. But you can not pass directly
to maintain independence.

I googled about that and see the architecture generally used by Java
or C# backend programmers named `n-Layered Architecture`. I realised that
the architecture I need to use in a monolithic web application should
be that. Because there is no domain specific abstractions so that
you can make a folder/package structure like;
```
   - entity/
   - services/
   - repository/
   - api/
   - web/
   - mobile/
   ...
```
At last, I chose that architecture as a project architecture and everything
become more comfortable, flexible and easier.

As a result, I tried to feel like a startup company that must keep things faster
to deploy their app. Generally they prefer monolithic way to deploy the first
release as much as faster and then seperated the system to the little services when needed( microservices ).
So I think that making all the domains seperated even doing a monolithic application effect the process 
after first release, so it will be easier to decouple the services as independent microservices.
But I learned that making a monolithic application in that
way is very hard and painful.

### External Modules Used In Project
- `html/templates` built-in module as a html parser and renderer technology.
- `go.uber.org/zap` logger module to handle logging operations.
- `go-ozzo/ozzo-validation` validation module to validate special structures and inputs.
- `/rs/xid` xid module to generate xid for database entities.
- `spf13/viper` configuration module to handle and structure environment specific configurations.
- `golangcollege/sessions` session module to handle all session operations on the application.

### Testing Modules
- `DATA-DOG/go-sqlmock` sql db mock module to mock the database on
testing repository layer.
- `golang/mock` as a mock module to mock the interfaces and test
the objects independently.
