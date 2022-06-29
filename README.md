<h1>movie-rater - a dead-simple tool for personal movie ratings</h1>

Just run `movie-rater` and it'll prompt you to enter a title, a comment and a rating (float 0-10). A file called `movieRatings.json` will be created 
in the same directory where `movie-rater` is located. That's it... That's all it does.

<h1>build from source</h1>
just clone the repo and run <code>go build -o movie-rater -ldflags "-s -w" *.go</code>
