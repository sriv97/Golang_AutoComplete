
# AutoComplete Backend API

An API Service to display the top 25 most frequent words used by Shakespeare based on searched fragment
Created using Shakespeare-Complete.txt 

## Running in Web Browser
To run the API in your web browser, change the term after ?=, and hit enter [Heroku](https://golang-autocomplete.herokuapp.com/autocomplete?term=th)

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
