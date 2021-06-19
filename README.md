# YTLocker
Automatic creation of playlists to manage YouTube subscriptions. Live: https://ytlocker.com/

- Create a playlist on https://ytlocker.com/.
- Add playlist to your youtube account.
- Add watched channels to playlist.
- Wait for your watched channels to upload videos.
- Watch the videos in your new playlist.

## Development
- Testability was a major focus, all services are fully tested.
- Typescript is awesome and super fun to work in.
- Api and DB hosted on DigitalOcean, Website on Netlify.

## Setup

- `git clone https://github.com/Killian264/YTLocker.git` clone this repo
- `cd YTLocker` move into main directory
- `docker-compose.yml` Set YOUTUBE_API_KEY retrieved from Google Cloud
- `go run golocker/scripts/oauth-generate/main.go` to generate oauth secrets
- `yarn install` is required due to docker memory issues
- `docker-compose up` run services listed below

## Docker Services
- See `docker-compose.yml` for service secrets

| Service | Url |
| ------ | ------ |
| Storybook | `localhost:6006` |
| Website | `localhost:3000` |
| Golang API | `localhost:8080` |
| MySQL DB | `localhost:3306` | 

## Contributing
- All pull requests are welcome, but creating a ticket first is suggested.
- All code will be reviewed before being merged into prod.
