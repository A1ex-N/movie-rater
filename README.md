<h1>movie-rater - a dead-simple tool for personal movie ratings</h1>

Just run `movie-rater` and it'll prompt you to enter a title, a comment and a rating (float 0-10 with 1 decimal place). A file called `movieRatings.json` will be created 
in the same directory where `movie-rater` is located. That's it... That's all it does.

<h1>build from source</h1>
just clone the repo and run <code>go build -o movie-rater -ldflags "-s -w" *.go</code>
<br>
<br>
<h1>Features i'd like to have but i'm too dumb/lazy to add</h1>
<ul>
  <li> Fetch rating for titles from an online database like IMDb or rotten tomatoes to compare to your ratings and store as <code>imdb_rating</code> in json file
  <li> Maybe fetch genre for titles from an online database and store it in a <code>genre</code> key in the json file
  <li> Sort movies by rating and output to user, either in ascending or descending order as chosen by the user
  <li> Fetch runtime of movies and store as <code>runtime</code> key in the json file and then add up all of the runtimes and display how much of your life you have wasted watching movies
</ul>


Feel free to make pull requests. This is my first Go project 
