# boggle-saleswhale
Assignment for saleswhale. Create api to play boggle

## Installation
* Run `docker build . -t boggle`
* Run `docker run -i -d  -p 8080:8080 -t boggle`
* You can now access the api server at `localhost:8080`

Refer `apidoc.html` for examples for below mentioned apis

## Apis
* `/api/health` for checking if server is up
* `/api/users/new/start` for starting a game for a new user. Can set duration, default is 1 minute. Returns userid, gameid and board. Userid and gameid are to be used in further requests.
* `/api/users/{userId}/start` for starting a game for existing user. Same as above but for a specific user
* `/api/users/{userId}` for getting user info
* `/api/users/{userId}/move` for user to guess a word in the game he is currently playing. Returns if guess is correct with score
* `/api/game/{gameId}` for getting info a particular game
* `/api/game/{gameId}/score` to get the score obtained in a particular game

One user can only play one game at a time 

To run locally without docker, you can run `go run main.go`

### Config
* `dictionary_path`: path of file storing dictionary words. One word in one line
* `board_path`: path of file storing board config. One board config per line. Board config can be of custom size. Default is 16 characters(4x4)