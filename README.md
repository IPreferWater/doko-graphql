# doko-graphql
go run github.com/99designs/gqlgen generate

## playground

query post {posts{posts{title,text,latitude,longitude}}}

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
      "txt": "ola2",
        "latitude":2.111111,
        "longitude":3.222222
    }
  ]
}

## TODO