# Backend

## Endpoints

Alle Endpoints hinter: /api/v0/

- /ping: pong
- /login: User Authentication (how?)
- /img/(image hash sha256) for all images (JPEG / WebP / AVIF Bilder)
- /me/img/list
  - JSON response with hash of all images from logged in user
- /me/img/upload
  - Push image to server
- /me/friends/list
  - JSON response with all friends of user
- /me/friends/add
  - Sent friend request (JSON response wether user exists or not)
- /me/friends/requests/sent
  - Sent friend requests that have not yet been accepted
- /me/friends/requests/self
  - Friend requests that have been sent to USER
