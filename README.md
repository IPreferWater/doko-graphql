# graphql-theory
go run github.com/99designs/gqlgen generate

## playground

query notes{
  notes{
    name,
    steps{
      title,
      txt,
      url}
  }
}

mutation CreateNote($newnote: NewNote!) {
  createNote(input: $newnote)
}

mutation DeleteNote($id: Int!) {
  deleteNote(input: $id)
}

query variables 

{
  "newnote": {
    "name": "random",
    "steps": [
      {
        "title": "title-1",
        "txt": "txt-1",
        "url": "url-1"
      },
      {
        "title": "title-2",
        "txt": "txt-2"
      },
      {
        "title": "title-3",
        "txt": "txt-3"
      }
    ]
  },
  "id":1
}



## TODO
middleware : 
- login

graphql : 
- use model instead of generated
#used autobind:
- query login