
# AutoComplete Backend API
An API Service to display the top 25 most frequent words used by Shakespeare based on searched fragment

**_The searched value can exist anywhere in the word, results are returned based on highest frequency_**


## Running in Web Browser
To run the API in your web browser, change the term after ?=, and hit enter <a href="https://golang-autocomplete.herokuapp.com/autocomplete?term=th" target="_blank">Run Online </a>
![Web Example](https://github.com/sriv97/Golang_AutoComplete/blob/master/images/web.png)


## Running by curl in Terminal / Command Prompt
Copy and Paste the command below into a terminal / command prompt window

In the below example the search fragment is set to ```th```, replace it with the actual fragment

```curl https://golang-autocomplete.herokuapp.com/autocomplete\?term\=th```
![Curl Example](https://github.com/sriv97/Golang_AutoComplete/blob/master/images/curl.png)

## Running in Postman

Open Postman and create a GET Request to 
```https://golang-autocomplete.herokuapp.com/autocomplete?term```

In the Query Params, enter in the desired fragment as the value for term, then click Send

![Postman Example](https://github.com/sriv97/Golang_AutoComplete/blob/master/images/postman.png)

## View sample data
Phrases (th, fr, pi, sh, wu, ar, il, ne, se, pl) [Sample Fragments](https://github.com/sriv97/Golang_AutoComplete/blob/master/results.md)

### References used
[Guide to setting up Golang on Heroku](https://github.com/heroku/go-getting-started)

[Original Text](https://github.com/sriv97/Golang_AutoComplete/blob/master/shakespeare-complete.txt)
