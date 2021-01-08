# Golang server

## Required packages

- Pq (Postgresql driver) `go get github.com/lib/pq`
- Gorilla mux `go get -u github.com/gorilla/mux`
- Gorilla handlers `go get -u github.com/gorilla/handlers`
- Whirlpool (far better than sha256) `go get github.com/jzelinskie/whirlpool`

## Usage

`go run server.go`

## Postgresql setup

Launch :
```
sudo service postgresql-13 start
```

Init :
```
sudo -u postgres psql
ALTER USER postgres WITH PASSWORD 'test';
```

## Database setup

```
CREATE TABLE users(
  id VARCHAR(32) PRIMARY KEY NOT NULL,
  username VARCHAR(24) NOT NULL,
  password VARCHAR(128) NOT NULL,
  email VARCHAR(255) NOT NULL,
  verified INT DEFAULT 0,
  verify_hash VARCHAR(32) NOT NULL,
  is_admin INT DEFAULT 0,
  is_banned INT DEFAULT 0,
  last_log DATE DEFAULT CURRENT_DATE,
  reg_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE articles(
	id VARCHAR(32) PRIMARY KEY NOT NULL,
	uid VARCHAR(32),
	titre VARCHAR(255) NOT NULL,
	contents TEXT,
	img_link VARCHAR(255) DEFAULT NULL,
	creation_date DATE DEFAULT CURRENT_DATE,
	CONSTRAINT fk_uid
		FOREIGN KEY(uid)
			REFERENCES users(id)
);

CREATE TABLE comment(
	id VARCHAR(32) PRIMARY KEY NOT NULL,
	uid VARCHAR(32),
	aid VARCHAR(32),
	contents VARCHAR(255),
	creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_uid
		FOREIGN KEY(uid)
			REFERENCES users(id),
	CONSTRAINT fk_aid
		FOREIGN KEY(aid)
			REFERENCES articles(id)
);
```

## API calls

```
(PATCH) /api/ban/{uid} (string => void) : Bans the given user (sets his is_banned flag to 1)
(GET) /api/is_auth/{type} (string => void) : Checks if the user is connected using his session token (also checks if the user is admin is type = admin)
(GET) /api/logout (void => void) : Ends the current session (deletes the session token)
(GET) /api/login (string, string => void) : Starts a session and generates a session token (initially planned to be a POST request, but set temporarly on GET due to cors issues)
(POST) /api/register (string, string, string => void) : Creates a new user account
(PATCH) /api/article/update (string, string, string, string => string) : Updates a given article, returns the article id
(POST) /api/article/create (string, string, string => string) : Creates a new article, returns the newly created article id
(DELETE) /api/article/delete (string => void) : Deletes a given article
(GET) /api/article/fetch (void => []Article) : Retrieves all articles
(GET) /api/article/read/{aid} (string => Article) : Retrieves an article
(PATCH) /api/comment/update (string, string => void) : Updates a given comment
(POST) /api/comment/create (string, string => void) : Creates a new comment under the given article
(DELETE) /api/comment/delete (string => void) : Deletes a given comment
(GET) /api/comment/fetch/{aid} (string => []Comment) : Retrieves all comments under the given article
```

## Technical choices (Asked by the uni subject)

- The passwords are salted and hashed in Whirlpool (strong hash algorithm)
- Each identifier is a randomly generated md5 hash
- Sessions token are salted (using a random nonce) and encrypted in AES-256
- A secret api key for token encryption is generated each time the server is opened, the devs don't even know the secret key, c'est fantabulistique
- The VueJS FrontEnd and GoLang BackEnd communicate through XHTTP Requests (Using Axios front-side, there still are lots of bugs about the linking)
- Sessions are fully stateless (The session token contains the user id hash)
- All cookies are treated back-end side
- All returned data are negociated in JSON (application/json and converted using the GoLang json marshall module)
- It's far more interesting to learn programming languages we never used before instead of using JAVA (And it was a very nice experience... EXCEPT FOR CORS !!!)
- Why PostgreSQL ? Because companies seems to like it, companies have money, and I like money, therefore I use it

## Main difficulties (Also asked by the uni subject)

- CORS
- CORS
- CORS
- Axios behaviour and JS Promises
- Divergences of data structures between JavaScript and GoLang

## Future features

- Multi-language support (French and Japanese translations)
- More features about User profiles like profile pictures
- Like/Dislike a comment
- Article categories
- Comment answers
- And more features, maybe someday, after fixing the remaning bugs
