# doko-graphql
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

query post {posts{posts{title,txt,latitude,longitude}}}

 mutation CreatePosts($newposts: [InputPost!]!) {
  createPosts(input: $newposts)
}

query headers
{
  "authorization":"Bearer xxx"
}

query variables 
{
  "newposts" : [
    {
      "title": "title1",
      "txt": "ola",
        "latitude":1.111111,
        "longitude":2.222222
    },
    {
      "title": "title2",
      "gps" : {
        "latitude":3.111111,
        "longitude":4.222222
      }
    }
  ],
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
  }
}


## TODO