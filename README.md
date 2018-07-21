# Go Deploy

Sample app as a reference for how to deploy Go to Heroku.
[This is a helpful link](https://devcenter.heroku.com/articles/getting-started-with-go)

### Notes
- Uses godep but maybe should use govendor instead.
- Procfile must specific entry point (aka compiled file). This would be more complicated
with background processes or more `main` entries.
- PORT var. Need it locally, is assigned by Heroku automatically. This is kind of a hacky solution,
otherwise you'd have to specific port every time you ran your server locally. There's probably
a cleaner way but it works.
