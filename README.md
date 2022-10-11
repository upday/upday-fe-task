Welcome to the UPDAY Front-end task assignment! :relaxed: 

Before we move on into the task let's introduce some important concepts:
- _Board_: contains a set of news with a specific language (determined by each individual board)
- _News_: content created by an editor to a specific board

## Task
In this task, you will need to create a small front-end system that allows our editors to log in and manage the news according to the boards.

### Pre-Requirements
- Run our micro-service locally: ```docker run -p 8080:8080 upday/upday-fe-task```
- Access the [API documentation](https://upday.github.io/upday-fe-task), or access [the micro-service repository](https://github.com/upday/upday-fe-task) to get more details about the API
- Choose a modern Typescript web framework. We're more familiar with VUE and React, but feel free to choose any framework that you enjoy and let us know why you did so

### Required Features
These are the required features, but this does not mean that you cannot add new ones. :smile: 

#### Login Screen
All users need to do the login before accessing the system. The access control should be made/managed on the front-end side, the micro-service doesn't provide any access control API.
- only the email field is mandatory
- all valid emails can be considered a "valid-user"

#### Home Screen
The user should be able to choose a board, view the board's news and manage the news.
- provide a board selector, the selector should display the board name instead of the language
- list of news filtered by board
- allow the editors to manage the news [create/update/delete/modify the status]

#### News Crud
Allow the editors to create/edit the news
- should use the logged user email as the default author
- one of the boards simulates an error, so please handle this accordingly

### Requirements
- you should store your project in a public git repository and let us participate in your commit progress. So please don't send us a git project with a single commit `first commit` :sweat_smile: 
- use typescript as the main language
- use unit tests, but you don't need to have 100% coverage!

> **Note:** We expect you to not use any frontend libraries (e.g. **Material UI, Formik, Formsy**). Reason for this is to also see your **core skills** when there is no good library exists.
